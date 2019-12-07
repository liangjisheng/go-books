package main

import "fmt"

// Named Type 有两类
// 内置类型，比如 int, int64, float, string, bool 等
// 使用关键字 type 声明的类型

// Unnamed Type 是基于已有的 Named Type 组合一起的类型
// 例如 struct{}、[]string、interface{}、map[string]bool

// T ...
type T []int

// F ...
func F(t T) {}

// Unnamed Type 不能作为方法的接收者
// func (m map[string]string) Set(key string, value string) {
// 	m[key] = value
// }

// User ...
type User map[string]string

// Set ...
func (m User) Set(key string, value string) {
	m[key] = value
}

func day001() {
	// 对于底层类型相同的变量可以相互赋值还有一个重要的条件
	// 即至少有一个不是有名类型 (named type), 因此可以编译通过
	var q []int
	F(q)

	m := make(User)
	m.Set("A", "One")
	fmt.Println(m)

	min(1, 2)
}

func min(a int, b uint) {
	var min = 0
	// 一行代码求两个数的最小值
	// 利用 copy() 函数的功能:切片复制，并且返回两者长度的较小值
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}
