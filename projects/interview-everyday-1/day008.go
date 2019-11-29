package main

import "fmt"

// T8 ...
type T8 struct {
	n int
}

// Set ...
func (t *T8) Set(n int) {
	t.n = n
}

func getT8() T8 {
	return T8{}
}

func day008Test1() {
	// 直接返回的 T{} 不可寻址
	// 不可寻址的结构体不能调用带结构体指针接收者的方法
	// getT8().Set(1)

	t := getT8()
	t.Set(2)
	fmt.Println(t.n)
}

func day008() {
	day008Test1()
}
