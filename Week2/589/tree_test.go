package _89

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name:       "1",
			args:       args{
				root: &Node{
					Val:      1,
					Children: []*Node{
						{
							Val:      2,
							Children: []*Node{
								{
									Val:      3,
									Children: nil,
								},
							},
						},
						{
							Val:      4,
							Children: nil,
						},
					},
				},
			},
			wantResult: []int{1,2,3,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := preorder(tt.args.root); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("preorder() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
