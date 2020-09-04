package bitoperation

import (
	"fmt"
	"math"
)

// 题目描述: 求1+2+3+…+n，要求不能使用乘除法、for、while、if、else、switch、case
// 等关键字及条件判断语句 (A?B:C)

// 如果不让用乘除法，也不准用循环的话,可以用递归求和
func sum(n int) int {
	if n == 0 {
		return 0
	}
	return sum(n-1) + n
}

// 如果不能用关键字的话,可以用逻辑运算符代替三元运算符, 但 go 里面好像不支持这种写法
func sum1(n int) int {
	sum := n
	// var b bool
	// b = ( (n != 0) && (sum += (sum1(n-1)) != 0) )
	return sum
}

// 题目描述: 写一个函数，求两个整数之和，要求在函数体内不得使用+、-、*、/四则运算符号
func sum2(num1, num2 int) int {
	tmp := 0
	fmt.Printf("n1: %08b\n", num1)
	fmt.Printf("n2: %08b\n", num2)
	for num1 != 0 {
		tmp = num1 ^ num2
		num1 = (num1 & num2) << 1
		num2 = tmp
		fmt.Printf("n1: %08b\n", num1)
		fmt.Printf("n2: %08b\n", num2)
	}
	return num2
}

// 如果我们把两个数进行异或, 例如num1 = 101, num2 = 001, 做异或运算:
// tmp = num1 ^ num2, 结果是 tmp = 100. 那么此时得到的结果 tmp 其实就是两个数
// (num1,num2)各个二进制位上相加, 不算进位的结果. 而 num1 = (num1 & num2) << 1
// 的结果就是两个数相加时那些需要进位的二进制位. 例如 （101 & 001）<< 1 = 010,
// 那么两个数第一位相加需要进位, 我们需要把进的那一位最后加到第二位上去
// 说白就是 a + b = a ^ b + (a & b) << 1
// 如果 num1 == 0 的话, 代表没有进位了, 此时就可以退出循环了

// 题目描述：实现两个整数的相乘，不能使用乘法运算符和循环
func mul(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a == 1 {
		return b
	}
	if b == 1 {
		return a
	}

	return a + mul(a, b-1)
}

func mult(a, b int) int {
	sum := mul(a, int(math.Abs(float64(b))))
	if b < 0 {
		return -sum
	}
	return sum
}
