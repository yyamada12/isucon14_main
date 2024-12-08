# DB Graph
node: ![](https://via.placeholder.com/16/795548/FFFFFF/?text=%20) `テーブル` ![](https://via.placeholder.com/16/1976D2/FFFFFF/?text=%20) `関数` 

edge: ![](https://via.placeholder.com/16/CDDC39/FFFFFF/?text=%20) `INSERT` ![](https://via.placeholder.com/16/FF9800/FFFFFF/?text=%20) `UPDATE` ![](https://via.placeholder.com/16/F44336/FFFFFF/?text=%20) `DELETE` ![](https://via.placeholder.com/16/78909C/FFFFFF/?text=%20) `SELECT` ![](https://via.placeholder.com/16/BBDEFB/FFFFFF/?text=%20) `関数呼び出し` 
```mermaid
graph LR
  classDef table fill:#795548,fill-opacity:0.5
  classDef func fill:#1976D2,fill-opacity:0.5
  func:github.com/isucon/isucon14/webapp/go.ownerPostOwners[ownerPostOwners]:::func --> table:owners[owners]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> func:github.com/isucon/isucon14/webapp/go.calculateDiscountedFare[calculateDiscountedFare]:::func
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func ==> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRides[appPostRides]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.chairGetNotification[chairGetNotification]:::func --> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.chairGetNotification[chairGetNotification]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.chairGetNotification[chairGetNotification]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.chairGetNotification[chairGetNotification]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.chairGetNotification[chairGetNotification]:::func --> table:users[users]:::table
  func:github.com/isucon/isucon14/webapp/go.ownerGetSales[ownerGetSales]:::func --> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.ownerGetSales[ownerGetSales]:::func ==> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.ownerGetSales[ownerGetSales]:::func ==> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostRideStatus[chairPostRideStatus]:::func --> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.chairPostRideStatus[chairPostRideStatus]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostRideStatus[chairPostRideStatus]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appAuthMiddleware[appAuthMiddleware]:::func --> table:users[users]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> func:github.com/isucon/isucon14/webapp/go.calculateDiscountedFare[calculateDiscountedFare]:::func
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> table:payment_tokens[payment_tokens]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRideEvaluatation[appPostRideEvaluatation]:::func --> table:settings[settings]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetNearbyChairs[appGetNearbyChairs]:::func ==> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.appGetNearbyChairs[appGetNearbyChairs]:::func ==> table:chair_locations[chair_locations]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetNearbyChairs[appGetNearbyChairs]:::func --> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetNearbyChairs[appGetNearbyChairs]:::func ==> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.internalGetMatching[internalGetMatching]:::func ==> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.internalGetMatching[internalGetMatching]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.internalGetMatching[internalGetMatching]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.getChairStats[getChairStats]:::func ==> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.getChairStats[getChairStats]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.ownerAuthMiddleware[ownerAuthMiddleware]:::func --> table:owners[owners]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostActivity[chairPostActivity]:::func --> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.chairAuthMiddleware[chairAuthMiddleware]:::func --> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostChairs[chairPostChairs]:::func --> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostChairs[chairPostChairs]:::func --> table:owners[owners]:::table
  func:github.com/isucon/isucon14/webapp/go.ownerGetChairs[ownerGetChairs]:::func --> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.calculateDiscountedFare[calculateDiscountedFare]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostCoordinate[chairPostCoordinate]:::func --> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.chairPostCoordinate[chairPostCoordinate]:::func --> table:chair_locations[chair_locations]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostCoordinate[chairPostCoordinate]:::func --> table:chair_locations[chair_locations]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostCoordinate[chairPostCoordinate]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.chairPostCoordinate[chairPostCoordinate]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetRides[appGetRides]:::func ==> func:github.com/isucon/isucon14/webapp/go.calculateDiscountedFare[calculateDiscountedFare]:::func
  func:github.com/isucon/isucon14/webapp/go.appGetRides[appGetRides]:::func ==> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.appGetRides[appGetRides]:::func ==> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetRides[appGetRides]:::func ==> table:owners[owners]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetRides[appGetRides]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostUsers[appPostUsers]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostUsers[appPostUsers]:::func --> table:coupons[coupons]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostUsers[appPostUsers]:::func --> table:users[users]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostUsers[appPostUsers]:::func --> table:users[users]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostRidesEstimatedFare[appPostRidesEstimatedFare]:::func --> func:github.com/isucon/isucon14/webapp/go.calculateDiscountedFare[calculateDiscountedFare]:::func
  func:github.com/isucon/isucon14/webapp/go.appGetNotification[appGetNotification]:::func --> func:github.com/isucon/isucon14/webapp/go.calculateDiscountedFare[calculateDiscountedFare]:::func
  func:github.com/isucon/isucon14/webapp/go.appGetNotification[appGetNotification]:::func --> func:github.com/isucon/isucon14/webapp/go.getChairStats[getChairStats]:::func
  func:github.com/isucon/isucon14/webapp/go.appGetNotification[appGetNotification]:::func --> func:github.com/isucon/isucon14/webapp/go.getLatestRideStatus[getLatestRideStatus]:::func
  func:github.com/isucon/isucon14/webapp/go.appGetNotification[appGetNotification]:::func --> table:chairs[chairs]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetNotification[appGetNotification]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetNotification[appGetNotification]:::func --> table:ride_statuses[ride_statuses]:::table
  func:github.com/isucon/isucon14/webapp/go.appGetNotification[appGetNotification]:::func --> table:rides[rides]:::table
  func:github.com/isucon/isucon14/webapp/go.appPostPaymentMethods[appPostPaymentMethods]:::func --> table:payment_tokens[payment_tokens]:::table
  linkStyle 1,2,12,20,24,25,32,49,54,55,63,64,65,66 stroke:#BBDEFB,stroke-width:2px
  linkStyle 3,5,7,13,29,37,42,69 stroke:#FF9800,stroke-width:2px
  linkStyle 4,6,9,11,14,15,16,17,18,19,22,23,26,28,30,31,33,34,35,36,38,39,40,41,43,45,46,47,48,50,53,56,57,58,59,62,67,68,70 stroke:#78909C,stroke-width:2px
  linkStyle 0,8,10,21,27,44,51,52,60,61,71 stroke:#CDDC39,stroke-width:2px

```