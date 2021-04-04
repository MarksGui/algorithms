package average_of_levels_in_binary_tree

import "algorithms/tree"

// 二叉树的层平均值
func averageOfLevels(root *tree.TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}
	queue := []*tree.TreeNode{root}

	for len(queue) > 0 {
		count := len(queue)
		var level int
		for i := 0; i < count; i++ {
			cur := queue[0]
			queue = queue[1:]
			level += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		avg := float64(level) / float64(count)
		res = append(res, avg)
	}

	return res
}
