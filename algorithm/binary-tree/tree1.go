package main

import "container/list"

// 给定一个二叉树, 找出其最大深度
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数

// 使用递归 maxDepth(root) = max(maxDepth(root.left), maxDepth(root.right)) + 1
// 因为每个节点都需要被访问一次, 因此时间复杂度为 O(n)
// 空间复杂度
// 考虑到递归使用调用栈(call stack)的情况
// 最坏情况: 树完全不平衡. 例如每个节点都只有右节点, 此时将递归 n 次, 需要保持调用栈的存储为 O(n)
// 最好情况: 树完全平衡. 即树的高度为 log(n), 此时空间复杂度为 O(log(n))
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.left)
	rightDepth := maxDepth(root.right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// 二叉树的非递归前序遍历
// 采用一个栈来辅助, 把结果放入 result 中作为返回值, 具体步骤如下:
// 1、把二叉树的根节点 root 放进栈
// 2、如果栈不为空, 从栈中取出一个节点, 把该节点追加到result;
// 如果该节点的右子树不为空, 则把右节点放入栈; 如果该节点的左子树不为空,
// 则把左节点放入栈中
// 3、一直重复步骤 2 ,直到栈为空, 此时遍历结束
func preOrderIteration(root *TreeNode) []int {
	result := make([]int, 0)
	// 使用 list 模拟 stack
	stack := list.New()
	if root == nil {
		return nil
	}
	stack.PushBack(root)

	for stack.Len() != 0 {
		tmp := stack.Back()
		stack.Remove(tmp)
		result = append(result, tmp.Value.(*TreeNode).data)

		right := tmp.Value.(*TreeNode).right
		if right != nil {
			stack.PushBack(right)
		}
		left := tmp.Value.(*TreeNode).left
		if left != nil {
			stack.PushBack(left)
		}
	}

	return result
}

// 二叉树的非递归前序遍历, 同样借助于一个辅助栈
// 1、进入 while 循环, 接着把根节点及其所有左子节点放入栈中
// 2、从栈中取出一个节点, 把该节点放入容器的尾部; 如果该节点的右子节点不为空
// 则把右子节点及其右子节点的所有左子节点放入队列
// 3、一直重复步骤 2, 直到栈为空并且当前节点也为空则退出循环
func inOrderIteration(root *TreeNode) []int {
	result := make([]int, 0)
	stack := list.New()

	for root != nil || stack.Len() != 0 {
		if root != nil {
			stack.PushBack(root)
			root = root.left
		} else {
			root = stack.Back().Value.(*TreeNode)
			stack.Remove(stack.Back())
			result = append(result, root.data)
			root = root.right
		}
	}

	return result
}

// 二叉树的非递归后序遍历, 同样借助于一个辅助栈, 并把结果存到 list 中
// 1、把二叉树的根节点 root 放进栈
// 2、如果栈不为空, 从栈中取出一个节点, 把该节点插入到容器的头部
// 如果该节点的左子树不为空, 则把该左子树放入栈中; 如果该节点的右子树不为空
// 则把右子树放入栈中
// 3、一直重复步骤 2, 直到栈为空, 此时遍历结束
func postOrderIteration(root *TreeNode) *list.List {
	result := list.New()
	stack := list.New()
	if root == nil {
		return result
	}
	stack.PushBack(root)

	for stack.Len() != 0 {
		tmp := stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		// 放在第一个位置
		result.PushFront(tmp)
		if tmp.left != nil {
			stack.PushBack(tmp.left)
		}
		if tmp.right != nil {
			stack.PushBack(tmp.right)
		}
	}

	return result
}

func listToArray(list *list.List) []int {
	array := make([]int, 0)
	n := list.Len()

	for i := 0; i < n; i++ {
		tmp := list.Front()
		list.MoveToBack(tmp)
		array = append(array, tmp.Value.(*TreeNode).data)
	}

	return array
}

// 从上往下打印出二叉树的每个节点, 同层节点从左至右打印
// 这个像相当于二叉树四种遍历中的层序遍历了, 其思想是采用广度优先遍历,
// 借助一个辅助队列
// 1、把二叉树的根节点 root 放进队列
// 2、如果队列不为空, 取出队列的一个节点, 把这个节点的左右孩子放进队列中
// (为空的话就不用放), 然后打印这个节点
// 3、一直重复步骤 2, 直到队列为空, 此时遍历结束
func printFromTopToBottom(root *TreeNode) *list.List {
	result := list.New()
	queue := list.New()
	if root == nil {
		return result
	}

	queue.PushBack(root)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		result.PushBack(node)
		if node.left != nil {
			queue.PushBack(node.left)
		}
		if node.right != nil {
			queue.PushBack(node.right)
		}
	}

	return result
}

// 输入一个整数数组, 判断该数组是不是某二叉搜索树的后序遍历的结果
// 如果是则输出Yes, 否则输出No. 假设输入的数组的任意两个数字都互不相同

// 一般对于二叉树的题目大多数都可以使用递归来做的, 这道题也是用递归来多, 思路如下:
// 二叉搜索树的后序序列有个这样的特点: 序列的最后一个值为二叉树的根 root
// 二叉搜索树左子树值都比 root 小, 右子树值都比 root 大
// 所以我们可以这样:
// 1、确定找出 root
// 2、遍历序列(除去root结点), 找到第一个大于root的位置, 则该位置左边为左子树,
// 右边为右子树
// 3、遍历右子树, 若发现有小于root的值, 则是不符合二叉树搜索树的规则的, 则直接返回false
// 4、分别判断左子树和右子树是否仍是二叉搜索树(即递归步骤1、2、3)

func judge(array []int, left, right int) bool {
	// 递归结束条件, 只有一个节点
	if left >= right {
		return true
	}

	// 最右边的节点为根节点
	root := array[right]
	// 用来记录序列中第一个比根节点大节点的下标
	index := 0
	for i := left; i < right; i++ {
		// 找到根节点的右孩子
		if array[i] > root {
			index = i
			i++
			// 如果右子树中有比根节点还小的树的话, 显然是不成立的
			for i <= right-1 {
				if array[i] < root {
					return false
				}
				i++
			}
		}
	}

	// 递归检查左右子树
	return judge(array, left, index-1) && judge(array, index, right-1)
}
