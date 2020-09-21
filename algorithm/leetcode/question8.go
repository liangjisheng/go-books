package main

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247485976&idx=2&sn=b3f3f69365e445e7bf2c661c2de66520&scene=21#wechat_redirect

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// 解法一 暴力法
// 时间复杂度 O(N^2), 空间复杂度 O(1)

func trap1(height []int) int {
	n := len(height)
	if n < 2 {
		return -1
	}
	res := 0                 // 能存储的总水量
	waters := make([]int, n) // 每个 i 能存储的雨水量

	// 索引为 0 和 n-1 个柱子作为边界不存储雨水
	for i := 1; i < n-1; i++ {
		lMax, rMax := 0, 0
		// 找到 i 左边最高的柱子 包括 i
		for j := i; j >= 0; j-- {
			lMax = max(height[j], lMax)
		}
		// 找到 i 右边最高的柱子 包括 i
		for j := i; j < n; j++ {
			rMax = max(height[j], rMax)
		}

		// 如果 i 是左边最高或者右边最高或者就是整个最高的, i 这个柱子是不能存储水的
		// 此时 water[i] 就是 0
		waters[i] = min(lMax, rMax) - height[i]
		res += waters[i]
	}

	// fmt.Println("waters:", waters)
	return res
}

// 解法二 备忘录优化
// 之前的暴力解法, 不是在每个位置 i 都要计算r_max和l_max吗?
// 直接把结果都缓存下来, 别傻不拉几的每次都遍历, 这时间复杂度不就降下来了嘛
// 开两个数组r_max和l_max充当备忘录, l_max[i]表示位置 i 左边最高的柱子高度
// r_max[i]表示位置 i 右边最高的柱子高度. 预先把这两个数组计算好, 避免重复计算
// 时间复杂度为 O(N), 已经是最优了, 但是空间复杂度是 O(N)

func trap2(height []int) int {
	n := len(height)
	if n < 2 {
		return -1
	}
	res := 0                 // 能存储的总水量
	waters := make([]int, n) // 每个 i 能存储的雨水量
	lMax := make([]int, n)   // 数组充当备忘录
	rMax := make([]int, n)

	// 初始化
	lMax[0] = height[0]
	rMax[n-1] = height[n-1]

	// 从左到右计算 lMax
	for i := 1; i < n; i++ {
		lMax[i] = max(height[i], lMax[i-1])
	}

	// 从右到左计算 rMax
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(height[i], rMax[i+1])
	}

	// 计算 waters 和 res
	for i := 1; i < n-1; i++ {
		waters[i] = min(lMax[i], rMax[i]) - height[i]
		res += waters[i]
	}

	// fmt.Println("waters:", waters)
	return res
}

// 解法三 双指针解法边走边算
// 时间复杂度为 O(N), 优化空间复杂度为 O(1)

func trap3(height []int) int {
	n := len(height)
	if n < 2 {
		return -1
	}
	res := 0                 // 能存储的总水量
	waters := make([]int, n) // 每个 i 能存储的雨水量
	left := 1
	right := n - 2

	lMax := height[0]
	rMax := height[n-1]

	for left <= right {
		// lMax 表示 height[0...left] 最高柱子高度
		lMax = max(height[left], lMax)
		// rMax 表示 height[right, n-1] 最高柱子高度
		rMax = max(height[right], rMax)

		// 计算 waters 和 res
		if rMax < lMax {
			// 右边最高高度比左边最高高度小, 计算 waters 和 res
			waters[right] = rMax - height[right]
			res += waters[right]
			right--
		} else {
			// 左边边最高高度比右边最高高度小或者相等, 计算 waters 和 res
			waters[left] = lMax - height[left]
			res += waters[left]
			left++
		}
	}

	// fmt.Println("waters:", waters)
	return res
}
