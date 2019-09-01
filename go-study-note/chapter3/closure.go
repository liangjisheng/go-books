package main

import "fmt"

// 闭包复制的是原对象指针
// 在汇编层⾯， closure 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、
// 闭包对象指针。当调⽤匿名函数时，只需以某个寄存器传递该对象即可.
func closure() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func testClosure() {
	f := closure()
	f()
}
