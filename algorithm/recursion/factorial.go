package recursion

// 递归求阶乘
// 等价关系式 f(n) = f(n-1)*n

// FactorialRecursion ...
func FactorialRecursion(n uint64) uint64 {
	if n <= 2 {
		return n
	}
	return FactorialRecursion(n-1) * n
}

// FactorialIteration ...
func FactorialIteration(n uint64) uint64 {
	if n <= 2 {
		return n
	}

	var res uint64 = 1
	var i uint64 = 2
	for ; i <= n; i++ {
		res = i * res
	}
	return res
}
