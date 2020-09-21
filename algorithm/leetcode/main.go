package main

import "fmt"

func main() {
	// array := []int{2, 3, 1, 0, 2, 5, 3}
	// res := duplicate(array)
	// fmt.Println(res)

	// array := []int{1, 1, 3, 6, 8, 8}
	// target := 9
	// res := twoSumOnlyOne(array, target)
	// fmt.Println(res)

	// resAll := twoSumAll(array, target)
	// fmt.Println(resAll)

	// resAll = twoSumAllNotRepeat(array, target)
	// fmt.Println(resAll)

	// array := []int{-1, 0, 1, 2, -1, -4}
	// resAll := threeSum(array)
	// fmt.Println(resAll)

	// array := []int{-1, 2, 1, -4}
	// sum := threeSumDiff(array, 1)
	// fmt.Println(sum)

	// alternatePrint()
	// alternatePrint1()

	// fmt.Println(lengthOfLongestSubstring1("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring2("pwwkew"))
	// fmt.Println(lengthOfLongestSubstring3("pwwkew"))

	// strs := []string{"flow", "flow", "flow"}
	// strs := []string{"dog", "racecar", "car"}
	// prefix := longestCommonPrefix(strs)
	// fmt.Println(prefix)

	// fmt.Printf("%5d = %s\n", 3, intToRoman(3))
	// fmt.Printf("%5d = %s\n", 4, intToRoman(4))
	// fmt.Printf("%5d = %s\n", 9, intToRoman(9))
	// fmt.Printf("%5d = %s\n", 58, intToRoman(58))
	// fmt.Printf("%5d = %s\n", 1994, intToRoman(1994))
	// fmt.Printf("%5d = %s\n", 9994, intToRoman(9994))

	// fmt.Println("intToRoman2")
	// fmt.Printf("%5d = %s\n", 3, intToRoman2(3))
	// fmt.Printf("%5d = %s\n", 4, intToRoman2(4))
	// fmt.Printf("%5d = %s\n", 9, intToRoman2(9))
	// fmt.Printf("%5d = %s\n", 58, intToRoman2(58))
	// fmt.Printf("%5d = %s\n", 1994, intToRoman2(1994))
	// fmt.Printf("%5d = %s\n", 9994, intToRoman2(9994))

	// fmt.Printf("%8s = %5d\n", "III", romanToInt("III"))
	// fmt.Printf("%8s = %5d\n", "IV", romanToInt("IV"))
	// fmt.Printf("%8s = %5d\n", "IX", romanToInt("IX"))
	// fmt.Printf("%8s = %5d\n", "LVIII", romanToInt("LVIII"))
	// fmt.Printf("%8s = %5d\n", "MCMXCIV", romanToInt("MCMXCIV"))

	// fmt.Println("romanToInt2")
	// fmt.Printf("%8s = %5d\n", "III", romanToInt2("III"))
	// fmt.Printf("%8s = %5d\n", "IV", romanToInt2("IV"))
	// fmt.Printf("%8s = %5d\n", "IX", romanToInt2("IX"))
	// fmt.Printf("%8s = %5d\n", "LVIII", romanToInt2("LVIII"))
	// fmt.Printf("%8s = %5d\n", "MCMXCIV", romanToInt2("MCMXCIV"))

	array := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("array :", array)

	fmt.Println(trap1(array))
	fmt.Println(trap2(array))
	fmt.Println(trap3(array))
}
