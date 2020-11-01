package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// 尽管我们可以使用 deadline 来模拟超时操作，但我们还是不能控制处理器完成操作所需的
// 耗时 deadline 作用于连接，因此我们的服务仅在处理器尝试访问连接的属性
// (如对 http.ResponseWriter 进行写操作)之后才会返回(错误)结果

func slowHandle(w http.ResponseWriter, req *http.Request) {
	time.Sleep(2 * time.Second)
	io.WriteString(w, "I am slow!\n")
}

func main() {
	srv := http.Server{
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 1 * time.Second,
		Handler:      http.HandlerFunc(slowHandle),
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

// 虽然这是个类似超时的处理，但它更大的作用是在到达超时时间时，阻止服务进行更多的操作
// 结束请求。在我们上面的例子中，handler 在完成之前一直在处理请求
// 即使已经超出响应写超时时间(1 秒)100%(耗时 2 秒)
// 最根本的问题是, 对于处理器来说, 我们应该怎么设置超时时间才更有效?
