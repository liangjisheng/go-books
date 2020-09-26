package main

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484469&idx=1&sn=e8d321c8ad62483874a997e9dd72da8f&chksm=9bd7fa3daca0732b316aa0afa58e70357e1cb7ab1fe0855d06bc4a852abb1b434c01c7dd19d6&scene=178#rd

// 解法一

// 第一个状态是剩余的按键次数, 用n表示; 第二个状态是当前屏幕上字符 A 的数量,
// 用a_num表示; 第三个状态是剪切板中字符 A 的数量, 用copy表示
// 如此定义状态, 就可以知道 base case: 当剩余次数n为 0 时, a_num就是我们想要的答案
// 结合刚才说的 4 种选择, 我们可以把这几种选择通过状态转移表示出来

// dp(n - 1, a_num + 1, copy),    # A
// 解释: 按下 A 键，屏幕上加一个字符
// 同时消耗 1 个操作数

// dp(n - 1, a_num + copy, copy), # C-V
// 解释: 按下 C-V 粘贴，剪切板中的字符加入屏幕
// 同时消耗 1 个操作数

// dp(n - 2, a_num, a_num)        # C-A C-C
// 解释: 全选和复制必然是联合使用的
// 剪切板中 A 的数量变为屏幕上 A 的数量
// 同时消耗 2 个操作数

// 我们可以把这个 dp 函数写成 dp 数组
// dp[n][a_num][copy]
// 状态的总数(时空复杂度)就是这个三维数组的体积
// 我们知道变量 n 最多为 N, 但是 a_num 和 copy 最多为多少我们很难计算
// 复杂度起码也有 O(N^3) 吧. 所以这个算法并不好, 复杂度太高, 且已经无法优化了
// 这也就说明, 这样定义状态是不太优秀的

type status1 struct {
	n    int
	aNum int
	copy int
}

func maxA1(n int) int {
	// 可以按 n 次键, 屏幕和剪切板里都还没有 A
	return maxA1Helper(n, 0, 0)
}

// 对于 (n, aNum, copy) 这个状态
// 屏幕上能最终有 maxA1Helper(n, aNum, copy) 个 A
func maxA1Helper(n int, aNum int, copy int) int {
	// base case
	if n <= 0 {
		return aNum
	}

	// 三种结果选择全试一遍, 选择最大的结果
	res := max(maxA1Helper(n-1, aNum+1, copy), // A
		maxA1Helper(n-1, aNum+copy, copy)) // C-V
	res = max(res,
		maxA1Helper(n-2, aNum, aNum)) // C-A, C-C

	return res
}

// 带备忘录的解法一
func maxA1Memo(n int) int {
	memo := make(map[status1]int)
	return maxA1MemoHelper(n, 0, 0, memo)
}

func maxA1MemoHelper(n int, aNum int, copy int, memo map[status1]int) int {
	// base case
	if n <= 0 {
		return aNum
	}

	// 避免计算重叠子问题
	_, ok := memo[status1{n, aNum, copy}]
	if ok {
		return memo[status1{n, aNum, copy}]
	}

	// 三种结果选择全试一遍, 选择最大的结果
	res := max(maxA1MemoHelper(n-1, aNum+1, copy, memo), // A
		maxA1MemoHelper(n-1, aNum+copy, copy, memo)) // C-V
	res = max(res,
		maxA1MemoHelper(n-2, aNum, aNum, memo)) // C-A, C-C

	return res
}

// 解法二
// 最后一次按键要么是 A 要么是 C-V
// 定义: dp[i] 表示 i 次操作后最多能显示多少个 A
// 已知 dp[0...i-1] 如何计算 dp[i]
// 对于按 A 键这种情况, 就是状态 i - 1 的屏幕上新增了一个 A 而已, 很容易得到结果
// 按 A 键，就比上次多一个 A 而已
// dp[i] = dp[i - 1] + 1;
// 刚才说了, 最优的操作序列一定是 C-A C-C 接着若干 C-V,
// 所以用一个变量j作为若干C-V的起点. 那么j之前的 2 个操作就应该是 C-A C-C 了

// 当 n 达到 20 以上时, 不管是暴力递归还是带备忘录的递归, 所需要的时间
// 就会远远大于解法二

func maxA2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		// 按 A 键
		dp[i] = dp[i-1] + 1

		// 按 C-V 键
		for j := 2; j < i; j++ {
			// 全选 & 复制 dp[j-2], 连续粘贴 i-j 次
			// 屏幕上共 dp[j-2] * (i-j+1) 个 A
			dp[i] = max(dp[i], dp[j-2]*(i-j+1))
		}
	}

	// N 次按键之后最多有几个 A
	return dp[n]
}
