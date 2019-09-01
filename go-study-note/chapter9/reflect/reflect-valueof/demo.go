package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Value 和 Type 使⽤方法类似，包括使用 Elem 获取指针目标对象
func demo1() {
	u := &Admin{User{"ljs", 24}, "NT"}
	v := reflect.ValueOf(u).Elem()

	fmt.Println(v.FieldByName("title").String())   // ⽤转换⽅方法获取字段值
	fmt.Println(v.FieldByName("age").Int())        // 直接访问嵌入字段成员
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int()) // 用多级序号访问嵌⼊入字段成员
}

func demo2() {
	v := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, v.Len(); i < n; i++ {
		fmt.Println(v.Index(i).Int())
	}

	fmt.Println("--------------------------")

	v = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	for _, k := range v.MapKeys() {
		fmt.Println(k.String(), v.MapIndex(k).Int())
	}
}

// 将对象转换为接⼝，会发生复制行为,该复制品只读，无法被修改。所以要通过接口改变
// 目标对象状态，必须是 pointer-interface
func demo3() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())
}

// 非导出字段无法直接修改，可改用指针操作
func demo4() {
	u := User{"ljs", 23}
	p := reflect.ValueOf(&u).Elem()

	p.FieldByName("Username").SetString("Tom")

	f := p.FieldByName("age")
	fmt.Println(f.CanSet())

	// 判断是否能获取地址
	if f.CanAddr() {
		age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		// age := (*int)(unsafe.Pointer(f.Addr().Pointer())) // 等同
		*age = 88
	}

	// p 是 Value 类型,需要还原成接口才能转型
	fmt.Println(u, p.Interface().(User))
}

// 复合类型修改示例
func demo5() {
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()

	v.SetLen(2)
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)

	fmt.Println(v.Interface(), s)

	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))

	fmt.Println(v2.Interface())
	fmt.Println("------------------------------")

	m := map[string]int{"a": 1}
	v = reflect.ValueOf(&m).Elem()

	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200))

	fmt.Println(v.Interface(), m)
}
