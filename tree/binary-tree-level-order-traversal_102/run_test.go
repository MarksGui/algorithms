package binary_tree_level_order_traversal_102

import (
	"algorithms/tree"
	"reflect"
	"testing"
)

//       3
//      / \
//     9  20
//       /  \
//      15   7

func Test_levelOrderIterating(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderIterating(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderIterating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrderIteratingL(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderIteratingL(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderIteratingL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrderRecursive(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
