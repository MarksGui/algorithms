package InvertBinaryTree_226

import (
	"algorithms/tree"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *tree.TreeNode
	}{
		{
			name: "递归解法",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{4, 2, 7, 1, 3, 6, 9}),
			},
			want: tree.CreateCompleteBinaryTree(0, []int{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTreeRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invertTree1(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *tree.TreeNode
	}{
		{
			name: "迭代:广度优先",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{4, 2, 7, 1, 3, 6, 9}),
			},
			want: tree.CreateCompleteBinaryTree(0, []int{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTreeIterating(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
