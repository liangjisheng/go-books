package main

import (
	"math"
)

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484731&idx=1&sn=f1db6dee2c8e70c42240aead9fd224e6&chksm=9bd7fb33aca07225bee0b23a911c30295e0b90f393af75eca377caa4598ffb203549e1768336&scene=178#rd

// 给你k种面值的硬币, 面值分别为c1, c2 ... ck, 每种硬币的数量无限, 再给一个总金额
// amount, 问你最少需要几枚硬币凑出这个金额, 如果不可能凑出, 算法返回 -1

// 比如说k = 3, 面值分别为 1, 2, 5, 总金额amount = 11. 那么最少需要 3 枚硬币凑出
// 即 11 = 5 + 5 + 1

// 这个问题是动态规划问题, 因为它具有「最优子结构」. 要符合「最优子结构」
// 子问题间必须互相独立

// 先确定「状态」, 也就是原问题和子问题中变化的变量. 由于硬币数量无限, 所以唯一的状态
// 就是目标金额 amount

// 然后确定dp函数的定义: 函数 dp(n)表示, 当前的目标金额是n, 至少需要dp(n)个硬币凑出该金额

// 然后确定「选择」并择优, 也就是对于每个状态, 可以做出什么选择改变当前状态.
// 具体到这个问题, 无论当的目标金额是多少, 选择就是从面额列表coins中选择一个硬币
// 然后目标金额就会减少

// # 伪码框架
// def coinChange(coins: List[int], amount: int):
//     # 定义：要凑出金额 n，至少要 dp(n) 个硬币
//     def dp(n):
//         # 做选择，需要硬币最少的那个结果就是答案
//         for coin in coins:
//             res = min(res, 1 + dp(n - coin))
//         return res
//     # 我们要求目标金额是 amount
//     return dp(amount)

// 最后明确 base case, 显然目标金额为 0 时, 所需硬币数量为 0;
// 当目标金额小于 0 时, 无解, 返回 -1

// 原文是 python 写的, 这里翻译成 go

// 解法一 暴力递归
// 时间复杂度 O(n^k) k 为 coins 的长度

func coinChange1(coins []int, amount int) int {
	// base case 递归结束条件
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	// 求最小值, 所以初始化为 int 最大值
	res := math.MaxInt32
	k := len(coins)
	for i := 0; i < k; i++ {
		subProblem := coinChange1(coins, amount-coins[i])
		// 子问题无解 跳过
		if subProblem == -1 {
			continue
		}

		res = min(res, 1+subProblem)
	}

	// 原问题也无解
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

// 解法二 带备忘录的递归
// 时间复杂度 O(kn) k 为 coins 的长度
// memo[i] 表示当 amount 为 i 的时候, 所需要最小硬币数
// 很显然「备忘录」大大减小了子问题数目, 完全消除了子问题的冗余, 所以子问题总数
// 不会超过金额数 amount, 即子问题数目为 O(n). 处理一个子问题的时间不变
// 仍是 O(k), 所以总的时间复杂度是 O(kn)

func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount+1)
	// 默认初始化为 -1, 表示默认问题是无解的
	for i := 0; i <= amount; i++ {
		memo[i] = -1
	}

	res := coinChange2Helper(coins, memo, amount)
	// fmt.Println("memo:", memo)

	return res
}

func coinChange2Helper(coins []int, memo []int, amount int) int {
	// base case
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		memo[amount] = 0
		return 0
	}
	if memo[amount] != -1 {
		return memo[amount]
	}

	res := math.MaxInt32
	k := len(coins)
	for i := 0; i < k; i++ {
		subProblem := coinChange2Helper(coins, memo, amount-coins[i])
		// 子问题无解 跳过
		if subProblem < 0 {
			continue
		}

		res = min(res, 1+subProblem)
	}

	// 原问题也无解
	if res == math.MaxInt32 {
		memo[amount] = -1
	}

	// 原问题有解
	memo[amount] = res
	return res
}

// 解法三 dp 数组的迭代解法
// 时间复杂度 O(kn) k 为 coins 的长度
// dp[i] = x 表示, 当目标金额为i时, 至少需要x枚硬币

func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 默认初始化为 amount+1
	// 为啥 dp 数组初始化为 amount + 1 呢, 因为凑成 amount 金额的硬币数最多只可能
	// 等于 amount (全用 1 元面值的硬币), 所以初始化为 amount + 1 就相当于初始化
	// 为正无穷, 便于后续取最小值
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	k := len(coins)

	for i := 1; i <= amount; i++ {
		// 内层 for 在求所有子问题 + 1 的最小值
		for j := 0; j < k; j++ {
			// 子问题无解 跳过
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coins[j]])
		}
	}

	// 原问题也无解
	if dp[amount] == amount+1 {
		return -1
	}

	// fmt.Println(dp)
	return dp[amount]
}
