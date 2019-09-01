package main

var x, y, z int
var s, n = "abc", 123

// 编译器会将未使用的局部变量当做错误，但全局变量没有问题
var (
	a int
	b float32
)

func main() {
	n, s := 0x1234, "Hello, World!"
	println(x, s, n)
}
