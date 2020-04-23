package main

import (
	"fmt"
	"net"
	"net/http"

	"golang.org/x/net/netutil"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.Body=", r.Body)
		fmt.Fprintf(w, "hello world")
	})

	server := &http.Server{
		Addr:    "",
		Handler: mux,
	}
	listener, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("服务器错误")
	}

	// 服务端设置了 limit 为 1 的时候，意味这我们服务端的最大连接数只能为 1
	// 我们客户端在使用 http keepalive，会将这个连接占满，如果这个时候又进行了 tcp 探活
	// 那么这个探活的请求就会被堵到 backlog 里
	err = server.Serve(netutil.LimitListener(listener, 1))
	if err != nil {
		fmt.Println("服务器错误2")
	}
}
