package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 限流的要求是在指定的时间间隔内，server 最多只能服务指定数量的请求
// 实现的原理是我们启动一个计数器，每次服务请求会把计数器加一
// 同时到达指定的时间间隔后会把计数器清零

// RequestLimitService ...
type RequestLimitService struct {
	Interval time.Duration
	MaxCount int
	Lock     sync.Mutex
	ReqCount int
}

// NewRequestLimitService ...
func NewRequestLimitService(interval time.Duration, maxCnt int) *RequestLimitService {
	reqLimit := &RequestLimitService{
		Interval: interval,
		MaxCount: maxCnt,
	}

	go func() {
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			reqLimit.Lock.Lock()
			fmt.Println("Reset Count...")
			reqLimit.ReqCount = 0
			reqLimit.Lock.Unlock()
		}
	}()

	return reqLimit
}

// Increase ...
func (reqLimit *RequestLimitService) Increase() {
	reqLimit.Lock.Lock()
	defer reqLimit.Lock.Unlock()

	reqLimit.ReqCount++
}

// IsAvailable ...
func (reqLimit *RequestLimitService) IsAvailable() bool {
	reqLimit.Lock.Lock()
	defer reqLimit.Lock.Unlock()

	return reqLimit.ReqCount < reqLimit.MaxCount
}

// RequestLimit ...
var RequestLimit = NewRequestLimitService(10*time.Second, 5)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if RequestLimit.IsAvailable() {
		RequestLimit.Increase()
		fmt.Println(RequestLimit.ReqCount)
		io.WriteString(w, "Hello world!\n")
	} else {
		fmt.Println("Reach request limiting!")
		io.WriteString(w, "Reach request limit!\n")
	}
}

// curl http://127.0.0.1:8080/

func main() {
	fmt.Println("Server Started!")
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
