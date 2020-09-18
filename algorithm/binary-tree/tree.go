package main

import (
	"fmt"
)

// 前序遍历: 根左右
// 中序遍历: 左根右
// 后续遍历: 左右根

// TreeNode ...
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// NewTreeNode ...
func NewTreeNode(data int) *TreeNode {
	return &TreeNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

// 前序遍历: 根左右
func preOrderRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	result = append(result, root.data)

	leftRes := preOrderRecursion(root.left)
	result = append(result, leftRes...)

	rightRes := preOrderRecursion(root.right)
	result = append(result, rightRes...)

	return result
}

// 中序遍历: 左根右
func inOrderRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)

	leftRes := inOrderRecursion(root.left)
	result = append(result, leftRes...)

	result = append(result, root.data)

	rightRes := inOrderRecursion(root.right)
	result = append(result, rightRes...)

	return result
}

// 后序遍历: 左右根
func postOrderRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)

	leftRes := postOrderRecursion(root.left)
	result = append(result, leftRes...)

	rightRes := postOrderRecursion(root.right)
	result = append(result, rightRes...)

	result = append(result, root.data)
	return result
}

// 题意: 给你一个前序遍历和中序遍历, 你要构造出一个二叉树
// 示例:
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]

// 前序遍历第一位数字一定是这个二叉树的根结点
// 中序遍历中, 根结点讲序列分为了左右两个区间. 左边的区间是左子树的结点集合
// 右边的区间是右子树的结点集合

func buildTreeByPreAndIn(preOrder, inOrder []int) *TreeNode {
	// 用 map 存储中序遍历, 目的是查找方便. 因为我们从前序遍历找到根节点后
	// 还要寻找根节点在中序遍历的哪个位置
	inOrderMap := make(map[int]int)
	n := len(inOrder)
	for i := 0; i < n; i++ {
		inOrderMap[inOrder[i]] = i
	}

	return buildByPreAndIn(preOrder, inOrderMap, 0, n-1, 0)
}

// 5个参数: 前序序列, 中序的 map, 前序的开始, 前序的结束, 中序的开始
func buildByPreAndIn(preOrder []int, inOrderMap map[int]int,
	preStart int, preEnd int, inStart int) *TreeNode {
	// 递归结束条件
	if preStart > preEnd {
		return nil
	}

	// 前序序列的第一位是根节点
	root := NewTreeNode(preOrder[preStart])
	// 找到中序序列中, 根节点的索引 index
	rootIndex := inOrderMap[root.data]
	// 代表左子树的结点个数
	leftLen := rootIndex - inStart

	root.left = buildByPreAndIn(preOrder, inOrderMap, preStart+1, preStart+leftLen, inStart)
	root.right = buildByPreAndIn(preOrder, inOrderMap, preStart+leftLen+1, preEnd, rootIndex+1)
	return root
}

// 通过中序序列和后序序列构建二叉树
func buildTreeByInAndPost(inOrder, postOrder []int) *TreeNode {
	// 用 map 存储中序遍历, 目的是查找方便. 因为我们从后序遍历找到根节点后
	// 还要寻找根节点在中序遍历的哪个位置
	inOrderMap := make(map[int]int)
	n := len(inOrder)
	for i := 0; i < n; i++ {
		inOrderMap[inOrder[i]] = i
	}

	return buildByInAndPost(postOrder, inOrderMap, 0, n-1, 0)
}

// 5个参数: 后序序列, 中序的 map, 后序的开始, 后序的结束, 中序的开始
func buildByInAndPost(postOrder []int, inOrderMap map[int]int,
	postStart int, postEnd int, inStart int) *TreeNode {
	// 递归结束条件
	if postStart > postEnd {
		return nil
	}

	// 后序序列的最后一位是根节点
	root := NewTreeNode(postOrder[postEnd])
	// 找到中序序列中, 根节点的索引 index
	rootIndex := inOrderMap[root.data]
	// 代表左子树的结点个数
	leftLen := rootIndex - inStart

	root.left = buildByInAndPost(postOrder, inOrderMap, postStart, postStart+leftLen-1, inStart)
	root.right = buildByInAndPost(postOrder, inOrderMap, postStart+leftLen, postEnd-1, rootIndex+1)
	return root
}

func buildTreeDemoByPreAndIn() {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	// postOrder := []int{9, 15, 7, 20, 3}

	root := buildTreeByPreAndIn(preOrder, inOrder)

	preOrderData := preOrderRecursion(root)
	fmt.Print("pre order: ")
	fmt.Println(preOrderData)

	inOrderData := inOrderRecursion(root)
	fmt.Print(" in order: ")
	fmt.Println(inOrderData)

	postOrderData := postOrderRecursion(root)
	fmt.Print("post order: ")
	fmt.Println(postOrderData)
}

func buildTreeDemoByInAndPost() {
	// preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	postOrder := []int{9, 15, 7, 20, 3}

	root := buildTreeByInAndPost(inOrder, postOrder)

	preOrderData := preOrderRecursion(root)
	fmt.Print("pre order: ")
	fmt.Println(preOrderData)

	inOrderData := inOrderRecursion(root)
	fmt.Print(" in order: ")
	fmt.Println(inOrderData)

	postOrderData := postOrderRecursion(root)
	fmt.Print("post order: ")
	fmt.Println(postOrderData)
}
