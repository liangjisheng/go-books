package main

import "fmt"

func day008_1(i int) {
	fmt.Println(i)
}

// People ...
type People struct{}

// ShowA ...
func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

// ShowB ...
func (p *People) ShowB() {
	fmt.Println("showB")
}

// Teacher ...
type Teacher struct {
	People
}

// ShowB ...
func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func day008() {
	i := 5
	defer day008_1(i)
	i = i + 10
	// Output
	// 5

	t := Teacher{}
	t.ShowA()
	// Output
	// showA
	// showB
}
