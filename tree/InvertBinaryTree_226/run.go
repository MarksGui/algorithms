package InvertBinaryTree_226

import "algorithms/tree"

// 翻转二叉树
func invertTreeRecursive(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	left := invertTreeRecursive(root.Left)
	right := invertTreeRecursive(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 迭代
func invertTreeIterating(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	var queue []*tree.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		left := curNode.Left
		curNode.Left = curNode.Right
		curNode.Right = left
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}

	return root
}
