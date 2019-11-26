package main

// 不能使用短变量声明设置结构体字段值

import "fmt"

// ConfigOne ...
type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

// 调用 day004Test1 会运行时错误 如果类型实现 String() 方法，当格式化输出时
// 会自动使用 String() 方法 上面这段代码是在该类型的 String() 方法内使用格式
// 化输出 导致递归调用 最后抛错
func day004Test1() {
	c := &ConfigOne{}
	c.String()
}

func day004Test2() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c)
	// Output
	// false true false

	// 不能对 nil 的 map 直接赋值，需要使用 make() 初始化
	// 但可以使用 append() 函数对为 nil 的 slice 增加元素
	var s []int
	s = append(s, 1)
	fmt.Println(s)
}

func day004() {
	// day004Test1()
	day004Test2()
}
