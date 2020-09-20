package main

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247486548&idx=2&sn=d327ea89a997e202e029ad1ccc201b6a&scene=21#wechat_redirect

// 给定一个字符串, 请你找出其中不含有重复字符的 最长子串 的长度

// 示例1:
// 输入: "abcabcbb"输出: 3
// 解释: 因为无重复字符的最长子串是 "abc", 所以其长度为 3

// 示例2：
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke", 所以其长度为 3
// 请注意, 你的答案必须是 子串 的长度, "pwke" 是一个子序列, 不是子串

// 解法一 暴力法
// 索引从字符串的第一位开始, 将后面的字符依次加入到 set 里面. 如果 set 里面
// 已经有了该字符, 此次循环结束, 内循环结束后记录 size. 字符串的每一位都用这种
// 方法去计算, 得到的最大的 size 即是答案
// 时间复杂度: O(n^2), 空间复杂度：O(m), m 为无重复字符的最长子串的长度
func lengthOfLongestSubstring1(s string) int {
	maxLen := 0
	b := []byte(s)
	n := len(b)
	for i := 0; i < n; i++ {
		// 创建一个存放字符的集合
		charMap := make(map[byte]bool)
		// 从 i 开始寻找对 i 来说的最长无重复子串
		for j := i; j < n; j++ {
			// 判断集合是否存在第 j 个字符
			_, ok := charMap[b[j]]
			if ok {
				break
			}
			charMap[b[j]] = true
		}

		// 得到第 i 个为止的最长无重复子串长度
		maxLen = max(maxLen, len(charMap))
	}

	return maxLen
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// 解法二
// 设置两根指针(i 和 j)与一个集合set. 两个指针之间是一个范围, 我们要维护这个范围
// 内不能出现重复的字符. 而这个 set 就是是用来判断范围内是否有重复的字符.
// 具体做法如下: i 和 j 开始的时候都指向 字符串首. 然后执行下面两个步骤:
// 1、如果 j 指针所指元素未在 set 里面, 我们将其 add 进 set. 继续后移 j
// 2、如果 j 指针所指元素在 set 里面, 我们将 i 指针所指元素从 set 中移除, 继续后移 i,
// i 会一直往后移, 直到 j 的元素不在 set 里面
// 这个过程中, i 与 j 的范围最大的时候, 就是我们要求的答案
// 在上面所提到的“范围”, 就是一个滑动窗口

// 时间复杂度: O(2n) = O(n), i 和 j 两个指针
// 空间复杂度: O(m), m 为无重复字符的最长子串的长度
func lengthOfLongestSubstring2(s string) int {
	maxLen := 0
	b := []byte(s)
	n := len(b)
	i, j := 0, 0
	charMap := make(map[byte]bool)

	for i < n && j < n {
		_, ok := charMap[b[j]]
		if !ok {
			charMap[b[j]] = true
			j++
			maxLen = max(maxLen, j-i)
		} else {
			delete(charMap, b[i])
			i++
		}
	}

	return maxLen
}

// 解法三 让 i 指针直接跳到重复元素的下一个位置. 那我们就需要保存每个元素以及它的位置
// 一个 Value, 一个 Index. 自然会想到 HashMap
// 时间复杂度: O(n), 最优解. 空间复杂度: O(m), m 为无重复字符的最长子串的长度
func lengthOfLongestSubstring3(s string) int {
	maxLen := 0
	b := []byte(s)
	n := len(b)
	i, j := 0, 0
	// 创建一个 key-value 为 字符-下标 相映射的哈希表
	charMap := make(map[byte]int)

	for ; j < n; j++ {
		// 判断是否存在第 j 个字符
		idx, ok := charMap[b[j]]
		if ok {
			// i 指针直接跳到重复元素的下一个位置和 i 的最大位置处
			// 因为重复元素可能在 i 之前
			i = max(idx+1, i)
		}

		// 计算最大长度
		maxLen = max(maxLen, j-i+1)
		// 覆盖之前的 key-value
		charMap[b[j]] = j
	}

	// fmt.Printf("%+v\n", charMap)
	return maxLen
}
