package main

// 某些时候，让函数直接 "实现" 接⼝能省不少事

type tester1 interface {
	Do()
}

type funcDo func()

func (f funcDo) Do() {
	f()
}

func funcInterface() {
	var t tester1 = funcDo(func() { println("hello, world") })
	t.Do()
}
