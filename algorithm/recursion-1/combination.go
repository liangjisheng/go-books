package recursion

import "fmt"

// 从 n 个不同元素中每次取出 m 个不同元素, 不管其顺序合成一组,
// 称为从 n 个元素中不重复地选取 m 个元素的一个组合. 所有这样的组合的种数称为组合数
// C(m, n) = n! / (m!(n-m)!)

// COMBINATIONCNT 组合中需要被选中的个数
var COMBINATIONCNT = 2

func selectedNum(selectArray []byte) int {
	n := len(selectArray)
	res := 0
	for i := 0; i < n; i++ {
		res = res + int(selectArray[i])
	}
	return res
}

// 从数组 array 中第 k 个位置开始取 m 个元素
// 这里我们额外引入了一个 selectArray 数组, 这个数组里的元素如果为1,
// 则代表相应位置的元素被选中了, 如果为 0 代表未选中
func combination(array []int, k int, selectArray []byte) {
	n := len(array)

	selectNum := selectedNum(selectArray)
	// 终止条件1: 选中的元素已经等于我们要选择的数组个数了
	if selectNum == COMBINATIONCNT {
		for i := 0; i < n; i++ {
			if selectArray[i] == 1 {
				fmt.Printf("%d ", array[i])
			}
		}
		fmt.Println()
	} else {
		// 终止条件2: 开始选取的数组索引 超出数组范围了
		if k >= n {
			return
		}

		// 第 k 位被选中
		selectArray[k] = 1
		// 则从第 k+1 位选择 COMBINATIONCNT - selectNum 个元素
		combination(array, k+1, selectArray)

		// 第 k 位未被选中
		selectArray[k] = 0
		// 则从第 k+1 位选择 COMBINATIONCNT - selectNum 个元素
		combination(array, k+1, selectArray)
	}
}
