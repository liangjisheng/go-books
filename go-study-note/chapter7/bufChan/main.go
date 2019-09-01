package main

import "fmt"

func main() {
	data := make(chan int, 3)
	exit := make(chan bool)

	data <- 1
	data <- 2
	data <- 3

	go func() {
		// 除⽤ range 外，还可⽤ ok-idiom 模式判断 channel 是否关闭
		// for d := range data {
		// 	fmt.Println(d)
		// }

		for {
			if d, ok := <-data; ok {
				fmt.Println(d)
			} else {
				break
			}
		}

		exit <- true
	}()

	data <- 4
	data <- 5
	close(data)

	<-exit
}
