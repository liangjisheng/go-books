package main

// 暴力匹配
// 假设现在文本串S匹配到 i 位置, 模式串P匹配到 j 位置, 则有:
// 如果当前字符匹配成功(即S[i] == P[j]), 则i++, j++, 继续匹配下一个字符;
// 如果失配(即S[i] != P[j]), 令i = i - (j - 1), j = 0.
// 相当于每次匹配失败时, i 回溯, j 被置为0

func violentMatch(s, p string) int {
	bs := []byte(s)
	bp := []byte(p)
	sLen := len(bs)
	pLen := len(bp)
	i, j := 0, 0

	for i < sLen && j < pLen {
		if bs[i] == bp[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	// 匹配成功, 返回模式串p在文本串s中的索引, 否则返回-1
	if j == pLen {
		return i - j
	}
	return -1
}
