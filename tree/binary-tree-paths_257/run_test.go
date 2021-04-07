package binary_tree_paths_257

import (
	"algorithms/tree"
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "深度优先迭代",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{1, 2, 3, 0, 5}),
			},
			want: []string{"1->2->5", "1->3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
