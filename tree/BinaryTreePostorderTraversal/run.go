package BinaryTreePostorderTraversal

import (
	"algorithms/tree"
)

// 二叉树后续遍历（递归）
func PostorderTraversalRecursive(root *tree.TreeNode) []int {
	res := make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *tree.TreeNode, res *[]int) {
	if root != nil {
		if root.Left != nil {
			helper(root.Left, res)
		}
		if root.Right != nil {
			helper(root.Right, res)
		}
		*res = append(*res, root.Val)
	}
}

// 二叉树后续遍历（迭代）
func PostorderTraversalIterating(root *tree.TreeNode) []int {
	stack := make([]*tree.TreeNode, 0)
	res := make([]int, 0)
	curr := root
	if curr == nil {
		return res
	}
	stack = append(stack, curr)
	for len(stack) > 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 插入数组头部
		res = append([]int{curr.Val}, res...)
		// 因为值是从头部入队，所以这里左节点先入栈，迟于右节点被遍历到插入
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}
	return res
}
