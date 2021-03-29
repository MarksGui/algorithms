package binary_tree_zigzag_level_order_traversal_103

import (
	"algorithms/tree"
	"container/list"
)

// 二叉树的锯齿形层次遍历（迭代）
func zigzagLevelOrder(root *tree.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushFront(root)
	i := 1 // 层数
	for queue.Len() > 0 {
		var curLevel []int
		count := queue.Len()
		for count > 0 {
			var ele *list.Element
			var node *tree.TreeNode
			if i%2 == 0 {
				// 右
				ele = queue.Back()
				queue.Remove(ele)
				node = ele.Value.(*tree.TreeNode)
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
			} else {
				// 左
				ele = queue.Front()
				queue.Remove(ele)
				node = ele.Value.(*tree.TreeNode)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}
			curLevel = append(curLevel, node.Val)
			count--
		}
		res = append(res, curLevel)
		i++
	}

	return res
}
