package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Go 有一个包名为 `context` 专门处理类似的场景
// 这个包定义了 Context 类型。它最初的目的是保存不同
// API 和不同处理的截止时间、取消信号和其他请求相关的值

// net/http 包中的的 Request 类型已经有 context 与之绑定。从 Go 1.7 开始
// Request 新增了一个返回请求的上下文的 `Context` 方法。对于进来的请求
// 在客户端关闭连接、请求被取消（HTTP/2 中）或 ServeHTTP 方法返回后，服务端会取消上下文

// 我们期望的现象是，当客户端取消请求（输入了 CTRL + C）或一段时间后 TimeoutHandler
// 继续执行然后终止请求时，服务端会停止后续的处理。进而关闭所有的连接，释放所有被运行中的
// 处理进程（及它的所有子协程）占用的资源

func slowAPICall(ctx context.Context) string {
	d := rand.Intn(5)
	select {
	case <-ctx.Done():
		log.Printf("slowAPICall was supposed to take %v seconds, but was canceled.", d)
		return ""
		// time.After() 可能会导致内存泄漏
	case <-time.After(time.Duration(d) * time.Second):
		log.Printf("Slow API call done after %v seconds.\n", d)
		return "foobar"
	}
}

func slowHandle(w http.ResponseWriter, req *http.Request) {
	result := slowAPICall(req.Context())
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

// Done 通道的关闭有效地取消了上下文，使我们的 slowAPICall 函数终止了它的执行
// 这就是 TimeoutHandler 终止耗时长的处理进程的原理
