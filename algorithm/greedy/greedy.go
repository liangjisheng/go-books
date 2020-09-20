package main

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247486367&idx=2&sn=b25c0701241325f644acda12599ec06a&scene=21#wechat_redirect

// 返回字符串 text 中按字典序排列最小的子序列, 该子序列包含 text 中所有的不同字符一次
// 示例 1:
// 输入: "cdadabcc"
// 输出: "adbc"

// 示例 2:
// 输入: "ecbacba"
// 输出: "eacb"

func smallestSubsequence(text string) string {
	res := make([]byte, 0)
	len := len(text)

	// 使用位掩码来保存 text 中的所有字符, 出现多次的字符保存在一个位置
	// text 中最多26个不同的英文字母, 默认 int 是32位
	all := 0
	for i := 0; i < len; i++ {
		all |= (1 << (text[i] - 'a'))
	}

	// pos 表示所选的当前字符在 text 中的位置
	pos := 0
	// text 中还有字符没有选
	for all != 0 {
		for i := 0; i < 26; i++ {
			// 使用 all & (1 << i) 来判断 text 中是否包含字符 i+'a'
			if (all&(1<<i)) != 0 && isOk(text, all, byte(i)+byte('a'), &pos) {
				res = append(res, byte(i+'a'))
				// all ^= (1 << i) 表示将所选字符从剩下的字符集all中移除
				all ^= (1 << i)
				break
			}
		}
	}

	return string(res)
}

// 函数 isOk(text, all, i+'a', pos) 用来判断字符 i+'a' 之后的所有字符能否包含
// text 中剩下未选的所有字符, 如果包含则附带更新pos的位置
func isOk(text string, all int, ch byte, pos *int) bool {
	len := len(text)
	i := *pos
	// 从 pos 位置开始, 找到 ch 在 text 中的位置
	for ; i < len; i++ {
		if byte(text[i]) == ch {
			break
		}
	}

	p := i + 1
	t := 0
	// 计算从 ch 所在的位置之后, 所有剩下未选的字符的位掩码
	for ; i < len; i++ {
		// ch 所在的位置之后的某个字符包含在剩下未选的字符
		if all&(1<<byte(text[i]-'a')) != 0 {
			t |= (1 << byte(text[i]-'a'))
		}
	}

	// ch 所在的位置之后的所有字符掩码和 all 相同
	if t == all {
		*pos = p
		return true
	}

	return false
}
