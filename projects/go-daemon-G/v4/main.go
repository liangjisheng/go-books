package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Fibonacci ...
type Fibonacci struct {
	a, b int
	stop func()
	mtx  sync.Mutex
}

// NewFibonacci ...
func NewFibonacci() *Fibonacci {
	return &Fibonacci{a: 0, b: 1}
}

// Run ...
func (f *Fibonacci) Run(ctx context.Context) {
	// 使用WithCancel派生一个可被取消的ctx, 用来控制后台协程
	ctx, f.stop = context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("close chan")
				return
			case <-f.loop(ctx):
				// f.loop() 在正常运行时errch是阻塞状态，如果
				// 出错了才有数据，此时select会被唤起，并重新
				// 启动 loop()，实现panic后自动恢复。
			}
		}
	}()
}

func (f *Fibonacci) loop(ctx context.Context) <-chan error {
	errch := make(chan error)
	go func() {
		t := time.NewTicker(time.Second)
		defer func() {
			t.Stop()
			if r := recover(); r != nil {
				errch <- fmt.Errorf("panic with error %v", r)
				close(errch)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				close(errch)
				return
			case <-t.C:
				f.nextFibonacci()
			}
		}
	}()
	return errch
}

func (f *Fibonacci) nextFibonacci() {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	fmt.Println(f.b)
	f.a, f.b = f.b, f.a+f.b
}

// Stop 调用 Stop 结束
func (f *Fibonacci) Stop() {
	if f.stop != nil {
		f.stop()
	}
}

// Value 获取当前的斐波那契数
func (f *Fibonacci) Value() int {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return f.b
}

func main() {
	f := NewFibonacci()
	f.Run(context.Background())
	time.Sleep(5 * time.Second)
	f.Stop()

	time.Sleep(2 * time.Second)
}
