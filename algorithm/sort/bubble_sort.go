package sort

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

	for i := n - 1; i > 0; i-- {
		sortedIndex := 1
		for j := 1; j <= i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				sortedIndex = j
			}
		}
		i = sortedIndex
	}
}
