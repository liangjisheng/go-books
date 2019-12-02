package main

// 闭包引用相同变量
func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func day009Test1() {
	a, b := test(100)
	a()
	b()
}

func day009() {
	day009Test1()
}
