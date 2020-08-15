package sort

import "container/list"

// 平均时间复杂度: O(NlogN)
// 最佳时间复杂度: O(NlogN)
// 最差时间复杂度: O(N^2)
// 空间复杂度: 根据实现方式的不同而不同

// 场景分析: 递归是一种使用相同的方法, 通过解决问题的子集以达到解决整个问题的方法,
// 是一种使用有限代码解决“无限”计算的方法. 在C/C++语言中递归表现在函数对自身的直接/间接
// 的调用上,在实现上,递归依赖于语言的运行时调用堆栈, 使用堆栈来保存每一次递归调用返回
// 时所需要的条件.递归通常具有简洁的编码和清晰的思路,但这种简洁是有代价的.
// 一方面,是函数调用的负担;另一方面,是堆栈占用的负担(堆栈的大小是有限的)

// 改进思路: 递归转化为迭代. 迭代的思想主要在于, 在同一栈帧中不断使用现有数据
// 计算出新的数据, 然后使用新的数据来替换原有数据

// QuickSortRecursive ...
func QuickSortRecursive(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	quickSortRecursive(array, 0, n-1)
}

func quickSortRecursive(array []int, start, end int) {
	if start >= end {
		return
	}

	// 选择序列的最后一个元素作为基准点 pivot
	pivot := array[end]
	left := start
	right := end
	for left < right {
		for array[left] < pivot && left < right {
			left++
		}
		for array[right] >= pivot && left < right {
			right--
		}

		array[left], array[right] = array[right], array[left]
	}

	array[left], array[end] = array[end], array[left]

	quickSortRecursive(array, start, left-1)
	quickSortRecursive(array, left+1, end)
}

// Range ...
type Range struct {
	start, end int
}

// QuickSortIteration ...
func QuickSortIteration(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	iterations := list.New()
	iterations.PushBack(Range{
		start: 0,
		end:   n - 1,
	})

	for iterations.Len() > 0 {
		iter := iterations.Front()
		iterations.Remove(iter)
		start := iter.Value.(Range).start
		end := iter.Value.(Range).end
		if start >= end {
			continue
		}

		left := iter.Value.(Range).start
		right := iter.Value.(Range).end
		pivot := array[right]

		for left < right {
			for array[left] < pivot && left < right {
				left++
			}
			for array[right] >= pivot && left < right {
				right--
			}

			array[left], array[right] = array[right], array[left]
		}

		array[left], array[end] = array[end], array[left]

		iterations.PushBack(Range{
			start: start,
			end:   left - 1,
		})
		iterations.PushBack(Range{
			start: left + 1,
			end:   end,
		})
	}
}
