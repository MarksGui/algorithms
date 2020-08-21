package main

import (
	"algorithms/tree"
	"fmt"
)

// 二叉树的中序遍历（递归）
func Recursive(root *tree.TreeNode) []int {
	res := make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *tree.TreeNode, res *[]int) {
	if root != nil {
		if root.Left != nil {
			helper(root.Left, res)
		}
		*res = append(*res, root.Val)
		if root.Right != nil {
			helper(root.Right, res)
		}
	}
}

// 二叉树的中序遍历（迭代）
func Iterating(root *tree.TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*tree.TreeNode, 0)
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

func main() {
	root := tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(Recursive(root))
	fmt.Println(Iterating(root))
}
