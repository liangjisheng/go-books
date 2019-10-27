package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 生产者: 生成 factor 整数倍的序列
func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64) // 成果队列

	go producer(3, ch) // 生成 3 的倍数的序列
	go producer(5, ch) // 生成 5 的倍数的序列
	go consumer(ch)    // 消费 生成的队列

	// 运行一定时间后退出 这种靠休眠方式是无法保证稳定的输出结果的
	// time.Sleep(2 * time.Second)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
