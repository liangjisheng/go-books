package main

import "fmt"

func day001Test1() {
	s := make([]int, 3, 9)
	fmt.Println(len(s)) // 3
	s2 := s[4:8]
	fmt.Println(len(s2)) // 4
}

// N ...
type N int

func (n N) test() {
	fmt.Println(n)
}

// 当目标方法的接收者是指针类型时,那么被复制的就是指针
// 当目标方法的接收者是值类型时,那么被复制的就是值
// 当指针值赋值给变量或者作为函数参数传递时 会立即计算并复制该方法执行所需的
// 接收者对象 与其绑定 以便在稍后执行时 能隐式第传入接收者参数
func day001Test2() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()
}

func day001() {
	// day001Test1()
	day001Test2() // 13 11 12
}
