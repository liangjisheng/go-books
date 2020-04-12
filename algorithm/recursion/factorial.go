package main

// 递归求阶乘
// 等价关系式 f(n) = f(n-1)*n

func factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return factorial(n-1) * n
}

// 递推
func factorialRecurrence(n uint64) uint64 {
	if n <= 2 {
		return n
	}
	var f1 uint64 = 1
	var f2 uint64 = 2
	var res uint64

	var i uint64 = 3
	for ; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}
