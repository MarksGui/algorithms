package main

import (
	"algorithms/tree"
	"container/list"
	"fmt"
)

// 二叉树的层序遍历（递归）
func levelOrderRecursive(root *tree.TreeNode) [][]int {
	res := make([][]int, 0)
	helper(root, &res)
	return res
}

func helper(root *tree.TreeNode, res *[][]int) {

}

func levelOrderIterating(root *tree.TreeNode) [][]int {
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

func main() {
	//       3
	//      / \
	//     9  20
	//       /  \
	//      15   7
	root := tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(levelOrderIterating(root))

}
