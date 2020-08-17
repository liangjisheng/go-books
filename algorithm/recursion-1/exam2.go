package recursion

// 给定两个有序数组arr1和arr2, 已知两个数组的长度分别为 m1 和 m2,
// 求两个数组中的第 K 小数. 要求时间复杂度O(log(m1 + m2))

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// FindKth ...
func FindKth(array1 []int, array2 []int, k int) int {
	n1 := len(array1)
	n2 := len(array2)
	if k > n1+n2 {
		return -1
	}

	if n1 == 0 {
		return array2[k-1]
	}
	if n2 == 0 {
		return array1[k-1]
	}

	// 这里我们假设k从0算起, 也就是有第0小元素, 相当于令 k = k - 1
	return findKth(array1, 0, n1-1, array2, 0, n2-1, k-1)
}

func findKth(array1 []int, l1, r1 int, array2 []int, l2, r2 int, k int) int {
	// 递归结束条件
	if l1 > r1 {
		return array2[l2+k]
	}
	if l2 > r2 {
		return array1[l1+k]
	}
	// 注意, k == 0的结束条件与上面两个结束条件不能颠倒
	if k == 0 {
		return min(array1[l1], array2[l2])
	}

	mid1 := min(l1+k/2, r1)
	mid2 := min(l2+k/2, r2)
	if array1[mid1] > array2[mid2] {
		return findKth(array1, l1, r1, array2, mid2+1, r2, k-k/2-1)
	} else if array1[mid1] < array2[mid2] {
		return findKth(array1, mid1+1, r1, array2, l2, r2, k-k/2-1)
	} else {
		return array1[mid1] // 返回 array2[mid2] 也可以
	}
}
