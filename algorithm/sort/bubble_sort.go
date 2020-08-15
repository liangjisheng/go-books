package sort

// 平均时间复杂度: O(N^2)
// 最佳时间复杂度: O(N)
// 最差时间复杂度: O(N^2)
// 空间复杂度: O(1)
// 排序方式: In-place
// 稳定性: 稳定

// 外层循环 n 次, 内层最多时循环 n – 1次、最少循环 0 次, 平均循环(n-1)/2

// 1.从头开始比较每一对相邻元素, 如果第1个比第2个大, 就交换它们的位置, 执行完一轮后, 最末尾那个元素就是最大的元素
// 2.忽略1中已经找到的最大元素, 重复执行步骤1, 直到全部元素有序
// 冒泡排序属于稳定排序

// BubbleSort ...
func BubbleSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := n - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
}

// BubbleSort1 如果序列已经完全有序,可以提前终止冒泡排序,相当于提前进行终止
func BubbleSort1(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := n - 1; i > 0; i-- {
		sorted := true
		for j := 1; j <= i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}
}

// BubbleSort2 如果序列尾部已经局部有序, 可以记录最后1次交换的位置, 减少比较次数
func BubbleSort2(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := n - 1; i > 0; {
		sortedIndex := 0
		for j := 1; j <= i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				sortedIndex = j
			}
		}
		i = sortedIndex
	}
}

// BubbleSort3 ...
func BubbleSort3(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
}

// BubbleSort4 在某次遍历中如果没有数据交换, 说明整个数组已经有序
func BubbleSort4(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		sorted := true
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
				sorted = false
			}
		}

		if sorted {
			return
		}
	}
}

// BubbleSort5 记录某次遍历时最后发生数据交换的位置pos, 这个位置之后的数据显然已经有序了
// 因此通过记录最后发生数据交换的位置就可以确定下次循环的范围了. 由于pos位置之后的记录均
// 已交换到位, 故在进行下一趟排序时只要扫描到pos位置即可
func BubbleSort5(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := n; i > 0; {
		pos := 0
		for j := 1; j < i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				pos = j
			}
		}

		i = pos
	}
}
