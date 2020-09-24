package main

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484498&idx=1&sn=df58ef249c457dd50ea632f7c2e6e761&chksm=9bd7fa5aaca0734c29bcf7979146359f63f521e3060c2acbf57a4992c887aeebe2a9e4bd8a89&scene=178#rd

// 给定一个无序的整数数组, 找到其中最长上升子序列的长度
// 注意子序列和子串这两个名词的区别, 子串一定是连续的, 而子序列不一定是连续的

// 示例
// 输入: [10, 9, 2, 5, 3, 7, 101, 18]
// 输出: 4
// 解释: 最长的上升子序列是 [2, 3, 7, 101] 长度为 4

// 定义 dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	// dp 数组应该全部初始化为 1, 因为子序列最少也要包含自己, 所以长度最小为 1
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	// for 循环求 dp[0...n]
	for i := 0; i < n; i++ {
		// 内层循环通过 dp[0...i-1] 求 dp[i], 这里是关键
		// 通过 dp[0...i-1] 这 i-1 个状态求第 i 个状态
		// 内层循环可以理解为状态转移方程
		for j := 0; j < i; j++ {
			// 找到比 nums[i] 小, 形成子序列, 可能有多种情况, 取最大值
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 最终结果(子序列的最大长度)应该是 dp 数组中的最大值
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}
