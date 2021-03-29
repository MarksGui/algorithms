package n_ary_tree_postorder_traversal_590

import (
	"algorithms/tree"
	"reflect"
	"testing"
)

func Test_postorderRecursive(t *testing.T) {
	type args struct {
		root *tree.NTreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "递归",
			args: args{
				root: &tree.NTreeNode{
					Val: 1,
					Children: []*tree.NTreeNode{
						&tree.NTreeNode{
							Val: 3,
							Children: []*tree.NTreeNode{
								&tree.NTreeNode{
									Val:      5,
									Children: nil,
								},
								&tree.NTreeNode{
									Val:      6,
									Children: nil,
								},
							},
						},
						&tree.NTreeNode{
							Val:      2,
							Children: nil,
						},
						&tree.NTreeNode{
							Val:      4,
							Children: nil,
						},
					},
				},
			},
			want: []int{5, 6, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
