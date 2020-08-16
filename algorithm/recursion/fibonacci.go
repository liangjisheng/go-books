package recursion

// 斐波那契数列的是这样一个数列：1、1、2、3、5、8、13、21、34….，即第一项
// f(1) = 1,第二项 f(2) = 1…..,第 n 项目为 f(n) = f(n-1) + f(n-2)
// 求第 n 项的值是多少

// FibonacciRecursion ...
func FibonacciRecursion(n uint64) uint64 {
	if n <= 2 {
		return 1
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// FibonacciIteration ...
func FibonacciIteration(n uint64) uint64 {
	if n <= 2 {
		return 1
	}
	var f1 uint64 = 1
	var f2 uint64 = 1
	var res uint64

	var i uint64 = 3
	for ; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}
