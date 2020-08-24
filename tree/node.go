package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//    3
//   / \
//  9  20
//    /  \
//   15   7
// [3,9,20,0,0,15,7]
// 通过一维数组构建一颗完全二叉树
// i   数组的下标
// arr 输入数组
func CreateCompleteBinaryTree(i int, arr []int) *TreeNode {
	t := &TreeNode{}
	if arr[i] != 0 {
		t.Val = arr[i]
	} else {
		t = nil
		return t
	}

	if i < len(arr) && 2*i+1 < len(arr) {
		t.Left = CreateCompleteBinaryTree(2*i+1, arr)
	}
	if i < len(arr) && 2*i+2 < len(arr) {
		t.Right = CreateCompleteBinaryTree(2*i+2, arr)
	}
	return t
}
