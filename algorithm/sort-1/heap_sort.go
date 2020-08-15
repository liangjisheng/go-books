package sort

// 平均时间复杂度: O(nlogn)
// 最佳时间复杂度: O(nlogn)
// 最差时间复杂度: O(nlogn)
// 稳定性: 不稳定

// 堆排序其实也是一种选择排序, 是一种树形选择排序. 只不过直接选择排序中
// 为了从R[1…n]中选择最大记录, 需比较n-1次, 然后从R[1…n-2]中选择最大记录需比较n-2次
// 事实上这n-2次比较中有很多已经在前面的n-1次比较中已经做过, 而树形选择排序恰好利用
// 树形的特点保存了部分前面的比较结果, 因此可以减少比较次数. 对于n个关键字序列
// 最坏情况下每个节点需比较log2(n)次, 因此其最坏情况下时间复杂度为nlogn
// 堆排序为不稳定排序, 不适合记录较少的排序

// 以start为索引的父节点下沉, 使之符合最大堆
func maxHeapify(array []int, start, end int) {
	par := start
	son := par*2 + 1
	for son <= end {
		// 先比较两个子节点大小, 选择最大的
		if son+1 <= end && array[son] < array[son+1] {
			son++
		}

		// 如果父节点大于子节点代表调整完毕, 直接跳出函数
		if array[par] > array[son] {
			return
		}

		// 否则交换父子内容再继续子节点和孙节点比较
		array[par], array[son] = array[son], array[par]
		par = son
		son = par*2 + 1
	}
}

// HeapSort ...
func HeapSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	// 初始化, i从最后一个父节点开始调整
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(array, i, n-1)
	}

	// 先将第一个元素和已排好元素前一位做交换, 再从新调整, 直到排序完毕
	for i := n - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		// 这里用i-1意思是最后有序的元素不参与调整
		maxHeapify(array, 0, i-1)
	}
}
