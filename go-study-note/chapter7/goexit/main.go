package main

import (
	"runtime"
	"sync"
)

// 调⽤ runtime.Goexit 将⽴即终⽌当前 goroutine 执⾏
// 调度器确保所有已注册 defer 延迟调⽤被执⾏

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer println("A.defer")

		func() {
			defer println("B.defer")
			runtime.Goexit()
			println("B") // 不会执行
		}()

		println("A") // 不会执行
	}()

	wg.Wait()
}

// output
// B.defer
// A.defer
