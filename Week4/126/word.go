package _26

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if !found(wordList, endWord) {
		return nil
	}
	visited := make([]bool, len(wordList))
	result, found, nextLevel := help(endWord, wordList, [][]string{{beginWord}}, map[string]struct{}{beginWord: {}}, visited)
	for !found && len(nextLevel) != 0 {
		result, found, nextLevel = help(endWord, wordList, result, nextLevel, visited)
	}
	if found {
		return findRightPath(result, endWord)
	}
	return nil
}

func findRightPath(result [][]string, endWord string) (newResult [][]string) {
	for _, strings := range result {
		if strings[len(strings)-1] == endWord {
			newResult = append(newResult, strings)
		}
	}
	return
}

func help(endWord string, wordList []string, path [][]string, cur map[string]struct{}, visited []bool) (result [][]string, found bool, nextLevel map[string]struct{}) {
	for item := range cur {
		if item == endWord {
			return path, true, nil
		}
		for i, s := range wordList {
			if !visited[i] && changeOnce(item, s) {
				if s == endWord {
					found = true
				}
				if len(nextLevel) == 0 {
					nextLevel = map[string]struct{}{s: {}}
				} else {
					nextLevel[s] = struct{}{}
				}
				for _, strings := range path {
					if strings[len(strings)-1] == item {
						newString := make([]string, 0, len(strings)+1)
						newString = append(newString, strings...)
						newString = append(newString, s)
						result = append(result, newString)
					}
				}
			}
		}
	}
	for i, s := range wordList {
		for s2 := range nextLevel {
			if s == s2 {
				visited[i] = true
				break
			}
		}
	}
	return
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
