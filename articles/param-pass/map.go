package main

import "fmt"

func mapParam() {
	persons := make(map[string]int)
	persons["asong"] = 8

	addr := &persons

	fmt.Printf("原始map的内存地址是: %p\n", addr)
	// 使用make函数返回的是一个hmap类型的指针*hmap.
	// func modifiedAge(person map[string]int)函数，其实就等于
	// func modifiedAge(person *hmap）,实际上在作为传递参数时还是
	// 使用了指针的副本进行传递, 属于值传递. 在这里, Go语言通过make函数
	// 字面量的包装，为我们省去了指针的操作，让我们可以更容易的使用map
	// 这里的map可以理解为引用类型，但是记住引用类型不是传引用
	modifiedAge(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}

func modifiedAge(person map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是: %p\n", &person)
	person["asong"] = 9
}
