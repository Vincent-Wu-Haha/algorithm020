package _2

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !found(wordList, endWord) {
		return 0
	}
	cur := []string{beginWord}
	visited := make([]bool, len(wordList))
	endCur := []string{endWord}
	endVisited := make([]bool, len(wordList))
	for i, s := range wordList {
		if s == beginWord {
			visited[i] = true
		}
		if s == endWord {
			endVisited[i] = true
		}
	}
	bLevel, eLevel := 0, 0
	for len(cur) != 0 && len(endCur) != 0 {
		i, done := accessNextLevel(&bLevel, &cur, wordList, endVisited, &eLevel, visited)
		if done {
			return i
		}
		j, done := accessNextLevel(&eLevel, &endCur, wordList, visited, &bLevel, endVisited)
		if done {
			return j
		}
	}
	return 0
}

func accessNextLevel(bLevel *int, cur *[]string, wordList []string, endVisited []bool, eLevel *int, visited []bool) (int, bool) {
	*bLevel++
	length := len(*cur)
	for i := 0; i < length; i++ {
		item := (*cur)[i]
		for i, s := range wordList {
			if !visited[i] && changeOnce(item, s) {
				visited[i] = true
				*cur = append(*cur, s)
			}
			if visited[i] && endVisited[i] {
				return *bLevel + *eLevel + 1, true
			}
		}
	}
	*cur = (*cur)[length:]
	return 0, false
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
