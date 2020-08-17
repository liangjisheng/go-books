package recursion

// 给定两个有序数组arr1和arr2, 已知两个数组的长度都为N, 求两个数组中所有数的上中位数
// 要求时间复杂度O(logN), 空间复杂度O(1)

// GetUpMedian ...
func GetUpMedian(array1 []int, array2 []int) int {
	n1 := len(array1)
	n2 := len(array2)
	if n1 == 0 || n2 == 0 {
		return -1
	}

	return find(array1, 0, n1-1, array2, 0, n2-1)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func find(array1 []int, l1, r1 int, array2 []int, l2, r2 int) int {
	mid1 := l1 + (r1-l1)/2
	mid2 := l2 + (r2-l2)/2

	// 表示数组只剩下一个数, 把两个数组中较小的数返回
	if l1 >= r1 {
		return min(array1[l1], array2[l2])
	}

	// 元素个数为奇数, 则offset为0, 为偶数则 offset 为 1
	// 用位运算比较快
	offset := (((r1 - l1) + 1) & 1) ^ 1
	if array1[mid1] > array2[mid2] {
		return find(array1, l1, mid1, array2, mid2+offset, r2)
	} else if array1[mid1] < array2[mid2] {
		return find(array1, mid1+offset, r1, array2, l2, mid2)
	} else {
		return array1[mid1] // 返回 array2[mid2] 也可以
	}
}

// GetUpMedianIteration ...
func GetUpMedianIteration(array1 []int, array2 []int) int {
	n1 := len(array1)
	n2 := len(array2)
	if n1 == 0 || n2 == 0 {
		return -1
	}

	l1 := 0
	r1 := n1 - 1
	l2 := 0
	r2 := n2 - 1
	mid1 := l1 + (r1-l1)/2
	mid2 := l2 + (r2-l2)/2
	offset := 0

	for l1 < r1 {
		mid1 = l1 + (r1-l1)/2
		mid2 = l2 + (r2-l2)/2
		offset = (((r1 - l1) + 1) & 1) ^ 1
		if array1[mid1] > array2[mid2] {
			r1 = mid1
			l2 = mid2 + offset
		} else if array1[mid1] < array2[mid2] {
			l1 = mid1 + offset
			r2 = mid2
		} else {
			return array1[mid1]
		}
	}

	return min(array1[l1], array2[l2])
}
