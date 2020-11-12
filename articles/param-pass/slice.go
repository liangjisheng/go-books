package main

import "fmt"

func sliceParam() {
	// 之所以对于引用类型的传递可以修改原内容的数据，这是因为在底层默认使用
	// 该引用类型的指针进行传递，但也是使用指针的副本，依旧是值传递。所以
	// slice传递的就是第一个元素的指针的副本，因为fmt.printf缘故造成了打印
	// 的地址一样，给人一种混淆的感觉
	var args = []int64{1, 2, 3}
	fmt.Printf("切片args的地址: %p\n", args)
	fmt.Printf("切片args第一个元素的地址: %p \n", &args[0])
	fmt.Printf("直接对切片args取地址%v \n", &args)
	modifiedNumber(args)
	fmt.Println(args)
}

func modifiedNumber(args []int64) {
	fmt.Printf("形参切片的地址: %p \n", args)
	fmt.Printf("形参切片args第一个元素的地址: %p \n", &args[0])
	fmt.Printf("直接对形参切片args取地址%v \n", &args)
	args[0] = 10
}
