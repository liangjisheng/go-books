package main

import "fmt"

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247486128&idx=2&sn=bcec6b9eb7374169d823963bb7ca8415&scene=21#wechat_redirect

// 给定一个非负整数数组, 你最初位于数组的第一个位置
// 数组中的每个元素代表你在该位置可以跳跃的最大长度
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置

// 输入: [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2
// 从下标为 0 跳到下标为 1 的位置, 跳 1 步, 然后跳 3 步到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置

// 返回跳跃次数 时间复杂度为 O(n^2), 空间复杂度为 O(1)
func jump1(array []int) int {
	right := len(array) - 1
	step := 0
	// path := []int{right} // 记录跳跃路径

	for right > 0 {
		// 寻找离 last 最远, 且能够到达 last 的点
		cur := right - 1
		for i := right - 2; i >= 0; i-- {
			// 是否能够达到 last 这个点
			if i+array[i] >= right {
				cur = i
			}
		}

		// 更新 last 点位置
		right = cur
		// path = append([]int{right}, path...)
		step++
	}

	return step
}

// 最优解是可以时间复杂度优化到 O(n) 的, 那就是采用贪心算法
// 从左边的起点开始跳跃的时候, 应该跳跃到哪一个点比较合适呢
// 要跳跃的那个点, 可以使得上一次 + 下一次的跳跃总距离最远
// 时间复杂度是 O(n), 空间复杂度是 O(1)
func jump2(array []int) int {
	n := len(array)
	if n < 2 {
		return 0
	}

	step := 0
	end := 0    // 能跳到的最远距离
	maxDis := 0 // 下一步可以跳到的最远距离

	for i := 0; i < n; i++ {
		// maxDis 表示到 i 这里下一步能跳到的最远索引处
		maxDis = max(maxDis, i+array[i])
		fmt.Printf("maxDis: %d\n", maxDis)
		// 如果上一步能跳的最远索引处, 还没超过数组的最大索引, 则更新
		// end 为到 i 这里能跳的最远索引处
		if i == end {
			end = maxDis
			fmt.Printf("end: %d\n", end)
			step++
		}
	}

	return step
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
