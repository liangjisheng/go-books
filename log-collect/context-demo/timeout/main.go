package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Result ...
type Result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	// req, err := http.NewRequest("GET", "http://www.google.com", nil)
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed,err:", err)
		return
	}

	// 如果请求成功了会将数据存入到管道中
	go func() {
		resp, err := client.Do(req)
		pack := Result{resp, err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		fmt.Println("timeout!")
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("server response:%s", out)
	}
	return
}

func main() {
	process()
}
