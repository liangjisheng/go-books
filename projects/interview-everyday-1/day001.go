package main

import "fmt"

// Math ...
type Math struct {
	x, y int
}

func day001Test1() {
	var m = map[string]Math{
		"foo": Math{2, 3},
	}
	// cannot assign to struct field m["foo"].x in map
	// 对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X
	// 但 go 中的 map 的 value 本身是不可寻址的
	// m["foo"].x = 4
	fmt.Println(m["foo"].x)

	// 解决方法1 使用临时变量
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	fmt.Println(m["foo"].x)

	// 解决方法2 修改数据结构
	var m1 = map[string]*Math{
		"foo": &Math{2, 3},
	}
	fmt.Println(m1["foo"].x)

	m1["foo"].x = 4
	fmt.Println(m1["foo"].x)
}

func day001() {
	day001Test1()
}
