package _26

import (
	"reflect"
	"testing"
)

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot","dot","dog","lot","log","cog"},
			},
			want: [][]string{
				{"hit","hot","dot","dog","cog"},
				{"hit","hot","lot","log","cog"},
			},
		},
		{
			name: "2",
			args: args{
				beginWord: "red",
				endWord:   "tax",
				wordList:  []string{"ted","tex","red","tax","tad","den","rex","pee"},
			},
			want: [][]string{
				{"red","ted","tad","tax"},
				{"red","ted","tex","tax"},
				{"red","rex","tex","tax"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
