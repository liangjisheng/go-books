package main

import (
	"container/list"
)

// 对于二叉树的节点来说, 有本身的值域, 有指向左孩子和右孩子的两个指针;
// 对双向链表的节点来说, 有本身的值域, 有指向上一个节点和下一个节点的指针
// 在结构上, 两种结构有相似性, 现有一棵搜索二叉树, 请将其转为成一个有序的双向链表

// 解法一: 中序遍历 时间复杂度为 O(n), 空间复杂度也为 O(n)
func (bst *BinarySearchTree) convert1(root *Node, list *list.List) {
	if root == nil {
		return
	}

	bst.convert1(root.left, list)
	list.PushBack(root)
	bst.convert1(root.right, list)
}

// 解法二: 通过递归的方式 返回最后面的节点
// 时间复杂度为O(n),空间复杂度为O(h), 其中h是二叉树的高度
func (bst *BinarySearchTree) convert2(root *Node) *Node {
	if root == nil {
		return root
	}

	leftEnd := bst.convert2(root.left)
	rightEnd := bst.convert2(root.right)
	var leftBegin, rightBegin *Node
	if leftEnd != nil {
		leftBegin = leftEnd.right
	}
	if rightEnd != nil {
		rightBegin = rightEnd.right
	}

	if leftEnd != nil && rightEnd != nil {
		leftEnd.right = root
		root.left = leftEnd
		root.right = rightBegin
		rightBegin.left = root
		rightEnd.right = leftBegin
		return rightEnd
	} else if leftEnd != nil {
		leftEnd.right = root
		root.left = leftEnd
		root.right = leftBegin
		return root
	} else if rightEnd != nil {
		root.right = rightBegin
		rightBegin.left = root
		rightEnd.right = root
		return rightEnd
	} else if leftEnd == nil && rightEnd == nil {
		// 把最右边节点的 right 指针指向了最左边的节点的
		root.right = root
		return root
	}
	return nil
}
