initdb ./data
pg_ctl -D ./data -l logfile start
createuser shupt -P
createdb shupt -O shupt -E UTF8 -e