package search

// 时间复杂度: O(n*n)
func minEatSpeed(piles []int, H int) int {
	if len(piles) == 0 {
		return -1
	}

	// piles 数组的最大值
	maxSpeed := maxArray(piles)
	// 注意这个 for 循环, 就是在连续的空间线性搜索, 这就是二分查找可以发挥作用的标志
	// 由于我们要求的是最小速度, 所以可以用一个搜索左侧边界的二分查找来代替线性搜索, 提升效率
	for speed := 1; speed <= maxSpeed; speed++ {
		// 以 speed 是否能在 H 小时内吃完香蕉
		if canFinish(piles, speed, H) {
			return speed
		}
	}
	return maxSpeed
}

// 时间复杂度: O(nlogn)
func minEatSpeedOptimize(piles []int, H int) int {
	if len(piles) == 0 {
		return -1
	}

	// piles 数组的最大值
	maxSpeed := maxArray(piles)
	// 注意这个 for 循环, 就是在连续的空间线性搜索, 这就是二分查找可以发挥作用的标志
	// 由于我们要求的是最小速度, 所以可以用一个搜索左侧边界的二分查找来代替线性搜索, 提升效率
	left := 1
	right := maxSpeed + 1
	for left < right {
		mid := left + (right-left)/2
		// 如果以mid速度能在规定的时间吃完,则缩小右边界,寻找左边界
		if canFinish(piles, mid, H) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 以 speed 这个速度是否能在 H 小时内吃完香蕉
func canFinish(piles []int, speed int, H int) bool {
	n := len(piles)
	hours := 0
	// 首先求吃完所有的香蕉需要多少小时, 然后和给定的时间 H 比较
	for i := 0; i < n; i++ {
		hours += piles[i] / speed
		if piles[i]%speed != 0 {
			hours++
		}
	}
	// 如果超出了规定的时间返回 false
	if hours > H {
		return false
	}
	return true
}
