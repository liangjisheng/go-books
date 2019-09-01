package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 只有 tab 和 data 都为 nil 时，接⼝才等于 nil
func interfaceNil() {
	var a interface{}               // tab = nil, data = nil
	var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil

	type iface struct {
		itab, data uintptr
	}

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println(a == nil, ia)                             // true {0 0}
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil()) // false {4796512 0} true
}
