package main

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247485335&idx=2&sn=8788fc59a9f02f2c0401efd428e00ea6&chksm=ce404c43f937c5554d0c47fd0618317f5739745e82166c2d4840a3b44a671ed9c2a667aefeef&scene=21#wechat_redirect

// 罗马数字包含以下七种字符: I, V, X, L, C, D 和 M
// 字符   数值
// I      1
// V      5
// X      10
// L      50
// C      100
// D      500
// M      1000

// 例如, 罗马数字 2 写做 II, 即为两个并列的 1. 12 写做 XII, 即为 X + II.
// 27 写做  XXVII, 即为 XX + V + II

// 通常情况下, 罗马数字中小的数字在大的数字的右边. 但也存在特例, 例如 4 不写做 IIII
// 而是 IV. 数字 1 在数字 5 的左边, 所表示的数等于大数 5 减小数 1 得到的数值 4.
// 同样地, 数字 9 表示为 IX. 这个特殊的规则只适用于以下六种情况:
// I 可以放在 V (5) 和 X (10) 的左边, 来表示 4 和 9
// X 可以放在 L (50) 和 C (100) 的左边, 来表示 40 和 90
// C 可以放在 D (500) 和 M (1000) 的左边, 来表示 400 和 900

// 给定一个整数, 将其转为罗马数字. 输入确保在 1 到 3999 的范围内

// 示例 1:
// 输入: 3
// 输出: "III"

// 示例 2:
// 输入: 4
// 输出: "IV"

// 示例 3:
// 输入: 9
// 输出: "IX"

// 示例 4:
// 输入: 58
// 输出: "LVIII"
// 解释: L = 50, V = 5, III = 3.

// 示例 5:
// 输入: 1994
// 输出: "MCMXCIV"

// 解释: M = 1000, CM = 900, XC = 90, IV = 4

func intToRoman(num int) string {
	res := make([]byte, 0)

	for num >= 1000 {
		res = append(res, byte('M'))
		num -= 1000
	}

	if num >= 900 {
		res = append(res, []byte("CM")...)
		num -= 900
	}
	if num >= 500 {
		res = append(res, byte('D'))
		num -= 500
	}
	if num >= 400 {
		res = append(res, []byte("CD")...)
		num -= 400
	}

	for num >= 100 {
		res = append(res, byte('C'))
		num -= 100
	}

	if num >= 90 {
		res = append(res, []byte("XC")...)
		num -= 90
	}
	if num >= 50 {
		res = append(res, byte('L'))
		num -= 50
	}
	if num >= 40 {
		res = append(res, []byte("XL")...)
		num -= 40
	}

	for num >= 10 {
		res = append(res, byte('X'))
		num -= 10
	}

	if num >= 9 {
		res = append(res, []byte("IX")...)
		num -= 9
	}
	if num >= 5 {
		res = append(res, byte('V'))
		num -= 5
	}
	if num >= 4 {
		res = append(res, []byte("IV")...)
		num -= 4
	}

	for num >= 1 {
		res = append(res, byte('I'))
		num--
	}

	return string(res)
}

func intToChar(num int) []byte {
	switch num {
	case 1:
		return []byte("I")
	case 4:
		return []byte("IV")
	case 5:
		return []byte("V")
	case 9:
		return []byte("IX")
	case 10:
		return []byte("X")
	case 40:
		return []byte("XL")
	case 50:
		return []byte("L")
	case 90:
		return []byte("XC")
	case 100:
		return []byte("C")
	case 400:
		return []byte("CD")
	case 500:
		return []byte("D")
	case 900:
		return []byte("CM")
	case 1000:
		return []byte("M")
	default:
		return []byte("")
	}
}

func intToRoman2(num int) string {
	res := make([]byte, 0)
	tmp := 100

	// 循环3次是由罗马数字到1000决定的
	for i := 0; i < 3; i++ {
		for num >= 10*tmp {
			res = append(res, intToChar(10*tmp)...)
			num -= 10 * tmp
		}

		if num >= 9*tmp {
			res = append(res, intToChar(9*tmp)...)
			num -= 9 * tmp
		}
		if num >= 5*tmp {
			res = append(res, intToChar(5*tmp)...)
			num -= 5 * tmp
		}
		if num >= 4*tmp {
			res = append(res, intToChar(4*tmp)...)
			num -= 4 * tmp
		}

		for num >= 1*tmp {
			res = append(res, intToChar(1*tmp)...)
			num -= 1 * tmp
		}

		tmp /= 10
	}

	return string(res)
}
