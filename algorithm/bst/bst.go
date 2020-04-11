package main

import (
	"fmt"

	"github.com/liyue201/gostl/ds/stack"
)

// Node ...
type Node struct {
	key         int
	left, right *Node
}

// BinarySearchTree ...
type BinarySearchTree struct {
	root *Node
}

// NewNode ...
func NewNode(key int) *Node {
	return &Node{
		key: key,
	}
}

// NewTree ...
func NewTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) search(key int) bool {
	return _search(bst.root, key)
}

func _search(node *Node, key int) bool {
	if node != nil {
		if key == node.key {
			return true
		}
		if key < node.key {
			return _search(node.left, key)
		}
		return _search(node.right, key)
	}
	return false
}

func (bst *BinarySearchTree) insert(key int) *BinarySearchTree {
	bst.root = _insert(bst.root, key)
	return bst
}

func _insert(root *Node, key int) *Node {
	if root == nil {
		return NewNode(key)
	}
	if key < root.key {
		root.left = _insert(root.left, key)
	} else if key > root.key {
		root.right = _insert(root.right, key)
	}
	return root
}

func (bst *BinarySearchTree) inOrder() {
	_inOrder(bst.root)
	fmt.Println()
}

func _inOrder(node *Node) {
	if node != nil {
		_inOrder(node.left)
		fmt.Printf("%d ", node.key)
		_inOrder(node.right)
	}
}

func (bst *BinarySearchTree) preOrder() {
	_preOrder(bst.root)
	fmt.Println()
}

func _preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.key)
		_preOrder(node.left)
		_preOrder(node.right)
	}
}

func (bst *BinarySearchTree) postOrder() {
	_postOrder(bst.root)
	fmt.Println()
}

func _postOrder(node *Node) {
	if node != nil {
		_postOrder(node.left)
		_postOrder(node.right)
		fmt.Printf("%d ", node.key)
	}
}

func minNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	cur := node
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func (bst *BinarySearchTree) delete(key int) {
	bst.root = _delete(bst.root, key)
}

// 1.删除的节点为叶子节点: 直接删除
// 2.删除的节点只存在左子树或右子树: 删除节点的父节点直接指向子树节点
// 3.删除的节点同时存在左子树和右子树: 将删除节点的左子树的最右节点或右子树的最左节点替换删除节点
//   再更新删除节点的左子树或者右子树
func _delete(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = _delete(root.left, key)
	} else if key > root.key {
		root.right = _delete(root.right, key)
	} else {
		// this is the node to delete, node with one child
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			// node with two children
			node := minNode(root.right)
			root.key = node.key
			root.right = _delete(root.right, node.key)
		}
	}
	return root
}

func (bst *BinarySearchTree) inOrderNoRecursion() {
	_inOrderNoRecursion(bst.root)
}

func _inOrderNoRecursion(root *Node) {
	if root == nil {
		return
	}
	node := root
	s := stack.New()

	for !s.Empty() || node != nil {
		for node != nil {
			s.Push(node)
			node = node.left
		}

		if !s.Empty() {
			node = s.Pop().(*Node)
			fmt.Printf("%d ", node.key)
			node = node.right
		}
	}
	fmt.Println()
}

func (bst *BinarySearchTree) preOrderNoRecursion() {
	_preOrderNoRecursion(bst.root)
}

func _preOrderNoRecursion(root *Node) {
	if root == nil {
		return
	}

	node := root
	s := stack.New()

	for !s.Empty() || node != nil {
		for node != nil {
			fmt.Printf("%d ", node.key)
			s.Push(node)
			node = node.left
		}

		if !s.Empty() {
			node = s.Pop().(*Node)
			node = node.right
		}
	}
	fmt.Println()
}

func (bst *BinarySearchTree) postOrderNoRecursion() {
	_postOrderNoRecursion(bst.root)
}

func _postOrderNoRecursion(root *Node) {
	if root == nil {
		return
	}

	s := stack.New()
	// cur:当前访问节点, lastVisit:上次访问节点
	cur := root
	var lastVisit *Node

	// 先把pCur移动到左子树最下边
	for cur != nil {
		s.Push(cur)
		cur = cur.left
	}

	for !s.Empty() {
		// 走到这里，pCur都是空，并已经遍历到左子树底端(看成扩充二叉树，则空，亦是某棵树的左孩子)
		cur = s.Pop().(*Node)
		// 一个根节点被访问的前提是：无右子树或右子树已被访问过
		if cur.right == nil || cur.right == lastVisit {
			fmt.Printf("%d ", cur.key)
			// 修改最近被访问的节点
			lastVisit = cur
		} else {
			/*这里的else语句可换成带条件的else if:
			else if cur->left == lastVisit //若左子树刚被访问过，则需先进入右子树(根节点需再次入栈)
			因为：上面的条件没通过就一定是下面的条件满足。仔细想想*/
			// 根节点再次入栈
			s.Push(cur)
			// 进入右子树，且可肯定右子树一定不为空
			cur = cur.right
			for cur != nil {
				s.Push(cur)
				cur = cur.left
			}
		}
	}
	fmt.Println()
}
