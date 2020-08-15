package sort

// 平均时间复杂度: O(N^2)
// 最佳时间复杂度: O(N^2)
// 最差时间复杂度: O(N^2)
// 空间复杂度: O(1)
// 排序方式: In-place
// 稳定性: 不稳定

// 选择排序的交换操作介于和(n-1)次之间. 选择排序的比较操作为n(n-1)/2次之间
// 选择排序的赋值操作介于0和3(n-1)次之间
// 比较次数O(n^2), 比较次数与关键字的初始状态无关, 总的比较次数
// N = (n-1) + (n-2) +…+ 1 = n x (n-1)/2. 交换次数O(n), 最好情况是
// 已经有序, 交换0次; 最坏情况是, 逆序, 交换n-1次

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
