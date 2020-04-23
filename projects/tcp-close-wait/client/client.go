package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var httpclient *http.Client

func main() {
	go func() {
		for i := 0; i < 1000; i++ {
			req, err := http.NewRequest("POST", "http://127.0.0.1:8080/hello", strings.NewReader("{}"))
			if err != nil {
				fmt.Println(err)
				return
			}
			// http请求结束后是否关闭连接, false表示不关闭,使用长连接
			req.Close = false

			httpclient = &http.Client{}
			resp, err := httpclient.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(string(body))
				fmt.Println(err)
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for i := 0; i < 2000; i++ {
		// time.Sleep(10 * time.Second)
		Dial("127.0.0.1:8080", time.Second)
		// time.Sleep(100 * time.Second)
	}
}

// Dial dial 指定的端口, 失败返回error
func Dial(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		fmt.Println("dial err------>", err)
		return err
	}
	defer conn.Close()
	return err
}
