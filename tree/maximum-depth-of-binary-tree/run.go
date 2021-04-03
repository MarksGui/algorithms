package maximum_depth_of_binary_tree

import "algorithms/tree"

// 二叉树的最大深度
func maxDepth(root *tree.TreeNode) int {
	var dep int
	if root == nil {
		return dep
	}
	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		dep++
	}

	return dep
}
