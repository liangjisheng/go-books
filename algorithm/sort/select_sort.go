package sort

// 1 从序列中找出最大的那个元素, 然后与最末尾的元素交换位置 执行完一轮后, 最末尾的那个元素就是最大的元素
// 2 忽略 1 中曾经找到的最大元素, 重复执行步骤 1
// 选择排序属于不稳定排序

func selectSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[i] {
				array[j], array[i] = array[i], array[j]
			}
		}
	}
}

func selectSort1(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		array[i], array[min] = array[min], array[i]
	}
}
