package main

import "fmt"

// Orange ...
type Orange struct {
	Quantity int
}

// Increase ...
func (o *Orange) Increase(n int) {
	o.Quantity += n
}

// Decrease ...
func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

// String ...
func (o *Orange) String() string {
	return fmt.Sprintf("point %#v", o.Quantity)
}

func day002Test1() {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	// String() 是指针方法，而不是值方法，所以使用 Println() 输出时不会调用到 String() 方法
	fmt.Println(orange)

	o1 := &Orange{}
	o1.Increase(10)
	o1.Decrease(1)
	fmt.Println(o1)
}

func day002() {
	day002Test1()
}
