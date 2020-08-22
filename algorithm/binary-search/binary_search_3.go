package search

func minLoadCap(weights []int, D int) int {
	n := len(weights)
	if n == 0 {
		return -1
	}

	// 货物不可分割, 所以load 的最小值和最大值分别为 max(weights) 和 sum(weights)
	minLoad := maxArray(weights)
	maxLoad := sumArray(weights)
	for load := minLoad; load <= maxLoad; load++ {
		if canLoad(weights, load, D) {
			return load
		}
	}
	return maxLoad
}

func minLoadCapOptisize(weights []int, D int) int {
	n := len(weights)
	if n == 0 {
		return -1
	}

	left := maxArray(weights)
	right := sumArray(weights) + 1
	for left < right {
		mid := left + (right-left)/2
		if canLoad(weights, mid, D) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 返回以每天 load 的速度是否能在 D 天内将所有 weights 装上船
func canLoad(weights []int, load int, D int) bool {
	days := 0
	n := len(weights)
	for i := 0; i < n; {
		tmp := load
		for i < n {
			tmp -= weights[i]
			if tmp < 0 {
				break
			}
			i++
		}
		days++
	}

	// 说明以 load 的速度在 D 天内不能将货物装上船
	if days > D {
		return false
	}
	return true
}
