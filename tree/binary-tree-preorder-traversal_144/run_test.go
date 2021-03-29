package binary_tree_preorder_traversal_144

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
func TestPreorderTraversalIterating(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreorderTraversalIterating(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversalIterating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPreorderTraversalRecursive(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreorderTraversalRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
