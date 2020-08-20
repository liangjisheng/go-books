package recursion

import "fmt"

// 从n个不同元素中, 任取 m (m≤n,m与n均为自然数,下同)个不同的元素按照一定的顺序排成一列
// 叫做从n个不同元素中取出m个元素的一个排列, 当 n = m 时, 我们称这样的排列为全排列
// A(m, n) = n!/(n-m)!

// 从n个元素中任意选3个, 求所有的排列
func permutation(n int) {
	if n < 3 {
		return
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if i != j && i != k && j != k {
					fmt.Printf("%d%d%d\n", i, j, k)
				}
			}
		}
	}
}

// 求从 k 位开始的全排列, 数组 array 存的是参与全排列的 1 到 n 这些数字
func permutationRecursion(array []int, k int) {
	n := len(array)
	if n == 0 {
		return
	}

	if k == n-1 {
		fmt.Println(array)
	} else {
		for i := k; i < n; i++ {
			array[k], array[i] = array[i], array[k]
			permutationRecursion(array, k+1)
			array[k], array[i] = array[i], array[k]
		}
	}
}

// 返回值表示是否还有下一个排列数
func nextPermutation(array []int) bool {
	n := len(array)
	currentIndex := n - 1
	beforeIndex := n - 1 // 记录从右到左寻找第一个左邻小于右邻的数对应的索引
	isAllReverse := true // 是否存在从右到左第一个左邻小于右邻的数对应的索引

	// 1. 从右到左（从个位数往高位数）寻找第一个左邻小于右邻的数
	for ; currentIndex > 0; currentIndex-- {
		beforeIndex = currentIndex - 1
		if array[beforeIndex] < array[currentIndex] {
			isAllReverse = false
			break
		}
	}

	// 如果不存在, 说明这个数已经是字典排序法里的最大值,
	// 此时已经找到所有的全排列了, 直接打印即可
	if isAllReverse {
		return false
	}

	// 2. 再从右往左找第一个比第一步找出的数更大的数的索引
	firstLargeIndex := n - 1
	for ; firstLargeIndex > beforeIndex; firstLargeIndex-- {
		if array[firstLargeIndex] > array[beforeIndex] {
			break
		}
	}

	// 3. 交换 上述 1, 2 两个步骤中得出的两个数
	array[beforeIndex], array[firstLargeIndex] = array[firstLargeIndex], array[beforeIndex]

	// 4. 对 beforeIndex 之后的数进行排序, 这里用了快排
	QuickSortRecursive(array[beforeIndex+1:])
	return true
}

func permutationDict(array []int) {
	QuickSortRecursive(array)
	fmt.Println(array)
	for nextPermutation(array) {
		fmt.Println(array)
	}
}

// QuickSortRecursive ...
func QuickSortRecursive(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	quickSortRecursive(array, 0, n-1)
}

func quickSortRecursive(array []int, start, end int) {
	if start >= end {
		return
	}

	// 选择序列的最后一个元素作为基准点 pivot
	pivot := array[end]
	left := start
	right := end
	for left < right {
		for array[left] < pivot && left < right {
			left++
		}
		for array[right] >= pivot && left < right {
			right--
		}

		array[left], array[right] = array[right], array[left]
	}

	array[left], array[end] = array[end], array[left]

	quickSortRecursive(array, start, left-1)
	quickSortRecursive(array, left+1, end)
}
