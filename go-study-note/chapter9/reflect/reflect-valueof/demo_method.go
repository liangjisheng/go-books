package main

import (
	"fmt"
	"reflect"
)

// Data ...
type Data struct{}

// Test ...
func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

// Sum ...
func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}

	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name)

	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf("  in[%d] %v\n", i, t.In(i))
	}

	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf("  out[%d] %v\n", i, t.Out(i))
	}
}

func demo6() {
	d := new(Data)
	t := reflect.TypeOf(d)

	test, _ := t.MethodByName("Test")
	info(test)

	sum, _ := t.MethodByName("Sum")
	info(sum)
}

func callMethod() {
	d := new(Data)
	v := reflect.ValueOf(d)

	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		out := m.Call(in)

		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}

	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})

	fmt.Println("--------------------------")

	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
}
