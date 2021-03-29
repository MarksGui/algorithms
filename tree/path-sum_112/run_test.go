package path_sum_112

import (
	"algorithms/tree"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *tree.TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				root:      tree.CreateCompleteBinaryTree(0, []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 1}),
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				root:      tree.CreateCompleteBinaryTree(0, []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 1}),
				targetSum: 23,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSumIterating(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasPathSumRecursive(t *testing.T) {
	type args struct {
		root      *tree.TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				root:      tree.CreateCompleteBinaryTree(0, []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 1}),
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				root:      tree.CreateCompleteBinaryTree(0, []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 1}),
				targetSum: 23,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSumRecursive(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSumRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
