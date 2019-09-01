package main

import (
	"math"
	"sync"
)

// 默认情况下，进程启动后仅允许⼀个系统线程服务于 goroutine
// 可使⽤环境变量或标准库函数 runtime.GOMAXPROCS 修改，
// 让调度器⽤多个线程实现多核并⾏，⽽不仅仅是并发
// go build -o test
// time -p ./test
// GOMAXPROCS=2 time -p ./test

func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}

	wg.Wait()
}
