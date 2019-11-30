package main

import "fmt"

func day002Test1() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	// panic 因为两个比较值的动态类型为同一个不可比较类型
	_ = y == y
}

// 第一次循环,写操作已经准备好,打印3, 第二次,读操作准备好,打印2
// 第三次,由于 c 为 nil,走的是 default 分支,输出 1
func day002Test2() {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			fmt.Println(1)
		case <-c:
			fmt.Println(2)
			c = nil
		case c <- 1:
			fmt.Println(3)
		}
	}
}

func day002() {
	// day002Test1()
	day002Test2()
}
