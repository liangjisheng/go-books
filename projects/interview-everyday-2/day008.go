package main

func day008Test1() {
	for i := 0; i < 10; i++ {
		// loop:
		println(i)
	}
	// goto 不能跳转到其他函数或者内层代码
	// goto loop jumps into block starting at .\day008.go:4:26
	// goto loop
}

func day008Test2() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v)
		}()
		y[i] = &v
	}
	print(*y[0], *y[1], *y[2])
}

func day008() {
	day008Test2()
}
