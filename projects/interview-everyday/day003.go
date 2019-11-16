package main

import "fmt"

// 相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关

func day003() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	fmt.Println("sm1:", sm1)
	fmt.Println("sm2:", sm2)

	// invalid operation: sm1 == sm2
	// (struct containing map[string]string cannot be compared)
	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }
}
