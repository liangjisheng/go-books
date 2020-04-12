package main

// 斐波那契数列的是这样一个数列：1、1、2、3、5、8、13、21、34….，即第一项
// f(1) = 1,第二项 f(2) = 1…..,第 n 项目为 f(n) = f(n-1) + f(n-2)
// 求第 n 项的值是多少

func fibonacci(n uint64) uint64 {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
