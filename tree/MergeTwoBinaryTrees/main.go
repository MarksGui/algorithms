package main

import "algorithms/tree"

// 合并两颗二叉树（递归解法）
func MergeTreesRecursive(t1 *tree.TreeNode, t2 *tree.TreeNode) *tree.TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = MergeTreesRecursive(t1.Left, t2.Left)
	t1.Right = MergeTreesRecursive(t1.Right, t2.Right)
	return t1
}

// 合并两颗二叉树（迭代解法）
func MergeTreesRecursiveIterating(t1 *tree.TreeNode, t2 *tree.TreeNode) *tree.TreeNode {

	return nil
}
