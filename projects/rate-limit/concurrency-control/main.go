package main

// 利用 channel 的缓冲设定，我们就可以来实现并发的限制。我们只要在执行并发的同时
// 往一个带有缓冲的 channel 里写入点东西(随便写啥，内容不重要) 让并发的 goroutine
// 在执行完成后把这个 channel 里的东西给读走。这样整个并发的数量就讲控制在这个
// channel的缓冲区大小上

import (
	"fmt"
	"time"
)

// Run ...
func Run(taskID, sleeptime, timeout int, ch chan string) {
	chRun := make(chan string)
	go run(taskID, sleeptime, chRun)
	select {
	case re := <-chRun:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", taskID)
		ch <- re
	}
}

func run(taskID, sleeptime int, ch chan string) {
	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", taskID, sleeptime)
}

func main() {
	input := []int{3, 2, 1}
	timeout := 2
	chLimit := make(chan bool, 1)
	// chLimit := make(chan bool, 2) // 修改并发限制为2
	chs := make([]chan string, len(input))
	limitFunc := func(chLimit chan bool, ch chan string, taskID, sleeptime, timeout int) {
		Run(taskID, sleeptime, timeout, ch)
		<-chLimit
	}
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string, 1)
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, sleeptime, timeout)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}
