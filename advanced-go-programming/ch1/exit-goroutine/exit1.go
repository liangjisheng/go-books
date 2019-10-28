package main

import (
	"fmt"
	"time"
)

// 通过select和default分支可以很容易实现一个Goroutine的退出控制
func worker(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello") // 正常工作
		case <-cancel:
			return // 退出
		}
	}
}

func exit1() {
	cancel := make(chan bool)
	go worker(cancel)

	time.Sleep(time.Second)
	cancel <- true
}
