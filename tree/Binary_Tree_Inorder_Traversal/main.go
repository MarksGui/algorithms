package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归方法
func Recursive(root *TreeNode) []int {
	res := make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *[]int) {
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

// 迭代处理
func Iterating(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
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
	var root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(Recursive(root))
	fmt.Println(Iterating(root))
}
