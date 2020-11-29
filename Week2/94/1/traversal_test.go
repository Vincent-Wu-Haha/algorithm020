package _1

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: []int{1,3,2},
		},
		{
			name: "success",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{1},
		},
		{
			name: "success",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{
						Val:   3,
						Left:  nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
					Right: nil,
				},
			},
			want: []int{3,1,2},
		},
		{
			name: "success",
			args: args{
				root: &TreeNode{
					Val:   2,
					Left:  &TreeNode{
						Val:   3,
						Left:  &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
			},
			want: []int{1,3,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
