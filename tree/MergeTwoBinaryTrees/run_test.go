package MergeTwoBinaryTrees

import (
	"algorithms/tree"
	"reflect"
	"testing"
)

func TestMergeTreesRecursive(t *testing.T) {
	type args struct {
		t1 *tree.TreeNode
		t2 *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *tree.TreeNode
	}{
		{
			name: "test",
			args: args{
				t1: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
				t2: tree.CreateCompleteBinaryTree(0, []int{3, 9, 20, 0, 0, 15, 7}),
			},
			want: tree.CreateCompleteBinaryTree(0, []int{6, 18, 40, 0, 0, 30, 14}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTreesRecursive(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTreesRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
