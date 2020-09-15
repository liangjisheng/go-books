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

	// list.head.next = list.Invert(list.head.next)
	// list.Print()

	// list.head.next = list.IterationInvert(list.head.next)
	// list.Print()

	// list.IterationInvertPartList(2, 4)
	// list.Print()

	// list.head = list.IterationInvertLinkedListEveryK(list.head, 2)
	// list.Print()
	// list.head.next = list.InvertLinkedListEveryK(list.head.next, 2)
	// list.Print()

	// list.ReverseIterationInvertLinkedListEveryK(2)
	// list.Print()

	// fmt.Printf("mid node: %d\n", list.findMiddleNodeMethod1().data)
	// fmt.Printf("mid node: %d\n", list.findMiddleNodeWithSlowFastPointer().data)

	// fmt.Printf("Kth node: %d\n", list.findKthToTailMethod1(1).data)
	// fmt.Printf("Kth node: %d\n", list.findKthToTailMethod2(5).data)

	// list.reversedKthToTail(2)
	// list.Print()

	// list.deleteKthToTail(1)
	// list.Print()

	// 构建环: 找到尾节点, 并将尾节点的 next 指向链表的第二个节点
	// tmp := list.head
	// for tmp.next != nil {
	// 	tmp = tmp.next
	// }
	// tmp.next = list.head.next.next

	// fmt.Printf("cross node: %d\n", list.detectCrossNode().data)
	// fmt.Printf("entry node: %d\n", list.getRingEntryNode().data)

	// fmt.Printf("palindromic: %v\n", list.isPalindromicList())
	// fmt.Printf("palindromic: %v\n", list.isPalindromicList2())
	// fmt.Printf("palindromic: %v\n", list.isPalindromicList3())

	// list1 := list.split1(3)
	// list1.Print()

	// list2 := list.split2(10)
	// list2.Print()

	// cycleList(list)
	// live := josephus1(list, 4)
	// fmt.Printf("%d live", live.data)
	// live := josephus2(list, 4)
	// fmt.Printf("%d live", live.data)
}
