package main

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484496&idx=1&sn=d04bb89cb1df241993c6b46ffcabae7e&chksm=9bd7fa58aca0734e7241b3459c6775f9ae30d16cb8d30b5ad897600a87c69fd77d7284742043&scene=178#rd

// Pair ...
type Pair struct {
	first, second int
}

// 返回游戏最后先手和后手的得分之差

func stoneGame(piles []int) int {
	n := len(piles)
	// 初始化 dp 二维数组
	dp := make([][]Pair, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]Pair, n)
	}

	// 填入 base case
	for i := 0; i < n; i++ {
		dp[i][i].first = piles[i]
		dp[i][i].second = 0
	}

	// 斜着遍历数组
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := l + i - 1
			// 先手选择最左边或者最右边的分数
			left := piles[i] + dp[i+1][j].second
			right := piles[j] + dp[i][j-1].second
			// 套用状态转移方程
			if left > right {
				dp[i][j].first = left
				dp[i][j].second = dp[i+1][j].first
			} else {
				dp[i][j].first = right
				dp[i][j].second = dp[i][j-1].first
			}
		}
	}

	// 斜着遍历数组打印
	// for k := 0; k < n; k++ {
	// 	for i := 0; i < n-k; i++ {
	// 		j := i + k
	// 		fmt.Printf("%+v ", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }

	res := dp[0][n-1]
	return res.first - res.second
}
