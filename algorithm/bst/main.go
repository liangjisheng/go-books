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

	fmt.Print("前序遍历: ")
	bst.preOrder()
	fmt.Print("前序遍历: ")
	bst.preOrderNoRecursion()

	fmt.Print("中序遍历: ")
	bst.inOrder()
	fmt.Print("中序遍历: ")
	bst.inOrderNoRecursion()

	fmt.Print("后序遍历: ")
	bst.postOrder()
	fmt.Print("后序遍历: ")
	bst.postOrderNoRecursion()

	// bst.delete(60)
	// bst.inOrder()

	val, err := bst.getLastCommonNode(54, 63)
	if err != nil {
		return
	}
	fmt.Println("common node is:", val)
}
