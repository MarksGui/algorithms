package n_ary_tree_postorder_traversal_590

import "algorithms/tree"

// N 叉树的后序遍历

// 递归
func postorderRecursive(root *tree.NTreeNode) []int {
	res := make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *tree.NTreeNode, res *[]int) {
	if root != nil {
		for _, node := range root.Children {
			helper(node, res)
		}
		*res = append(*res, root.Val)
	}
}

// 迭代
func postorder(root *tree.NTreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := []*tree.NTreeNode{root}

	for len(stack) > 0 {

	}

	return res
}
