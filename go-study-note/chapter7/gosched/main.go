package main

import (
	"runtime"
	"sync"
	"time"
)

// 和协程 yield 作⽤类似， Gosched 让出底层线程，将当前 goroutine 暂停
// 放回队列等待下次被调度执⾏

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// 只需在函数调⽤语句前添加 go 关键字，就可创建并发执⾏单元
	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			println(i)
			time.Sleep(time.Millisecond)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}()

	time.Sleep(time.Millisecond * 4)
	go func() {
		wg.Done()
		println("hello, world")
	}()

	wg.Wait()
}
