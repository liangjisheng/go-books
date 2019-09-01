package main

import (
	"fmt"
	"reflect"
)

func demo1() {
	var u Admin
	t := reflect.TypeOf(u)

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

// 如果是指针，应该先使⽤ Elem 方法获取目标类型，指针本⾝身是没有字段成员的
func demo2() {
	u := new(Admin)

	t := reflect.TypeOf(u)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

// value-interface 和 pointer-interface 也会导致方法集存在差异
func demo3() {
	var u Admin

	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}

	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(u))

	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&u))
}

// 可从基本类型获取所对应复合类型
func demo4() {
	var (
		Int    = reflect.TypeOf(0)
		String = reflect.TypeOf("")
	)

	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)

	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)

	t := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p)
}
