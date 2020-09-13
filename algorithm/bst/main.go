package main

import (
	"fmt"
)

func main() {
	bst := NewTree()
	bst.insert(50)
	bst.insert(38)
	bst.insert(60)
	bst.insert(20)
	bst.insert(40)
	bst.insert(52)
	bst.insert(62)
	bst.insert(53)
	bst.insert(48)
	bst.insert(54)
	bst.insert(63)

	fmt.Println("len:", bst.len)

	// fmt.Print("前序遍历: ")
	// bst.preOrder()
	// fmt.Print("前序遍历: ")
	// bst.preOrderNoRecursion()

	fmt.Print("中序遍历: ")
	bst.inOrder()
	// fmt.Print("中序遍历: ")
	// bst.inOrderNoRecursion()

	// fmt.Print("后序遍历: ")
	// bst.postOrder()
	// fmt.Print("后序遍历: ")
	// bst.postOrderNoRecursion()

	// bst.delete(60)
	// bst.inOrder()

	// val, err := bst.getLastCommonNode(54, 63)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("common node is:", val)

	// list := list.New()
	// bst.convert1(bst.root, list)
	// fmt.Print("list: ")
	// for i := 0; i < bst.len; i++ {
	// 	front := list.Front()
	// 	list.Remove(front)
	// 	fmt.Printf("%d ", front.Value.(*Node).key)
	// }
	// fmt.Println()

	fmt.Print("list: ")
	head := bst.convert2(bst.root).right
	for i := 0; i < bst.len && head != nil; i++ {
		fmt.Printf("%d ", head.key)
		head = head.right
	}
	fmt.Println()
}
