package main

import "fmt"

// ConfigOne ...
type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	// Sprintf format %v with arg c causes recursive String method call
	return fmt.Sprintf("print: %v", c)
}

// 无限递归循环，栈溢出
// 如果类型定义了 String() 方法，使用 Printf() Print() Println() Sprintf()
// 等格式化输出时会自动使用 String() 方法
func day006Test1() {
	c := &ConfigOne{}
	c.String()
}

func day006() {
	day006Test1()
}
