package main

// 将给定的二叉树左右反转
func reverseTree(root *TreeNode) {
	if root == nil {
		return
	}

	root.left, root.right = root.right, root.left

	reverseTree(root.left)
	reverseTree(root.right)
}

// 输入两棵二叉树A, B, 判断B是不是A的子结构 (ps:约定空树不是任意一个树的子结构)
// 判断 subRoot 是否包含在 root 的子结构中
func hasSubtree(root, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	// root 和 subRoot 代表的二叉树是一样的
	if judgeSubTree(root, subRoot) {
		return true
	}

	// 如果不一样, 则递归判断左右子树根节点是否和 subRoot 一样
	return judgeSubTree(root.left, subRoot) && judgeSubTree(root.right, subRoot)
}

// 判断 root, subRoot 代表的二叉树是否相同
func judgeSubTree(root, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if root.data == subRoot.data {
		// 如果根节点一样, 则递归判断左右子树
		return judgeSubTree(root.left, subRoot.left) && judgeSubTree(root.right, subRoot.right)
	}
	return false
}
