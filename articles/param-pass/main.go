package main

// go就是值传递，可以确认的是Go语言中所有的传参都是值传递（传值）
// 都是一个副本, 一个拷贝. 因为拷贝的内容有时候是非引用类型
// (int、string、struct等这些), 这样就在函数中就无法修改原内容数据
// 有的是引用类型(指针、map、slice、chan等这些)这样就可以修改原内容数据

// 是否可以修改原内容数据，和传值、传引用没有必然的关系。在C++中
// 传引用肯定是可以修改原内容数据的，在Go语言里，虽然只有传值
// 但是我们也可以修改原内容数据，因为参数是引用类型

// 引用类型和传引用是2个概念

func main() {
	// sliceParam()
	// mapParam()
	chanParam()
}
