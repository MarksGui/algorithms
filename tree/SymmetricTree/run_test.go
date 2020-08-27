package SymmetricTree

import (
	"algorithms/tree"
	"testing"
)

//       3
//      / \
//     10  10
//   /  \  /  \
//  7  15  15   7
func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				root: tree.CreateCompleteBinaryTree(0, []int{3, 10, 10, 7, 0, 0, 7}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
