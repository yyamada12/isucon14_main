# git
alias g='git'
alias ga='git add'
alias gd='git diff'
alias gdc='git diff --cached'
alias gs='git status'
alias gp='git push'
alias gb='git branch'
alias gst='git status'
alias gco='git checkout'
alias gf='git fetch'
alias gci='git commit'
alias gcia='git commit --amend'
alias gl='git log'
alias ggr="git log --all --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit --date=relative"
alias gsr='gst && echo && ggr'
alias gres='git reset'

# shell
alias lh='ls -lah'
alias rm='rm -i'
alias cp='cp -i'
alias mv='mv -i'

# dstat
alias dstata='dstat -tlcmgdr --socket --tcp -n'

# tail
alias tlf='tail -F'

# watch
alias watch='watch '

# alp
alias al='sudo alp ltsv -c ~/alp.yml'
alias al_bak='sudo alp ltsv -c ~/alp.yml --file /var/log/nginx/access_bak.log'
# alias als='alp ltsv -c ~/alp.yml | slackcat -t -c cancer_acropolis -n alp.txt'
alias als='sudo alp ltsv -c ~/alp.yml > alp-result.txt && ~/upload_file_slack.sh alp-result.txt isucon && cat alp-result.txt && rm -f alp-result.txt'
alias als_bak='sudo alp ltsv -c ~/alp.yml --file /var/log/nginx/access_bak.log > alp-result.txt && ~/upload_file_slack.sh alp-result.txt isucon && cat alp-result.txt && rm -f alp-result.txt'

# pt-query-digest
alias pt='sudo pt-query-digest --limit 10 --report-format profile,query_report /var/log/mysql/slow.log | less'
alias pt_bak='sudo pt-query-digest --limit 10 --report-format profile,query_report /var/log/mysql/slow_bak.log | less'
# alias pts='sudo pt-query-digest --limit 10 --report-format profile,query_report /var/log/mysql/slow.log | slackcat -c cancer_acropolis -n slowlog.txt'
alias pts='sudo pt-query-digest --limit 10 --report-format profile,query_report /var/log/mysql/slow.log ~/alp.yml > pt-result.txt && ~/upload_file_slack.sh pt-result.txt isucon && cat pt-result.txt && rm -f pt-result.txt'
alias pts_bak='sudo pt-query-digest --limit 10 --report-format profile,query_report /var/log/mysql/slow_bak.log ~/alp.yml > pt-result.txt && ~/upload_file_slack.sh pt-result.txt isucon && cat pt-result.txt && rm -f pt-result.txt'

# pprof
alias pp='go tool pprof -http=":1234" ~/pprof/pprof.pb.gz'
alias pp_bak='go tool pprof -http=":1234" ~/pprof/pprof_bak.pb.gz'
# alias pps='go tool pprof -png -output ~/pprof/pprof.png http://localhost:6060/debug/pprof/profile && slackcat -c cancer_acropolis -n pprof.png ~/pprof/pprof.png'
alias pps='go tool pprof -png -output pprof.png ~/pprof/pprof.pb.gz && ~/upload_file_slack.sh pprof.png isucon && rm -f pprof.png'
alias pps_bak='go tool pprof -png -output pprof.png ~/pprof/pprof_bak.pb.gz && ~/upload_file_slack.sh pprof.png isucon && rm -f pprof.png'
alias ppb='go tool pprof -http=":1234" http://localhost:6060/debug/pprof/profile'

alias fgp='go tool pprof -http=":1235" ~/pprof/fgprof.pb.gz'
alias fgp_bak='go tool pprof -http=":1235" ~/pprof/fgprof_bak.pb.gz'

# app
alias deploy='~/deploy.sh'
alias applog='sudo journalctl -u $APP_SERVICE_NAME'

# systemctl
alias sc='sudo systemctl'
alias scl='sudo systemctl list-unit-files --type=service'
alias scla='sudo systemctl list-units --type=service --state=running'
alias scs='sudo systemctl status'
alias scr='sudo systemctl restart'
alias scsn='sudo systemctl status nginx'
alias scrn='sudo systemctl restart nginx'
alias scsm='sudo systemctl status mysql'
alias scrm='sudo systemctl restart mysql'
alias scss='sudo systemctl status $APP_SERVICE_NAME'
alias scrs='sudo systemctl restart $APP_SERVICE_NAME'
