package main

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484513&idx=1&sn=e5fc3cce76c1b916195e1793122c28b8&chksm=9bd7fa69aca0737fe704ea5c6da28f47b9e3f0961df2eb40ef93a7d507ace8def1a18d013515&scene=178#rd

// 两个普通的字符串进行比较

func isMatch1(text, pattern string) bool {
	lenT := len(text)
	lenP := len(pattern)
	if lenT != lenP {
		return false
	}

	for i := 0; i < lenT; i++ {
		if text[i] != pattern[i] {
			return false
		}
	}
	return true
}

// 这个是另外的一种写法
func isMatch2(text, pattern string) bool {
	lenT := len(text)
	lenP := len(pattern)
	i := 0 // text 索引
	j := 0 // pattern 索引

	for j < lenP {
		if i >= lenT {
			return false
		}
		if pattern[j] != text[i] {
			return false
		}

		i++
		j++
	}

	// 判断 pattern 和 text 是否一样长
	if j == lenT {
		return true
	}
	return false
}

// 改造成递归写法
func isMatch3(text, pattern string) bool {
	if pattern == "" && text == "" {
		return true
	}
	if pattern == "" && text != "" {
		return false
	}
	if pattern != "" && text == "" {
		return false
	}

	firstMatch := text[0] == pattern[0]

	return firstMatch && isMatch3(text[1:], pattern[1:])
}

// 处理点号 . 和 * 通配符
func isMatch4(text, pattern string) bool {
	if pattern == "" {
		if text == "" {
			return true
		}
		return false
	}

	// 处理点号 . 通配符
	firstMatch := text != "" && (pattern[0] == text[0] || pattern[0] == '.')

	// 发现 * 通配符
	if len(pattern) >= 2 && pattern[1] == '*' {
		// 如果发现有字符和 * 结合,
		// 或者匹配该字符 0 次, 然后跳过该字符和 *
		// 或者当 pattern[0] 和 text[0] 匹配后, 移动 text
		// 可以看到, 我们是通过保留 pattern 中的 *, 同时向后推移 text
		// 来实现 * 让字符出现多次的功能
		return isMatch4(text, pattern[2:]) || (firstMatch && isMatch4(text[1:], pattern))
	}

	// 第一个字符匹配后, 接着匹配后面的字符
	return firstMatch && isMatch3(text[1:], pattern[1:])
}

// Range ...
type Range struct {
	i, j int
}

// 使用备忘录递归的方法来降低复杂度
func isMatch5(text, pattern string) bool {
	memo := make(map[Range]bool)
	// text, pattern 分别从 0, 0 处开始匹配
	return isMatch5Helper(text, pattern, memo, 0, 0)
}

func isMatch5Helper(text, pattern string, memo map[Range]bool, i, j int) bool {
	// 如果备忘录有值, 直接从备忘录里拿
	_, ok := memo[Range{i, j}]
	if ok {
		return memo[Range{i, j}]
	}

	// pattern 已经处理到最后, 或者最开始都是空的
	if j == len(pattern) {
		return i == len(text)
	}

	// 处理 . 通配符
	firstMatch := i < len(text) && (pattern[j] == text[i] || pattern[j] == '.')

	res := false
	if j <= len(pattern)-2 && pattern[j+1] == '*' {
		// 发现 * 通配符
		// 或者匹配该字符 0 次, 然后跳过该字符和 *
		// 或者当 pattern[0] 和 text[0] 匹配后, 移动 text
		// 可以看到, 我们是通过保留 pattern 中的 *, 同时向后推移 text
		// 来实现 * 让字符出现多次的功能
		res = isMatch5Helper(text, pattern, memo, i, j+2) || (firstMatch && isMatch5Helper(text, pattern, memo, i+1, j))
	} else {
		// 第一个字符匹配后, 接着匹配后面的字符
		res = firstMatch && isMatch5Helper(text, pattern, memo, i+1, j+1)
	}

	memo[Range{i, j}] = res
	return res
}
