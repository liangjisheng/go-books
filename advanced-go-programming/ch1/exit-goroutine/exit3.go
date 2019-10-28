package main

import (
	"fmt"
	"sync"
	"time"
)

// 不过exit2()这个程序依然不够稳健：当每个Goroutine收到退出指令退出时一般会进行
// 一定的清理工作，但是退出的清理工作并不能保证被完成，因为main线程并没有等待各
// 个工作Goroutine退出工作完成的机制。我们可以结合sync.WaitGroup来改进

func worker1(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-cancel:
			return
		}
	}
}

// 现在每个工作者并发体的创建、运行、暂停和退出都是在exit3函数的安全控制之下了
func exit3() {
	cancel := make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker1(&wg, cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
