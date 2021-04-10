package binary_tree_level_order_traversal_102

import (
	"algorithms/tree"
	"container/list"
)

// 二叉树的层序遍历（递归）
func levelOrderRecursive(root *tree.TreeNode) [][]int {
	res := make([][]int, 0)
	helper(root, -1, &res)
	return res
}

func helper(node *tree.TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	currLevel := level + 1
	for len(*res) <= currLevel {
		*res = append(*res, []int{})
	}
	(*res)[currLevel] = append((*res)[currLevel], node.Val)
	helper(node.Left, currLevel, res)
	helper(node.Right, currLevel, res)
}

// 二叉树的层序遍历（迭代）
func levelOrderIterating(root *tree.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := make([]*tree.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		var curLevel []int
		count := len(queue)
		for count > 0 {
			node := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			count--
		}
		res = append(res, curLevel)
	}
	return res
}

// 二叉树的层序遍历（迭代-广度优先）  优化
func levelOrderIteratingNew(root *tree.TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*tree.TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		p := []*tree.TreeNode{}
		for j := 0; j < len(queue); j++ {
			cur := queue[j]
			res[i] = append(res[i], cur.Val)
			if cur.Left != nil {
				p = append(p, cur.Left)
			}
			if cur.Right != nil {
				p = append(p, cur.Right)
			}
		}
		queue = p
	}

	return res
}

// 二叉树的层序遍历（迭代）
// 使用内置 linked list
func levelOrderIteratingL(root *tree.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushFront(root)
	for queue.Len() > 0 {
		// 保存当前层数据的切片
		var curLevel []int
		// 保存当前层节点个数，使得新追加的元素不会被循环到
		count := queue.Len()
		for count > 0 {
			element := queue.Back()
			node := element.Value.(*tree.TreeNode)
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			queue.Remove(element)
			count--
		}
		res = append(res, curLevel)
	}

	return res
}
