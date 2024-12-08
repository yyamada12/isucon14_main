package main

import (
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	_ "net/http/pprof"

	"github.com/felixge/fgprof"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type SyncMap[K comparable, V any] struct {
	m  map[K]*V
	mu sync.RWMutex
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{m: map[K]*V{}}
}

func (sm *SyncMap[K, V]) Add(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = &value
}

func (sm *SyncMap[K, V]) Get(key K) *V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.m[key]
}

func (sm *SyncMap[K, V]) Clear() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m = map[K]*V{}
}

type ChairLocationSummary struct {
	ChairID       string     `db:"chair_id"`
	TotalDistance int64      `db:"total_distance"`
	UpdatedAt     *time.Time `db:"updated_at"`
}

var userMap = NewSyncMap[string, User]()
var chairLocationSummaryMap = NewSyncMap[string, ChairLocationSummary]() // chair_id -> ChairLocationSummary

func LoadMap() {
	LoadUserFromDB()
	LoadChairLocationSummaryFromDB()
}

func LoadUserFromDB() {
	// clear sync map
	userMap.Clear()

	var rows []*User
	if err := db.Select(&rows, `SELECT * FROM users`); err != nil {
		log.Fatalf("failed to load users: %+v", err)
		return
	}
	for _, row := range rows {
		// add to sync map
		userMap.Add(row.ID, *row)
	}
}

func LoadChairLocationSummaryFromDB() {
	// clear sync map
	chairLocationSummaryMap.Clear()

	var rows []*ChairLocationSummary
	if err := db.Select(&rows, `SELECT chair_id,
									SUM(IFNULL(distance, 0)) AS total_distance,
									MAX(created_at)          AS updated_at
								FROM (SELECT chair_id,
											created_at,
											ABS(latitude - LAG(latitude) OVER (PARTITION BY chair_id ORDER BY created_at)) +
											ABS(longitude - LAG(longitude) OVER (PARTITION BY chair_id ORDER BY created_at)) AS distance
										FROM chair_locations) tmp
								GROUP BY chair_id`); err != nil {
		log.Fatalf("failed to load chair_locations summary: %+v", err)
		return
	}
	for _, row := range rows {
		chairLocationSummaryMap.Add(row.ChairID, *row)
	}
}

func main() {
	http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	mux := setup()
	LoadMap()

	slog.Info("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}

func setup() http.Handler {
	host := os.Getenv("ISUCON_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("ISUCON_DB_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("failed to convert DB port number from ISUCON_DB_PORT environment variable into int: %v", err))
	}
	user := os.Getenv("ISUCON_DB_USER")
	if user == "" {
		user = "isucon"
	}
	password := os.Getenv("ISUCON_DB_PASSWORD")
	if password == "" {
		password = "isucon"
	}
	dbname := os.Getenv("ISUCON_DB_NAME")
	if dbname == "" {
		dbname = "isuride"
	}

	dbConfig := mysql.NewConfig()
	dbConfig.User = user
	dbConfig.Passwd = password
	dbConfig.Addr = net.JoinHostPort(host, port)
	dbConfig.Net = "tcp"
	dbConfig.DBName = dbname
	dbConfig.ParseTime = true

	_db, err := sqlx.Connect("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}
	db = _db

	for {
		err := db.Ping()
		if err == nil {
			break
		}
		log.Print(err)
		time.Sleep(time.Second * 2)
	}
	log.Print("DB ready!")

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.HandleFunc("POST /api/initialize", postInitialize)

	// app handlers
	{
		mux.HandleFunc("POST /api/app/users", appPostUsers)

		authedMux := mux.With(appAuthMiddleware)
		authedMux.HandleFunc("POST /api/app/payment-methods", appPostPaymentMethods)
		authedMux.HandleFunc("GET /api/app/rides", appGetRides)
		authedMux.HandleFunc("POST /api/app/rides", appPostRides)
		authedMux.HandleFunc("POST /api/app/rides/estimated-fare", appPostRidesEstimatedFare)
		authedMux.HandleFunc("POST /api/app/rides/{ride_id}/evaluation", appPostRideEvaluatation)
		authedMux.HandleFunc("GET /api/app/notification", appGetNotification)
		authedMux.HandleFunc("GET /api/app/nearby-chairs", appGetNearbyChairs)
	}

	// owner handlers
	{
		mux.HandleFunc("POST /api/owner/owners", ownerPostOwners)

		authedMux := mux.With(ownerAuthMiddleware)
		authedMux.HandleFunc("GET /api/owner/sales", ownerGetSales)
		authedMux.HandleFunc("GET /api/owner/chairs", ownerGetChairs)
	}

	// chair handlers
	{
		mux.HandleFunc("POST /api/chair/chairs", chairPostChairs)

		authedMux := mux.With(chairAuthMiddleware)
		authedMux.HandleFunc("POST /api/chair/activity", chairPostActivity)
		authedMux.HandleFunc("POST /api/chair/coordinate", chairPostCoordinate)
		authedMux.HandleFunc("GET /api/chair/notification", chairGetNotification)
		authedMux.HandleFunc("POST /api/chair/rides/{ride_id}/status", chairPostRideStatus)
	}

	// internal handlers
	{
		mux.HandleFunc("GET /api/internal/matching", internalGetMatching)
	}

	return mux
}

type postInitializeRequest struct {
	PaymentServer string `json:"payment_server"`
}

type postInitializeResponse struct {
	Language string `json:"language"`
}

func postInitialize(w http.ResponseWriter, r *http.Request) {
	go func() {
		if out, err := exec.Command("/home/isucon/local/golang/bin/go", "tool", "pprof", "-seconds=30", "-proto", "-output", "/home/isucon/pprof/pprof.pb.gz", "localhost:6060/debug/pprof/profile").CombinedOutput(); err != nil {
			fmt.Printf("pprof failed with err=%s, %s", string(out), err)
		} else {
			fmt.Printf("pprof.pb.gz created: %s", string(out))
		}
	}()
	go func() {
		if out, err := exec.Command("/home/isucon/local/golang/bin/go", "tool", "pprof", "-seconds=30", "-proto", "-output", "/home/isucon/pprof/fgprof.pb.gz", "localhost:6060/debug/fgprof").CombinedOutput(); err != nil {
			fmt.Printf("fgprof failed with err=%s, %s", string(out), err)
		} else {
			fmt.Printf("fgprof.pb.gz created: %s", string(out))
		}
	}()

	ctx := r.Context()
	req := &postInitializeRequest{}
	if err := bindJSON(r, req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if out, err := exec.Command("../sql/init.sh").CombinedOutput(); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Errorf("failed to initialize: %s: %w", string(out), err))
		return
	}

	if _, err := db.ExecContext(ctx, "UPDATE settings SET value = ? WHERE name = 'payment_gateway_url'", req.PaymentServer); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	LoadMap()

	writeJSON(w, http.StatusOK, postInitializeResponse{Language: "go"})
}

type Coordinate struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

func bindJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func writeJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	buf, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(buf)
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(statusCode)
	buf, marshalError := json.Marshal(map[string]string{"message": err.Error()})
	if marshalError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"marshaling error failed"}`))
		return
	}
	w.Write(buf)

	slog.Error("error response wrote", err)
}

func secureRandomStr(b int) string {
	k := make([]byte, b)
	if _, err := crand.Read(k); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", k)
}
