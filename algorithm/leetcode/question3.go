package main

import (
	"fmt"
	"time"
)

// 两个线程轮流打印数字1-N

// N ...
var N = 10

func alternatePrint() {
	msg := make(chan int)
	go func(msg chan int) {
		for i := 1; i <= N; i++ {
			msg <- i
			if i&1 == 1 {
				fmt.Printf("G1: %d\n", i)
			}
		}
		close(msg)
	}(msg)

	go func(msg chan int) {
		for i := 1; i <= N; i++ {
			<-msg
			if i&1 == 0 {
				fmt.Printf("G2: %d\n", i)
			}
		}
	}(msg)

	time.Sleep(1 * time.Second)
}

func alternatePrint1() {
	chan1 := make(chan struct{})
	chan2 := make(chan struct{})

	go func() {
		for i := 1; i <= N; i += 2 {
			<-chan1
			fmt.Printf("G1: %d\n", i)
			chan2 <- struct{}{}
		}
	}()

	go func() {
		for i := 2; i <= N; i += 2 {
			<-chan2
			fmt.Printf("G2: %d\n", i)
			chan1 <- struct{}{}
		}
	}()

	// 通知G1开始执行,G1通知G2, G2通知G1
	chan1 <- struct{}{}
	time.Sleep(time.Second)

	close(chan1)
	close(chan2)
}
