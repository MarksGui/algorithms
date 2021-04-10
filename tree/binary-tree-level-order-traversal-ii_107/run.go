package binary_tree_level_order_traversal_ii_107

import "algorithms/tree"

// 二叉树的程序遍历2
// 给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
func levelOrderBottom(root *tree.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		count := len(queue)
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append([][]int{level}, res...)
	}
	return res
}
