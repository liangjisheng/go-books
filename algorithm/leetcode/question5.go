package main

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247485354&idx=2&sn=d817c66b7e107575243a18e4d5aef1b1&scene=21#wechat_redirect

// 编写一个函数来查找字符串数组中的最长公共前缀
// 如果不存在公共前缀, 返回空字符串 ""
// 所有输入只包含小写字母 a-z

// 示例 1:
// 输入: ["flower","flow","flight"]
// 输出: "fl"

// 示例 2:
// 输入: ["dog","racecar","car"]
// 输出: ""
// 解释: 输入不存在公共前缀

// 不横向一个一个字符串遍历, 而是采用纵向的方式
// f l o w e r
// f l o w
// f l i g h t
// 然后纵向遍历, 一列一列遍历, 只要发现某一列出现不同的字符, 就遍历结束

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	n := len(strs)
	bStrs := make([][]byte, 0) // 将 string 转换成 []byte 存储
	minLen := len(strs[0])     // 所有字符串最小长度
	minStr := []byte(strs[0])  // 长度最小字符串
	for i := 0; i < n; i++ {
		bStrs = append(bStrs, []byte(strs[i]))
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
			minStr = []byte(strs[i])
		}
	}

	res := make([]byte, 0)

	// 最长公共前缀不会超过字符串中最小长度
	for i := 0; i < minLen; i++ {
		char := minStr[i] // 第 i 列的字符
		for j := 0; j < n; j++ {
			// 判断所有字符串的第 i 列是否都是一样的,如果有某个不一样
			// 就说明找到了最长公共前缀, 直接返回之前的结果
			if char != bStrs[j][i] {
				return string(res)
			}
		}

		// 第 i 列的字符属于公共前缀的一部分
		res = append(res, char)
	}

	return string(res)
}
