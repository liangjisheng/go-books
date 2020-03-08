package main

import (
	"fmt"
	"net"
	"os"
	"time"

	cache "github.com/UncleBig/goCache"
)

var (
	limit = 10
	c     *cache.Cache
)

func main() {
	c = cache.New(10*time.Minute, 30*time.Second)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(&conn)
	}
}

func handle(conn *net.Conn) {
	defer (*conn).Close()
	t := time.Now().Unix()
	key := fmt.Sprintf("%d", t)
	if n, found := c.Get(key); found {
		num := n.(int)
		fmt.Printf("key:%d num:%d\n", t, num)
		if num >= limit {
			(*conn).Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\nError, too many request, please try again."))
		} else {
			(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!"))
			c.Increment(key, 1)
		}
	} else {
		(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!"))
		c.Set(key, 1, 2*time.Second)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
