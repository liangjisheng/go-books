package bitoperation

// 如果一个数是奇数的话,那么最后一个二进制位为1
// 如果一个数是偶数的话,那么最后一个二进制位为0
// 用 n & (n - 1) 可以消去 n 最后的一位 1

// IsOdd 判断一个数是否是奇数
func IsOdd(n int) bool {
	if n&1 == 1 {
		return true
	}
	return false
}

// 2个相同的数异或运算的结果为0, 即 n^n = 0
// 0与任何数异或运算的结果为任何数, 即 0^n = n
// 异或运算支持运算的交换律和结合律

// 使用异或运算不借助额外的变量交换2个变量
func swap(x, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}

// 给你一组整型数据, 这些数据中, 其中有一个数只出现了一次, 其他的数都出现了两次
// 让你来找出一个数

// 解法1: 使用数据作为数组的下标, 遍历数据结束后, 只有某个数的索引对应的值为1

// 解法2: 使用异或运算性质 2个相同的数异或运算的结果为0
func findOnlyOne(array []int) int {
	n := len(array)
	tmp := 0
	for i := 0; i < n; i++ {
		tmp = tmp ^ array[i]
	}
	return tmp
}

// 位运算求 m 的 n 次方
// 例如 n = 13, 则 n 的二进制表示为 1101, 那么 m 的 13 次方可以拆解为:
// m^1101 = m^0001 * m^0100 * m^1000
// 可以通过 & 1和 >>1 来逐位读取 1101, 为1时将该位代表的乘数累乘到最终结果

// 时间复杂度近似为 O(logn)
func pow(m, n int) int {
	tmp := m
	res := 1
	for n != 0 {
		if n&1 == 1 {
			res = res * tmp
		}
		tmp *= tmp
		n >>= 1
	}
	return res
}

// 求 n 二进制表示中 1 的个数
func numberOf1(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
