package main

import (
	"fmt"
	"sort"
)

// 有一个班级有 n 个人, 给出 n 个元素, 第 i 个元素代表 第 i 位同学的考试成绩,
// 接下来进行 m 次询问, 每次询问给出一个数值 t, 表示第 t 个同学,
// 然后需要我们输出第 t 个同学的成绩超过班级百分之几的人, 百分数 p 可以这样算:
// p = (不超过第 t 个同学分数的人数 ) / n * 100% 输出的时候保留到小数点后 6 位
// 并且需要四舍五入

// 输入描述: 第一行输入两个数 n 和 m, 两个数以空格隔开, 表示 n 个同学和 m 次询问
// 第二行输入 n 个数值 ni, 表示每个同学的分数, 第三行输入 m 个数值mi,
// 表示每次询问是询问第几个同学 (注意，这里 2<=n, m<=100000, 0<=ni<=150, 1<=mi<=n)
// 输出描述: 输出 m 行, 每一行输出一个百分数 p, 代表超过班级百分之几的人

// 示例1:
// 输入:
// 3  2
// 50 60 70
// 1  2
// 输出
// 33.333333%
// 66.666667%

// 方法1 二分法, 就是先对所有人的成绩进行排序, 不过排序的时候我们需要开一个新的数组来存储
// 然后每次查询的时候可以通过二分查找进行匹配, 由于用到了排序, 需要花 O(nlogn) 的时间,
// m 次查询花的时间大致为 O(mlogn). 所以平均时间复杂度可以算是 O((m+n)logn)

func method1() {
	fmt.Print("input n, m: ")
	var n, m int
	fmt.Scanln(&n, &m)

	scores := make([]int, n)
	fmt.Print("input scores: ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &scores[i])
	}
	fmt.Scanln()
	fmt.Println("scores:", scores)
	sort.Ints(scores)
	fmt.Println("scores sort:", scores)

	asks := make([]int, m)
	fmt.Print("input ask index: ")
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &asks[i])
	}
	fmt.Scanln()
	fmt.Println("asks:", asks)

	fmt.Println("method1:")
	for i := 0; i < m; i++ {
		score := scores[asks[i]-1]
		fmt.Printf("score: %d ", score)
		index := binarySearch(scores, score) + 1
		res := float64(index) / float64(n) * 100
		fmt.Printf("%.6f%%\n", res)
	}

	fmt.Println("method2:")
	// method2(scores, asks)
}

// 寻找 target 的右边界
func binarySearch(array []int, target int) int {
	n := len(array)
	if n == 0 {
		return -1
	}

	left := 0
	right := n
	for left < right {
		mid := left + (right-left)/2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] == target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid
		}
	}

	// target 比数组中所有的数都小
	if left == 0 {
		return -1
	}

	if array[left-1] == target {
		return left - 1
	}

	// target 比数组中所有的数都大
	return -1
}

// 方法2 使用前缀和 假设每个同学的分数不超过 150, 不小于 0, 那么我们可以用一个数组 arr
// 然后让 arr[i] 表示分数不超过 i 的人数. 通过这种方式, 可以把时间复杂度控制在 O(n+m)
// scores[]存放成绩, m[]存放m次询问

func method2(scores []int, m []int) {
	prefixSum := make([]int, 151)
	n := len(scores)

	// 统计分数为 i 的有多少人
	for i := 0; i < n; i++ {
		prefixSum[scores[i]]++
	}

	// 构造前缀和, prefixSum[i] 表示分数不超过 i 的人有多少个
	for i := 1; i < 151; i++ {
		prefixSum[i] = prefixSum[i-1] + prefixSum[i]
	}

	// 进行 m 次询问
	askLen := len(m)
	for i := 0; i < askLen; i++ {
		score := scores[m[i]-1] // 取出分数
		fmt.Printf("score: %d ", score)
		sum := prefixSum[score] // 有多少人的成绩不超过 score
		precent := float64(sum) / float64(n) * 100
		fmt.Printf("%.6f%%\n", precent)
	}
}
