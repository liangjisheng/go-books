package main

// 前缀和应用
// 输入n个数的数列, 所有相邻m数的和有n-m+1个
// 求其中的最小值(最大值同样的道理)
// 比如：
// 数组为：[10, 4, 1, 5, 5, 2]
// m为：3
// 结果为：10

func minM(array []int, m int) int {
	if m <= 0 {
		return 0
	}

	n := len(array)
	if n <= 0 {
		return 0
	}

	// 求前缀和
	prefixSum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			prefixSum[i] = array[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + array[i]
		}
	}

	result := prefixSum[m-1] // 数组前 m 个数的和
	for i := m; i < n; i++ {
		current := prefixSum[i] - prefixSum[i-m]
		if current < result {
			result = current
		}
	}
	return result
}
