package main

import (
	"container/list"
)

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247486410&idx=1&sn=3327affe430dcacffa76a29b0fcaa9ca&scene=21#wechat_redirect

// 给定一个只包括 '(', ')'的字符串, 判断字符串是否有效. 注: 空字符串属于有效字符串

// 解法一: 使用栈

func isValid1(s string) bool {
	if s == "" {
		return true
	}

	b := []byte(s)
	n := len(b)
	stack := list.New()

	for i := 0; i < n; i++ {
		if b[i] == byte('(') {
			stack.PushBack(b[i])
		} else {
			if stack.Len() == 0 {
				// 到 ) 入栈, 此时栈空说明字符串无效
				return false
			}
			// 弹出栈顶元素
			stack.Remove(stack.Back())
		}
	}

	// 判断栈是否为空
	if stack.Len() == 0 {
		return true
	}
	return false
}

// 解法二
// 由于我们栈里面存放的都是同一种字符 "(", 其实我们可以用一个变量来取代栈的
// 这个变量就记录 "(" 的个数, 遇到 "(" 变量就加 1, 遇到 ")" 变量就减 1
// 栈为空就相当于变量的值为 0

func isValid2(s string) bool {
	if s == "" {
		return true
	}

	b := []byte(s)
	n := len(b)
	// 用来记录最后 ( 的个数
	sum := 0

	for i := 0; i < n; i++ {
		if b[i] == byte('(') {
			sum++
		} else {
			if sum == 0 {
				// 到 ), 如果此时 ( 的个数为 0, 说明无效
				return false
			}
			sum--
		}
	}

	if sum == 0 {
		return true
	}
	return false
}

// 给定一个只包含 '(' 和 ')' 的字符串, 找出最长的包含有效括号的子串的长度

// 示例 1:

// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"
// 示例 2:

// 输入: ")()())"
// 输出: 4
// 解释: 最长有效括号子串为 "()()"

// 时间复杂度为 O(n), 空间复杂度为 O(n)

func longestValidParentheses1(s string) int {
	longest := 0
	b := []byte(s)
	n := len(s)
	stack := list.New()
	// 初始化 push -1
	stack.PushBack(-1)

	for i := 0; i < n; i++ {
		if b[i] == byte('(') {
			stack.PushBack(i) // 下标入栈
		} else {
			stack.Remove(stack.Back()) // 出栈
			// 看栈顶是否为空, 为空的话就不能作差了
			if stack.Len() == 0 {
				stack.PushBack(i)
			} else {
				// i - 栈顶, 获得当前有效括号长度
				top := stack.Back()
				longest = max(longest, i-top.Value.(int))
			}
		}
	}

	return longest
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// 解法二
// 时间复杂度为 O(n), 空间复杂度为 O(1)

func longestValidParentheses2(s string) int {
	left := 0
	right := 0
	b := []byte(s)
	n := len(b)
	longest := 0

	// 从左到右遍历
	for i := 0; i < n; i++ {
		if b[i] == byte('(') {
			left++
		} else {
			right++
		}

		// 从左到右的时候, left记录的是 ( 的个数, 而到第 i 个字符的时候 ( 的个数
		// 肯定是不会小于 ) 的个数的, 当 left < right 的时候, 说明 ) 比 ( 多
		// 出现了, 此时重置 left, right
		if left == right {
			longest = max(longest, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0

	// 从右到左遍历
	for i := n - 1; i >= 0; i-- {
		if b[i] == byte('(') {
			left++
		} else {
			right++
		}

		// 从右到左的时候, ) 的个数是不会少于 ( 的个数的, 当 right < left
		// 的时候, 说明 ) 的个数是少于 ( 的个数, 此时重置 left, right
		if left == right {
			longest = max(longest, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return longest
}
