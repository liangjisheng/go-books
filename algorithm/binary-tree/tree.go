package main

import "fmt"

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

// Tree ...
type Tree struct {
	root          *TreeNode
	preOrderData  []int
	inOrderData   []int
	postOrderData []int
	length        int
}

// NewTree ...
func NewTree(root *TreeNode) *Tree {
	return &Tree{
		root:          root,
		preOrderData:  make([]int, 0),
		inOrderData:   make([]int, 0),
		postOrderData: make([]int, 0),
		length:        0,
	}
}

// 前序遍历: 根左右
func (tree *Tree) preOrder(root *TreeNode) {
	if root == nil {
		return
	}

	tree.preOrderData = append(tree.preOrderData, root.data)
	tree.preOrder(root.left)
	tree.preOrder(root.right)
}

// 中序遍历: 左根右
func (tree *Tree) inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	tree.inOrder(root.left)
	tree.inOrderData = append(tree.inOrderData, root.data)
	tree.inOrder(root.right)
}

// 后序遍历: 左右根
func (tree *Tree) postOrder(root *TreeNode) {
	if root == nil {
		return
	}

	tree.postOrder(root.left)
	tree.postOrder(root.right)
	tree.postOrderData = append(tree.postOrderData, root.data)
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
	preOrderData := []int{3, 9, 20, 15, 7}
	inOrderData := []int{9, 3, 15, 20, 7}
	// postOrderData := []int{9, 15, 7, 20, 3}

	root := buildTreeByPreAndIn(preOrderData, inOrderData)
	tree := NewTree(root)

	tree.preOrder(tree.root)
	fmt.Print("pre order: ")
	fmt.Println(tree.preOrderData)

	tree.inOrder(tree.root)
	fmt.Print("in order: ")
	fmt.Println(tree.inOrderData)

	tree.postOrder(tree.root)
	fmt.Print("post order: ")
	fmt.Println(tree.postOrderData)
}

func buildTreeDemoByInAndPost() {
	// preOrderData := []int{3, 9, 20, 15, 7}
	inOrderData := []int{9, 3, 15, 20, 7}
	postOrderData := []int{9, 15, 7, 20, 3}

	root := buildTreeByInAndPost(inOrderData, postOrderData)
	tree := NewTree(root)

	tree.preOrder(tree.root)
	fmt.Print("pre order: ")
	fmt.Println(tree.preOrderData)

	tree.inOrder(tree.root)
	fmt.Print("in order: ")
	fmt.Println(tree.inOrderData)

	tree.postOrder(tree.root)
	fmt.Print("post order: ")
	fmt.Println(tree.postOrderData)
}
