package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	n := len(array)
	list := NewList()
	for i := 0; i < n; i++ {
		// list.HeadInsert(array[i])
		list.TailInsert(array[i])
	}
	fmt.Println("len:", list.length)
	list.Print()

	// list.InvertList()
	// list.Print()

	// list.IterationInvert()
	// list.Print()

	// list.IterationInvertPartList(2, 4)
	// list.Print()

	// list.IterationInvertLinkedListEveryK(2)
	// list.Print()

	// list.ReverseIterationInvertLinkedListEveryK(2)
	// list.Print()

	// fmt.Printf("mid node: %d\n", list.findMiddleNodeMethod1().data)
	// fmt.Printf("mid node: %d\n", list.findMiddleNodeWithSlowFastPointer().data)

	// fmt.Printf("Kth node: %d\n", list.findKthToTailMethod1(1).data)
	// fmt.Printf("Kth node: %d\n", list.findKthToTailMethod2(5).data)

	// list.reversedKthToTail(2)
	// list.Print()

	list.deleteKthToTail(1)
	list.Print()
}
