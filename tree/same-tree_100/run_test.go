package same_tree_100

import (
	"algorithms/tree"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *tree.TreeNode
		q *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				p: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
				q: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				p: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
				q: tree.CreateCompleteBinaryTree(0, []int{5, 9, 20, 0, 0, 15, 7}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
