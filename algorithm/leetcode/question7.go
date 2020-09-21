package main

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247485335&idx=2&sn=8788fc59a9f02f2c0401efd428e00ea6&scene=21#wechat_redirect

// 给定一罗马数字, 将其转为整数. 输入确保在 1 到 3999 的范围内

// 示例 1:
// 输入: "III"
// 输出: 3

// 示例 2:
// 输入: "IV"
// 输出: 4

// 示例 3:
// 输入: "IX"
// 输出: 9

// 示例 4:
// 输入: "LVIII"
// 输出: 58
// 解释: L = 50, V= 5, III = 3.

// 示例 5:
// 输入: "MCMXCIV"
// 输出: 1994
// 解释: M = 1000, CM = 900, XC = 90, IV = 4.

// 我们把这些字符一个一个判断就可以了, 例如遇到 V 就加 5, 然后向右前进一个字符
// 遇到 IV 就加 4, 然后向右前进 2 个字符. 不过在判断的是时候,
// 我们要优先 V 再判断 IV

func romanToInt(s string) int {
	b := []byte(s)
	n := len(s)
	i := 0
	num := 0

	// 1000
	for i < n && b[i] == byte('M') {
		num += 1000
		i++
	}

	// 900
	if i+1 < n && b[i] == byte('C') && b[i+1] == byte('M') {
		num += 900
		i += 2
	}

	// 500
	if i < n && b[i] == byte('D') {
		num += 500
		i++
	}

	// 400
	if i+1 < n && b[i] == byte('C') && b[i+1] == byte('D') {
		num += 400
		i += 2
	}

	// 100
	for i < n && b[i] == byte('C') {
		num += 100
		i++
	}

	// 90
	if i+1 < n && b[i] == byte('X') && b[i+1] == byte('C') {
		num += 90
		i += 2
	}

	// 50
	if i < n && b[i] == byte('L') {
		num += 50
		i++
	}

	// 40
	if i+1 < n && b[i] == byte('X') && b[i+1] == byte('L') {
		num += 40
		i += 2
	}

	// 10
	for i < n && b[i] == byte('X') {
		num += 10
		i++
	}

	// 9
	if i+1 < n && b[i] == byte('I') && b[i+1] == byte('X') {
		num += 9
		i += 2
	}

	// 5
	if i < n && b[i] == byte('V') {
		num += 5
		i++
	}

	// 4
	if i+1 < n && b[i] == byte('I') && b[i+1] == byte('V') {
		num += 4
		i += 2
	}

	// 1
	for i < n && b[i] == byte('I') {
		num++
		i++
	}

	return num
}

func romanToInt2(s string) int {
	b := []byte(s)
	n := len(s)
	num := charToInt(b[0])

	for i := 1; i < n; i++ {
		// 如果出现 IV,IX,XL,XC,CD,CM 这种情况, 那么我们需要减去两倍的 I,X 等
		// 因为它前面被我们多加了一次
		tmpCur := charToInt(b[i])
		tmpPre := charToInt(b[i-1])
		if tmpCur > tmpPre {
			num = num + tmpCur - 2*tmpPre
		} else {
			num = num + tmpCur
		}
	}

	return num
}

func charToInt(c byte) int {
	switch c {
	case byte('I'):
		return 1
	case byte('V'):
		return 5
	case byte('X'):
		return 10
	case byte('L'):
		return 50
	case byte('C'):
		return 100
	case byte('D'):
		return 500
	case byte('M'):
		return 1000
	default:
		return 0
	}
}
