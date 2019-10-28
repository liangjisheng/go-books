package main

import "time"

// 但是管道的发送操作和接收操作是一一对应的，如果要停止多个Goroutine那么可能需要创建
// 同样数量的管道，这个代价太大了。其实我们可以通过close关闭一个管道来实现广播的效果
// 所有从关闭管道接收的操作均会收到一个零值和一个可选的失败标志

func exit2() {
	cancel := make(chan bool)

	for i := 0; i < 10; i++ {
		go worker(cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
}
