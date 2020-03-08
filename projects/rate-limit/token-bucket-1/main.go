package main

import (
	"fmt"
	"net"
	"os"
	"sync/atomic"
	"time"
)

var (
	limiting int32 = 10 // 这就是我的令牌桶
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr) // 监听一个端口
	checkError(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept() // 在此处阻塞，每次来一个请求才往下运行handle函数
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 起一个单独的协程处理，有多少个请求，就起多少个协程，协程之间共享同一个全局变量
		// limiting 对其进行原子操作
		go handle(&conn)
	}
}

func handle(conn *net.Conn) {
	defer (*conn).Close()
	// dcr 1 by atomic，获取一个令牌，总数减1 这是一个原子性的操作
	// 并发情况下，数据不会写错
	n := atomic.AddInt32(&limiting, -1)
	if n < 0 {
		// 令牌不够用了，限流，抛弃此次请求
		(*conn).Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\nError, too many request, please try again."))
	} else {
		// 还有剩余令牌可用
		time.Sleep(1 * time.Second) // 假设我们的应用处理业务用了1s的时间
		// 业务处理结束后，回复200成功
		(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!"))
	}
	atomic.AddInt32(&limiting, 1) // add 1 by atomic，业务处理完毕，放回令牌
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
