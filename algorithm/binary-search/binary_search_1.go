package search

// BinarySearch1 二分查找
func BinarySearch1(array []int, target int) int {
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
			return mid
		} else if array[mid] > target {
			right = mid - 1
		}
	}

	// target 比所有的数都大
	if left == n {
		return -1
	}

	// target 不在 array 中
	if array[left] == target {
		return left
	}
	return -1
}

// 比如说给你有序数组 nums = [1,2,2,2,3], target = 2, 二分查找算法返回的索引是 2, 没错
// 但是如果我想得到 target 的左侧边界, 即索引 1, 或者我想得到 target 的右侧边界, 即索引 3
// 你也许会说, 找到一个 target 索引, 然后向左或向右线性搜索不行吗? 可以, 但是不好,
// 因为这样难以保证二分查找对数级的复杂度了

// 寻找左侧边界的二分搜索, 返回对应的索引
func leftBound(array []int, target int) int {
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
			// 找到 target 时不要立即返回, 而是缩小「搜索区间」的上界 right,
			// 在区间 [left, mid) 中继续搜索, 即不断向左收缩, 达到锁定左侧边界的目的
			right = mid
		} else if array[mid] > target {
			right = mid
		}
	}

	// target 比所有的数都大
	if left == n {
		return -1
	}

	// target 不在 array 中
	if array[left] == target {
		// 也可以返回 right, 因为 for 循环结束的条件是 left == right
		return left
	}
	return -1
}

// 寻找右侧边界的二分搜索, 返回对应的索引
func rightBound(array []int, target int) int {
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
			// 当 nums[mid] == target 时, 不要立即返回, 而是增大「搜索区间」的下界 left
			// 使得区间不断向右收缩, 达到锁定右侧边界的目的
			left = mid + 1
		} else if array[mid] > target {
			right = mid
		}
	}

	// target 比所有的数都小
	if left == 0 {
		return -1
	}

	// target 不在 array 中
	if array[left-1] == target {
		// 也可以返回 right, 因为 for 循环结束的条件是 left == right
		return left - 1
	}
	return -1
}
