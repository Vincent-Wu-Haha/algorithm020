package _05

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				preorder: []int{3, 2, 1, 4},
				inorder:  []int{1, 2, 3, 4},
			},
			want: []int{3, 2, 4, 1},
		},
		{
			name: "1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: []int{9, 15, 7, 20, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(postorder(got), tt.want) {
				t.Errorf("buildTree() = %v, want %v", postorder(got), tt.want)
			}
		})
	}
}

func postorder(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	result = append(result, postorder(root.Left)...)
	result = append(result, postorder(root.Right)...)
	result = append(result, root.Val)
	return result
}
