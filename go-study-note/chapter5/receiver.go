package main

import "fmt"

type data struct {
	x int
}

func (d data) valueTest() { // func valueTest(self data)
	fmt.Printf("value: %p\n", &d)
}

func (d *data) pointTest() { // func pointTest(self *data)
	fmt.Printf("point: %p\n", d)
}

func receiverTest() {
	d := data{}
	p := &d
	fmt.Printf("data:  %p\n", p)

	d.valueTest()
	d.pointTest()

	p.valueTest()
	p.pointTest()
}
