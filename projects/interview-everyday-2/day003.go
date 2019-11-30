package main

func day003Test1() {
	x := make([]int, 2, 10)
	_ = x[6:10]
	// panic 截取符号 [i:j] 如果 j 省略 默认是原切片或者数组的长度 x 的长度是 2
	// 小于起始下标 6 所以 panic
	_ = x[6:]
	_ = x[2:]
}

func day003() {
	// day003Test1()
}
