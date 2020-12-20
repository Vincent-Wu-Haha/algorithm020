package _5

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1,0,2,3},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				nums: []int{1,1,2,3},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				nums: []int{0},
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				nums: []int{0,13},
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				nums: []int{2,0,2,3},
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				nums: []int{2,3,1,1,4},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
