package sort

// 平均时间复杂度: O(nlogn)
// 最佳时间复杂度: O(nlogn)
// 最差时间复杂度: O(nlogn)
// 空间复杂度: O(n)
// 排序方式: In-place
// 稳定性: 稳定

// 不管元素在什么情况下都要做这些步骤, 所以花销的时间是不变的,
// 所以该算法的最优时间复杂度和最差时间复杂度及平均时间复杂度都是一样的
// 归并的空间复杂度就是那个临时的数组和递归时压入栈的数据占用的空间: n + logn;
// 所以空间复杂度为: O(n)
// 归并排序算法中, 归并最后到底都是相邻元素之间的比较交换,
// 并不会发生相同元素的相对位置发生变化, 故是稳定性算法

// 递归法
// 1.将序列每相邻两个数字进行归并操作, 形成floor(n/2)个序列, 排序后每个序列包含两个元素
// 2.将上述序列再次归并, 形成floor(n/4)个序列, 每个序列包含四个元素
// 3.重复步骤2, 直到所有元素排序完毕

// MergeSortRecursive ...
func MergeSortRecursive(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	reg := make([]int, n)
	mergeSortRecursive(array, reg, 0, n-1)
}

func mergeSortRecursive(array []int, reg []int, start, end int) {
	if start >= end {
		return
	}

	len := end - start
	mid := len/2 + start
	start1 := start
	end1 := mid
	start2 := mid + 1
	end2 := end

	mergeSortRecursive(array, reg, start1, end1)
	mergeSortRecursive(array, reg, start2, end2)

	k := start
	for start1 <= end1 && start2 <= end2 {
		if array[start1] <= array[start2] {
			reg[k] = array[start1]
			start1++
		} else {
			reg[k] = array[start2]
			start2++
		}
		k++
	}

	for start1 <= end1 {
		reg[k] = array[start1]
		k++
		start1++
	}

	for start2 <= end2 {
		reg[k] = array[start2]
		k++
		start2++
	}

	for k = start; k <= end; k++ {
		array[k] = reg[k]
	}
}

// 迭代法
// 1.申请空间, 使其大小为两个已经排序序列之和, 该空间用来存放合并后的序列
// 2.设定两个指针, 最初位置分别为两个已经排序序列的起始位置
// 3.比较两个指针所指向的元素, 选择相对小的元素放入到合并空间, 并移动指针到下一位置
// 4.重复步骤3直到某一指针到达序列尾
// 5.将另一序列剩下的所有元素直接复制到合并序列尾

// Range ...
type Range struct {
	start1, end1 int
	start2, end2 int
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// MergeSortIteration ...
func MergeSortIteration(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	reg := make([]int, n)
	seg := 0
	start := 0

	// 这个for是分2路归并,4路归并,8路归并...
	for seg = 1; seg < n; seg += seg {
		// 这个for是对多个2路进行归并
		for start = 0; start < n; start += seg + seg {
			start1 := start
			end1 := min(start+seg, n)
			start2 := min(start+seg, n)
			end2 := min(start+seg+seg, n)

			k := start
			for start1 < end1 && start2 < end2 {
				if array[start1] <= array[start2] {
					reg[k] = array[start1]
					start1++
				} else {
					reg[k] = array[start2]
					start2++
				}
				k++
			}

			for start1 < end1 {
				reg[k] = array[start1]
				k++
				start1++
			}

			for start2 < end2 {
				reg[k] = array[start2]
				k++
				start2++
			}

			for k = start; k < end2; k++ {
				array[k] = reg[k]
			}
		}
	}
}
