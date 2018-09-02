export CONNSTR="user=shupt password=shupt dbname=shupt host=127.0.0.1 sslmode=disable"
pg_ctl -D ./data -l logfile start
go run ./script/SyncDB/main.go