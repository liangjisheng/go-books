package recursion

import "github.com/shopspring/decimal"

// 一只青蛙一次可以跳上1级台阶，也可以跳上2级
// 求该青蛙跳上一个n级的台阶总共有多少种跳法
// 等价关系式: f(n) = f(n-1) + f(n-2)

// ForgJumpStep ...
func ForgJumpStep(n uint64) uint64 {
	// if n <= 1 {
	// 	return 1
	// }

	// 只要你觉得参数是什么时 你能够直接知道函数的结果
	// 那么你就可以把这个参数作为结束的条件,这2个结束条件是等价的
	if n <= 2 {
		return n
	}

	return ForgJumpStep(n-1) + ForgJumpStep(n-2)
}

// ForgJumpStepIteration ...
func ForgJumpStepIteration(n uint64) uint64 {
	if n <= 2 {
		return n
	}

	var f1 uint64 = 1
	var f2 uint64 = 2
	var i uint64 = 3
	var res uint64
	for ; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}

// ForgJumpStepIterationDecimal ...
func ForgJumpStepIterationDecimal(n uint64) decimal.Decimal {
	if n <= 2 {
		return decimal.New(int64(n), 0)
	}

	f1 := decimal.New(1, 0)
	f2 := decimal.New(2, 0)
	var i uint64 = 3
	var res decimal.Decimal
	for ; i <= n; i++ {
		res = f1.Add(f2)
		f1 = f2
		f2 = res
	}
	return res
}
