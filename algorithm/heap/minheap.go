package main

import (
	"errors"
	"fmt"
)

// 索引从0开始算

// MinHeap ...
type MinHeap struct {
	size  int
	elems []int
}

// NewMinHeap ...
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func parentIndex(index int) int {
	if index == 0 {
		return 0
	}
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func (heap *MinHeap) print() {
	for _, v := range heap.elems {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func (heap *MinHeap) insert(key int) {
	heap.elems = append(heap.elems, key) // 将元素插入到最下堆中最后一个位置
	curI := heap.size                    // 当前节点索引
	parI := parentIndex(curI)            // 当前节点父节点索引

	// 如果当前插入的节点值小于父节点的值,则交换上浮
	// 交换后,原来的父节点变成当前节点,继续比较交换上浮
	for heap.elems[curI] < heap.elems[parI] {
		heap.elems[parI], heap.elems[curI] = heap.elems[curI], heap.elems[parI]
		curI = parI
		parI = parentIndex(curI)
	}
	heap.size++
}

// GetMinVal ...
func (heap *MinHeap) GetMinVal() (int, error) {
	if heap.size == 0 {
		return 0, errors.New("heap is empty")
	}
	return heap.elems[0], nil
}

// siftDown 从index的元素开始下沉调整,使之符合最小堆
func (heap *MinHeap) siftDown(index int) {
	curI := index
	var minI int

	for {
		leftI := leftChild(curI)
		rightI := rightChild(curI)

		// 当前节点已经到最底下没有左孩子返回
		if leftI > heap.size-1 {
			return
		}
		// 当前节点有左孩子,没有右孩子,比较当前节点和左孩子大小,并下沉
		if rightI > heap.size-1 {
			if heap.elems[curI] > heap.elems[leftI] {
				heap.elems[curI], heap.elems[leftI] = heap.elems[leftI], heap.elems[curI]
			}
			return
		}

		// 当前节点既有左孩子又有右孩子,获取值较小的孩子与当前节点进行比较,并下沉
		if heap.elems[leftI] < heap.elems[rightI] {
			minI = leftI
		} else {
			minI = rightI
		}
		if heap.elems[curI] > heap.elems[minI] {
			heap.elems[curI], heap.elems[minI] = heap.elems[minI], heap.elems[curI]
		}
		curI = minI
	}
}

// Replace 用key代替跟节点,并下沉调整使得符合最小堆
func (heap *MinHeap) Replace(key int) (int, error) {
	minVal, err := heap.GetMinVal()
	if err != nil {
		return minVal, err
	}

	heap.elems[0] = key
	heap.siftDown(0)
	return minVal, nil
}

// Sift 删除跟节点并返回,用最后一个节点代替,并下沉调整使得符合最小堆
func (heap *MinHeap) Sift() (int, error) {
	minVal, err := heap.GetMinVal()
	if err != nil {
		return minVal, err
	}

	heap.size--
	heap.elems[0], heap.elems[heap.size] = heap.elems[heap.size], 0
	heap.siftDown(0)
	return minVal, nil
}
