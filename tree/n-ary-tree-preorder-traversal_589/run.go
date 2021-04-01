package n_ary_tree_preorder_traversal_589

import "algorithms/tree"

// N 叉树的前序遍历 - 迭代
func preorder(root *tree.NTreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := []*tree.NTreeNode{root}

	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		res = append(res, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			queue = append(queue, cur.Children[i])
		}
	}

	return res
}

// N 叉树的前序遍历 - 递归
func preorderRecursive(root *tree.NTreeNode) []int {
	res := make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *tree.NTreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for _, child := range root.Children {
			helper(child, res)
		}
	}
}
