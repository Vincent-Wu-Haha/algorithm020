package _27

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !found(wordList, endWord) {
		return 0
	}
	cur := []string{beginWord}
	visited := make([]bool, len(wordList))
	level := 0
	for len(cur) != 0 {
		level++
		length := len(cur)
		for i := 0; i < length; i++ {
			item := cur[i]
			if item == endWord {
				return level
			}
			for i, s := range wordList {
				if !visited[i] && changeOnce(item, s) {
					visited[i] = true
					cur = append(cur, s)
				}
			}
		}
		cur = cur[length:]
	}
	return 0
}

func changeOnce(a, b string) bool {
	count := 0
	for i := range a {
		if a[i] != b[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	if count == 1 {
		return true
	}
	return false
}

func found(arr []string, s string) bool {
	for _, s2 := range arr {
		if s == s2 {
			return true
		}
	}
	return false
}
