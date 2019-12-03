package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func day003Test1() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

var f = func(i int) {
	print("x")
}

// 这道题一眼看上去会输出 109876543210 其实这是错误的答案 这里不是递归 假设
// day003Test2() 函数里为 f2() 外面的为 f1()，当声明 f2() 时 调用的是已经完成声明的 f1()
func day003Test2() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}

func day003() {
	// day003Test1()
	day003Test2() // 10x
}
