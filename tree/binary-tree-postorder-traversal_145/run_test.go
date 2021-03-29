package binary_tree_postorder_traversal_145

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
func TestPostorderTraversalIterating(t *testing.T) {
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
			want: []int{9, 15, 7, 20, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostorderTraversalIterating(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversalIterating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostorderTraversalRecursive(t *testing.T) {
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
			want: []int{9, 15, 7, 20, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PostorderTraversalRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversalRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
