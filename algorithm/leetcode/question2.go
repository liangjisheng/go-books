package main

// 给定一个有序整数数组和一个目标值, 找出数组中和为目标值的两个数. 你可以假设每个输入
// 只对应一种答案, 且同样的元素不能被重复利用
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

// 解法1: 从左到右遍历数组, 在遍历的过程中, 取一个元素 a, 然后让 sum 减去 a, 这样可以得到
// b, 即 b = sum - a. 然后由于数组是有序的, 我们再利用二分查找, 在数组中查询 b 的下标
// 在这个过程中, 二分查找的时间复杂度是 O(logn), 从左到右扫描遍历是 O(n), 所以这种方法的
// 时间复杂度是 O(nlogn)

// 解法2: 采用双指针的方法, 从数组的头尾两边向中间夹击的方法来做的话, 时间复杂度仅需为 O(n)
// 而且代码也会更加简洁

func towSum(array []int, target int) []int {
	res := []int{-1, -1}
	start := 0
	end := len(array) - 1

	for start < end {
		if array[start]+array[end] < target {
			start++
		} else if array[start]+array[end] == target {
			res[0] = start
			res[1] = end
			return res
		} else if array[start]+array[end] > target {
			end--
		}
	}
	return res
}
