package BinaryTreePreorderTraversal

import (
	"algorithms/tree"
)

// 二叉树前序遍历（递归）
func PreorderTraversalRecursive(root *tree.TreeNode) []int {
	res := make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *tree.TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		if root.Left != nil {
			helper(root.Left, res)
		}
		if root.Right != nil {
			helper(root.Right, res)
		}
	}
}

// 二叉树前序遍历（迭代）
func PreorderTraversalIterating(root *tree.TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*tree.TreeNode, 0)
	curr := root
	if curr == nil {
		return res
	}
	stack = append(stack, curr)
	for len(stack) > 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

	}
	return res
}
