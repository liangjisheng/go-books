package main

func main() {
	topNTest()
}

func minHeapTest() {
	minHeap := NewMinHeap()
	for i := 9; i > 0; i-- {
		minHeap.insert(i)
	}
	minHeap.print()

	minHeap.insert(0)
	minHeap.print()
}
