package main

import (
	"fmt"
	"os"
	"runtime"
)

func day004Test1() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()

	// for {} 独占 CPU 资源导致其他 Goroutine 饿死
	// for {
	// }

	// 可以通过阻塞的方式避免 CPU 占用，修复代码
	select {}
}

func day004() {
	day004Test1()
}
