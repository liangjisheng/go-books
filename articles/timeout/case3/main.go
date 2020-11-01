package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 稍微展开下我们的例子

func slowAPICall() string {
	d := rand.Intn(5)
	select {
	case <-time.After(time.Duration(d) * time.Second):
		log.Printf("Slow API call done after %v seconds.\n", d)
		return "foobar"
	}
}

func slowHandle(w http.ResponseWriter, req *http.Request) {
	result := slowAPICall()
	io.WriteString(w, result+"\n")
}

func main() {
	srv := http.Server{
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 5 * time.Second,
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandle), 1*time.Second, "Timeout!\n"),
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

// 通过观察服务的输出，我们会发现，它是在几秒之后打出日志的，而不是在超时 handler 生效时打出
// 这个现象表明：虽然 1 秒之后请求超时了，但是服务仍然完整地处理了请求
// 这就是在几秒之后才打出日志的原因

// 虽然在这个例子里问题很简单，但是类似的现象在生产中可能变成一个严重的问题。例如
// 当 slowAPICall 函数开启了几个百个协程，每个协程都处理一些数据时。或者当它向
// 不同系统发出多个不同的 API 发出请求时。这种耗时长的的进程，它们的请求方/客户端
// 并不会使用服务端的返回结果，会耗尽你系统的资源。
// 所以，我们怎么保护系统，使之不会出现类似的未优化的超时或取消请求呢？
