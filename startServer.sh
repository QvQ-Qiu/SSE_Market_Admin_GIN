# 启动原本的文件
./serve &

# 启动 cron.go 文件
go run ./cron/cron.go

# 等待进程结束
wait
