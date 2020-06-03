package main

import (
	"fmt"
	"sync"
	"time"
)

// Fibonacci ...
type Fibonacci struct {
	a, b int
	stop chan struct{}
	mtx  sync.Mutex
}

// NewFibonacci ...
func NewFibonacci() *Fibonacci {
	return &Fibonacci{
		a:    0,
		b:    1,
		stop: make(chan struct{}),
	}
}

// Run ...
func (f *Fibonacci) Run() {
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				f.mtx.Lock()
				fmt.Println(f.b)
				f.a, f.b = f.b, f.a+f.b
				f.mtx.Unlock()
			case <-f.stop:
				t.Stop()
				return
			}
		}
	}()
}

// Stop 调用 Stop 结束
func (f *Fibonacci) Stop() {
	close(f.stop)
}

// Value 获取当前的斐波那契数
func (f *Fibonacci) Value() int {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return f.b
}

func main() {
	f := NewFibonacci()
	f.Run()
	time.Sleep(5 * time.Second)
	f.Stop()

	// 用 select {} 代替 time.sleep 会导致死锁
	// select {}
	time.Sleep(2 * time.Second)
}
