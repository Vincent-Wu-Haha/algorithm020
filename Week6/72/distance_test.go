package _2

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "1",
		//	args: args{
		//		word1: "horse",
		//		word2: "ros",
		//	},
		//	want: 3,
		//},
		{
			name: "1",
			args: args{
				word1: "dinitrophenylhydrazine",
				word2: "benzalphenylhydrazone",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}