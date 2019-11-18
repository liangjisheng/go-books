package main

func defertest() {
	println(deferFunc1(1)) // 4
	println(deferFunc2(1)) // 1
	println(deferFunc3(1)) // 3
}

func deferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
