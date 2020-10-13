package main

import "fmt"

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484475&idx=1&sn=8e9518d67ae8f4c16f14fb0c4d584c79&chksm=9bd7fa33aca07325c056c017b7ff5b434a11fe7fee1a0c14aacbc9f1dd317bb7770cb1faef36&scene=178#rd

// KMP ...
type KMP struct {
	dp  [][]int
	pat string
}

func (kmp *KMP) setDP(pat string) {
	kmp.pat = pat
	bPat := []byte(pat)
	m := len(pat)
	// dp[状态][字符] = 下个状态
	kmp.dp = make([][]int, m)
	for i := 0; i < m; i++ {
		kmp.dp[i] = make([]int, 256)
	}

	// base case
	// 只有遇到 pat[0] 这个字符才能使状态从 0 转移到 1, 遇到其它字符的话
	// 还是停留在状态 0
	kmp.dp[0][bPat[0]] = 1

	// 影子状态 X 初始为 0
	x := 0

	// 当前状态 j 从 1 开始
	for j := 1; j < m; j++ {
		for c := 0; c < 256; c++ {
			if bPat[j] == byte(c) {
				kmp.dp[j][c] = j + 1
			} else {
				kmp.dp[j][c] = kmp.dp[x][c]
			}
		}

		// 更新影子状态
		// 当前是状态 X, 遇到字符 pat[j]
		// pat 应该转移到哪个状态?
		x = kmp.dp[x][bPat[j]]
	}
}

func (kmp *KMP) search(txt string) int {
	m := len(kmp.pat)
	n := len(txt)
	bTxt := []byte(txt)

	// pat 的初始态为 0
	j := 0
	for i := 0; i < n; i++ {
		// 计算 pat 的下一个状态
		j = kmp.dp[j][bTxt[i]]
		// 到达终止态, 返回结果
		if j == m {
			return i - m + 1
		}
	}

	// 没到达终止态, 匹配失败
	return -1
}

func kmp3() {
	kmp := &KMP{}
	pat := "ABCDABD"
	txt := "BBC ABCDAB ABCDABCDABDE"
	kmp.setDP(pat)
	fmt.Println(kmp.search(txt))
}
