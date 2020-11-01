package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// 我们的目标是确保我们的 slowHandler 在 1s 内完成处理。如果超过了 1s
// 我们的服务会停止运行并返回对应的超时错误
// 基于此, net/http 包提供了`TimeoutHandler`— 返回了一个在给定的时间限制内运行的 handler
// func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
// 第一个参数是 Handler, 第二个参数是 time.Duration(超时时间), 第三个参数是 string 类型
// 当到达超时时间后返回的信息

func slowHandle(w http.ResponseWriter, req *http.Request) {
	time.Sleep(2 * time.Second)
	io.WriteString(w, "I am slow!\n")
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

// 现在我们的程序不会有过长耗时的处理；也避免了有人恶意发送导致长耗时处理的请求时
// 导致的潜在的 DoS 攻击
// 我们的 slowHandler 仅仅是个简单的 demo. 但是，如果我们的程序复杂些，
// 能向其他服务和资源发出请求会发生什么呢？如果我们的程序在超时时向诸如
// S3 的服务发出了请求会怎么样? 会发生什么?
