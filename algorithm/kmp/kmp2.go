package main

import "fmt"

// next 数组各值的含义: 代表当前字符之前的字符串中, 有多大长度的相同前缀后缀
// 例如如果next [j] = k, 代表j 之前的字符串中有最大长度为k 的相同前缀后缀

// next 数组值默认为-1
func getNext(p string, next []int) {
	bp := []byte(p)
	pLen := len(bp)
	k, j := -1, 0

	for j < pLen-1 {
		// bp[k] 表示前缀, bp[j] 表示后缀
		if k == -1 || bp[j] == bp[k] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
}

// 优化后的求next
// 问题出在不该出现 p[j] = p[ next[j] ]. 为什么呢? 理由是: 当p[j] != s[i] 时
// 下次匹配必然是 p[ next [j]] 跟s[i]匹配, 如果 p[j] = p[ next[j] ]
// 必然导致后一步匹配失败(因为p[j]已经跟s[i]失配, 然后你还用跟p[j]等同的值
// p[next[j]]去跟s[i]匹配, 很显然, 必然失配), 所以不能允许
// p[j] = p[ next[j ]]. 如果出现了p[j] = p[ next[j] ]咋办呢?如果出现了
// 则需要再次递归, 即令next[j] = next[ next[j] ]
func getNextOptisize(p string, next []int) {
	bp := []byte(p)
	pLen := len(bp)
	k, j := -1, 0

	for j < pLen-1 {
		// bp[k] 表示前缀, bp[j] 表示后缀
		if k == -1 || bp[j] == bp[k] {
			k++
			j++

			// 较之前 next 数组求法, 改动在下面
			if bp[j] != bp[k] {
				next[j] = k
			} else {
				// 因为不能出现 p[j] = p[ next[j ]], 所以当出现时需要继续递归
				// k = next[k] = next[next[k]]
				next[j] = next[k]
			}
		} else {
			k = next[k]
		}
	}
}

// 该字符对应的 next 值会告诉你下一步匹配中, 模式串应该跳到哪个位置
// (跳到next [j] 的位置). 如果 next [j] 等于0或-1, 则跳到模式串的开头字符
// 若next [j] = k 且 k > 0, 代表下次匹配跳到j 之前的某个字符
// 而不是跳到开头, 且具体跳过了k 个字符

func kmpSearch(s, p string, next []int) int {
	i, j := 0, 0
	bs := []byte(s)
	bp := []byte(p)
	sLen := len(bs)
	pLen := len(bp)

	for i < sLen && j < pLen {
		// j = -1, 或者当前字符匹配成功(即S[i] == P[j]), 都令i++, j++
		if j == -1 || bs[i] == bp[j] {
			i++
			j++
		} else {
			// 如果j != -1, 且当前字符匹配失败(即S[i] != P[j])
			// 则令 i 不变, j = next[j]
			j = next[j]
		}
	}

	if j == pLen {
		return i - j
	}
	return -1
}

func kmp2() {
	p := "ABCDABD"
	n := len(p)
	next := make([]int, n)
	for i := 0; i < n; i++ {
		next[i] = -1
	}

	// getNext(p, next)
	getNextOptisize(p, next)

	s := "BBC ABCDAB ABCDABCDABDE"
	fmt.Println(kmpSearch(s, p, next))
}
