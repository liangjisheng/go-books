package main

import (
	"sort"
)

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

// 找到第一个二元组就返回
func twoSumOnlyOne(array []int, target int) []int {
	// sort
	sort.Ints(array)

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

// 给定一个有序整数数组和一个目标值, 找出数组中和为目标值的两个数的所有组合
// 找到所有的二元组包括重复的
func twoSumAll(array []int, target int) [][]int {
	// sort
	sort.Ints(array)

	res := make([][]int, 0)
	i := 0
	j := len(array) - 1

	for i < j {
		if array[i]+array[j] < target {
			i++
		} else if array[i]+array[j] == target {
			tmp := make([]int, 0)
			tmp = append(tmp, i)
			tmp = append(tmp, j)
			res = append(res, tmp)
			i++
			j--
		} else if array[i]+array[j] > target {
			j--
		}
	}

	return res
}

// 给定一个有序整数数组和一个目标值, 找出数组中和为目标值的两个数的所有组合
// 找到所有的二元组包括重复的
func twoSumAllNotRepeat(array []int, target int) [][]int {
	// sort
	sort.Ints(array)

	res := make([][]int, 0)
	i := 0
	j := len(array) - 1

	for i < j {
		if array[i]+array[j] < target {
			i++
		} else if array[i]+array[j] == target {
			tmp := make([]int, 0)
			tmp = append(tmp, i)
			tmp = append(tmp, j)
			res = append(res, tmp)

			// 过滤掉重复的组合
			for i < j && array[i] == array[i+1] {
				i++
			}
			for i < j && array[j] == array[j-1] {
				j--
			}
			i++
			j--
		} else if array[i]+array[j] > target {
			j--
		}
	}

	return res
}

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247485359&idx=2&sn=78e779f614eff8a3d66132621670b9a7&chksm=ce404c7bf937c56ddf662f69a11f93a7e0afa991480f7c985ceb6fb775897a64bcaf8315c9f0&token=665909378&lang=zh_CN&scene=21#wechat_redirect

// 给定一个包含 n 个整数的数组 nums, 判断 nums 中是否存在三个元素 a, b, c,
// 使得 a + b + c = 0? 找出所有满足条件且不重复的三元组
// 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]
// 满足要求的三元组集合为:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]
// 题目中的三数之和, 即找出所有 a + b + c = 0的三元组, 其实和二元组差不多
// 我们可以锁定一个数, 然后把问题转换为两数之和, 即找出 a + b = -c 的所有二元组
// 其中的 -c 相当于我们的 target

func threeSum(array []int) [][]int {
	if len(array) == 0 {
		return nil
	}
	res := make([][]int, 0)
	// sort
	sort.Ints(array)
	n := len(array)

	for k := 0; k < n; k++ {
		// 优化一下, array[k] > 0的后, 那 k 以及 k 后面的数都大于 0,
		// 此时不可能会出现 a + b +c = 0 的三元组
		if array[k] > 0 {
			break
		}

		// 过滤掉重复的
		if k > 0 && array[k] == array[k-1] {
			continue
		}

		// 将问题转化为两数之和
		target := 0 - array[k]
		i := k + 1
		j := n - 1
		for i < j {
			if array[i]+array[j] < target {
				i++
			} else if array[i]+array[j] == target {
				tmp := make([]int, 0)
				tmp = append(tmp, array[k])
				tmp = append(tmp, array[i])
				tmp = append(tmp, array[j])
				res = append(res, tmp)

				// 过滤掉重复的组合
				for i < j && array[i] == array[i+1] {
					i++
				}
				for i < j && array[j] == array[j-1] {
					j--
				}
				i++
				j--
			} else if array[i]+array[j] > target {
				j--
			}
		}
	}

	return res
}

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247485364&idx=3&sn=f61017b0b29910a9a57fa96ae0d575b5&scene=21#wechat_redirect

// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target. 找出 nums 中的三个整数
// 使得它们的和与 target 最接近. 返回这三个数的和. 假定每组输入只存在唯一答案
// 例如给定数组 nums = [-1, 2, 1, -4], 和 target = 1.
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2)
// 这道题和三数之和类似, 两道题的解决思路几乎一样
// 只不过这道题需要不停着记录三个数的和与 target 之间的差

func threeSumDiff(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}

	sum := array[0] + array[1] + array[2] // 存三数之和
	diff := abs(sum, target)              // 存差值

	// res := make([][]int, 0)
	// sort
	sort.Ints(array)
	n := len(array)
	i, j, k := 0, 0, 0

	for ; k < n; k++ {
		// 过滤掉重复的
		if k > 0 && array[k] == array[k-1] {
			continue
		}

		// 将问题转化为两数之和与 target - array[k] 的差值问题
		tmp := target - array[k]
		i = k + 1
		j = n - 1
		for i < j {
			if array[i]+array[j] < tmp {
				// 当前差值较大, 更新差值
				if diff > tmp-(array[i]+array[j]) {
					diff = tmp - (array[i] + array[j])
					sum = array[k] + array[i] + array[j]
					// fmt.Println(k, i, j)
				}
				i++
			} else if array[i]+array[j] == tmp {
				sum = array[k] + array[i] + array[j]
				return sum

				// tmp := make([]int, 0)
				// tmp = append(tmp, array[k])
				// tmp = append(tmp, array[i])
				// tmp = append(tmp, array[j])
				// res = append(res, tmp)

				// 过滤掉重复的组合
				// for i < j && array[i] == array[i+1] {
				// 	i++
				// }
				// for i < j && array[j] == array[j-1] {
				// 	j--
				// }
				// i++
				// j--
			} else if array[i]+array[j] > tmp {
				// 当前差值较大, 更新差值
				if diff > (array[i]+array[j])-tmp {
					diff = (array[i] + array[j]) - tmp
					sum = array[k] + array[i] + array[j]
					// fmt.Println(k, i, j)
				}
				j--
			}
		}
	}

	return sum
}

func abs(x, y int) int {
	if x >= y {
		return x - y
	}
	return y - x
}
