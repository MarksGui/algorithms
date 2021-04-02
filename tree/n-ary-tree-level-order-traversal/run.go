package n_ary_tree_level_order_traversal

import "algorithms/tree"

// N 叉树的层序遍历
func levelOrder(root *tree.NTreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*tree.NTreeNode{root}

	for len(queue) > 0 {
		level := make([]int, 0)
		count := len(queue)
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			level = append(level, cur.Val)
			for _, child := range cur.Children {
				queue = append(queue, child)
			}
		}
		res = append(res, level)
	}

	return res
}
