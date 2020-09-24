package main

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484731&idx=1&sn=f1db6dee2c8e70c42240aead9fd224e6&chksm=9bd7fb33aca07225bee0b23a911c30295e0b90f393af75eca377caa4598ffb203549e1768336&scene=178#rd

// 重叠子问题、最优子结构、状态转移方程就是动态规划三要素
// 明确「状态」 -> 定义 dp 数组/函数的含义 -> 明确「选择」-> 明确 base case

// 一、斐波那契数列

// 解法一 暴力递归
// 递归算法的时间复杂度怎么计算? 子问题个数乘以解决一个子问题需要的时间

// 时间复杂度 O(2^n) 空间复杂度 O(n)
func fibonacci1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibonacci1(n-1) + fibonacci1(n-2)
}

// 解法二 带备忘录的递归解法 时间复杂度 O(n) 空间复杂度 O(n)
// 这种方法叫做自顶向下
func fibonacci2(n int) int {
	if n <= 0 {
		return 0
	}

	// 初始化备忘录为 0
	memo := make([]int, n+1)
	return helper2(memo, n)
}

func helper2(memo []int, n int) int {
	// base case 递归结束条件
	if n == 1 || n == 2 {
		return 1
	}

	// 如果已经计算过, 则直接从备忘录里拿
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = helper2(memo, n-1) + helper2(memo, n-2)
	return memo[n]
}

// 解法三 dp 数组的迭代解法
// 有了上一步「备忘录」的启发, 我们可以把这个「备忘录」独立出来成为一张表
// 就叫做 DP table 吧, 在这张表上完成「自底向上」的推算
// 时间复杂度: O(n) 空间复杂度: O(n)

func fibonacci3(n int) int {
	dp := make([]int, n+1)
	// base case
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 解法四: 解法三的细节优化
// 当前状态只和之前的两个状态有关, 其实并不需要那么长的一个 DP table 来存储所有的状态
// 只要想办法存储之前的两个状态就行了. 所以, 可以进一步优化, 把空间复杂度降为 O(1)

func fibonacci4(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	prev, cur := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + cur
		prev = cur
		cur = sum
	}

	return cur
}
