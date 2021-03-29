package path_sum_112

import "algorithms/tree"

// 路径总和
// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
// 叶子节点是指没有子节点的节点。
func hasPathSumIterating(root *tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queNode := []*tree.TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) > 0 {
		now := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]

		// 当到达叶子节点后（left、right为nil），判断值
		if now.Left == nil && now.Right == nil {
			if temp == targetSum {
				return true
			}
			continue
		}

		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+temp)
		}

		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+temp)
		}
	}
	return false
}

// 递归解法
func hasPathSumRecursive(root *tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSumRecursive(root.Left, targetSum-root.Val) || hasPathSumRecursive(root.Right, targetSum-root.Val)
}
