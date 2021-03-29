package symmetric_tree_101

import "algorithms/tree"

// 对称二叉树（递归）
func isSymmetric(root *tree.TreeNode) bool {
	return check(root, root)
}

func check(p, q *tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
