package main

const cl = 100

var bl = 123

func consttest() {
	println(&bl, bl)
	// 常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用
	// println(&cl, cl)
}
