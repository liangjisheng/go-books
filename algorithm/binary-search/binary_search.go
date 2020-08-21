package search

// BinarySearch 搜索和目标值相等的数
func BinarySearch(array []int, target int) int {
	n := len(array)
	if n <= 0 {
		return -1
	}

	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// BinarySearchFirstLargeOrEqual 查找第一个不小于目标值的数, 目标数并不一定就出现在数组中
func BinarySearchFirstLargeOrEqual(array []int, target int) int {
	n := len(array)
	if n <= 0 {
		return -1
	}

	left := 0
	right := n
	for left < right {
		mid := left + (right-left)/2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] == target {
			right = mid
		} else if array[mid] > target {
			right = mid
		}
	}

	// target 比所有的数都大
	if left == n {
		return -1
	}

	return left
}

// BinarySearchLastLargeOrEqual 查找最后一个不小于目标值的数, 目标数并不一定就出现在数组中
func BinarySearchLastLargeOrEqual(array []int, target int) int {
	n := len(array)
	if n <= 0 {
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

	if array[max(left-1, 0)] < target {
		return left
	}
	return max(left-1, 0)
}

// BinarySearchFirstLarge 查找第一个大于目标值的数, 目标数并不一定就出现在数组中
func BinarySearchFirstLarge(array []int, target int) int {
	n := len(array)
	if n <= 0 {
		return -1
	}
	// 如果最后一个数小于等于目标值的话,返回-1
	if array[n-1] <= target {
		return -1
	}

	left := 0
	right := n
	for left < right {
		mid := left + (right-left)/2
		if array[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// 用二分法的思想求x的n次方
func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res *= x
		}
		x = x * x
		n = n >> 1
	}
	return res
}

// 搜索选择排序数组: 假设按照升序排序的数组在预先未知的某个点上进行了旋转
// (例如, 数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2])
// 搜索一个给定的目标值, 如果数组中存在这个目标值, 则返回它的索引, 否则返回 -1
// 你可以假设数组中不存在重复的元素. 你的算法时间复杂度必须是 O(log n) 级别

// 示例 1:
// 输入: nums = [4,5,6,7,0,1,2], target = 0
// 输出: 4

// 示例 2:
// 输入: nums = [4,5,6,7,0,1,2], target = 3
// 输出: -1

// RotatedBinarySearch ...
func RotatedBinarySearch(array []int, target int) int {
	n := len(array)
	if n == 0 {
		return -1
	}

	left := 0
	right := n - 1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == target {
			return mid
		}

		// 情况1: 如果中间元素在旋转点左侧
		if array[mid] >= array[left] {
			// target 如果位于中间元素的左侧
			if array[mid] > target && target >= array[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 情况2: 中间元素在旋转点的右侧
			// target 如果位于中间元素的右侧
			if array[mid] < target && target <= array[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
