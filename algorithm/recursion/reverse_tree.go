package recursion

// 反转二叉树 将左边的二叉树反转成右边的二叉树

// TreeNode ...
type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

// 时间复杂度分析 由于我们会对每一个节点都去做翻转, 所以时间复杂度是 O(n)
// 空间复杂度非常有意思, 我们一起来看下, 由于每次调用 ReverseTree 函数都相当于一次压栈操作
// 那最多压了几次栈呢? 从根节点出发不断对左结果调用翻转函数, 直到叶子节点,
// 每调用一次都会压栈, 左节点调用完后, 出栈, 再对右节点压栈...., 如果是完全二叉树,
// 则树的高度为logn, 即空间复杂度为O(logn), 最坏情况, 如果此二叉树只有左节点,
// 没有右节点, 则树的高度即结点的个数 n, 此时空间复杂度为 O(n),总的来看空间复杂度为O(n)

// ReverseTree ...
func ReverseTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := ReverseTree(root.left)
	right := ReverseTree(root.right)

	root.left = right
	root.right = left
	return root
}
