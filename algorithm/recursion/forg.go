package main

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级
// 求该青蛙跳上一个n级的台阶总共有多少种跳法
// 等价关系式: f(n) = f(n-1) + f(n-2)

func forgJumpStep(n uint64) uint64 {
	// if n <= 1 {
	// 	return 1
	// }

	// 只要你觉得参数是什么时 你能够直接知道函数的结果
	// 那么你就可以把这个参数作为结束的条件,这2个结束条件是等价的
	if n <= 2 {
		return n
	}

	return forgJumpStep(n-1) + forgJumpStep(n-2)
}
