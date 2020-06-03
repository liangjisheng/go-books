package main

import (
	"fmt"
	"sync"
	"time"
)

// Fibonacci ...
type Fibonacci struct {
	a, b int
	stop bool
	mtx  sync.Mutex
}

// NewFibonacci ...
func NewFibonacci() *Fibonacci {
	return &Fibonacci{a: 0, b: 1}
}

// Run ...
func (f *Fibonacci) Run() {
	go func() {
		for {
			if f.isStop() {
				break
			}
			time.Sleep(time.Second)
			f.mtx.Lock()
			fmt.Println(f.b)
			f.a, f.b = f.b, f.a+f.b
			f.mtx.Unlock()
		}
	}()
}

// Stop 调用 Stop 结束
func (f *Fibonacci) Stop() {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	f.stop = true
}

func (f *Fibonacci) isStop() bool {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return f.stop
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
	fmt.Println(f.isStop())
	time.Sleep(5 * time.Second)
	f.Stop()

	fmt.Println(f.isStop())
	// 用 select {} 代替 time.sleep 会导致死锁
	time.Sleep(2 * time.Second)
}
