package main

import "fmt"

// People ...
type People interface {
	SpeakValue(string) string
	SpeakPoint(string) string
}

// Student ...
type Student struct{}

// SpeakValue 值接收器实现的方法默认指针接收器也会实现,因为不会对接受对象造成改变
// 但指针接收器实现的方法默认值接收器不会实现,因为指针可能改变对象
func (Student) SpeakValue(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

// SpeakPoint ...
func (*Student) SpeakPoint(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func interface1() {
	think := "speak"

	// var peoValue People = Student{}
	// fmt.Println(peoValue.SpeakValue(think))

	var peoPoint People = &Student{}
	fmt.Println(peoPoint.SpeakValue(think))
}
