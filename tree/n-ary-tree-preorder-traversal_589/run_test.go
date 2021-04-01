package n_ary_tree_preorder_traversal_589

import (
	"algorithms/tree"
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *tree.NTreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "迭代",
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
			want: []int{1, 3, 5, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderRecursive(t *testing.T) {
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
			want: []int{1, 3, 5, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderRecursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
