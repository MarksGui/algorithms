package BinaryTreeInorderTraversal

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
func TestIterating(t *testing.T) {
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
			want: []int{9, 3, 15, 20, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iterating(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Iterating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursive(t *testing.T) {
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
			want: []int{9, 3, 15, 20, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Recursive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Recursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
