#!/bin/bash

# trap 命令用于在 shell 脚本退出时 删掉临时文件 结束子进程
trap "rm server.exe; kill 0" EXIT

go build -o server.exe
./server.exe -port=8081 &
./server.exe -port=8082 &
./server.exe -port=8083 -api=1 &

sleep 2

echo ">>> start test"
curl "http://localhost:8080/api?key=Tom" &
curl "http://localhost:8080/api?key=Tom" &
curl "http://localhost:8080/api?key=Tom" &

wait
