package _1

func ladderLength(beginWord string, endWord string, wordList []string) int {
	res := map[int]struct{}{}
	backtrack(beginWord, endWord, wordList, 0, res)
	return minOrNotExist(res) + 1
}

func backtrack(begin, end string, wordlist []string, count int, res map[int]struct{}) {
	if begin == end {
		res[count] = struct{}{}
		return
	}
	if len(wordlist) == 0 {
		return
	}

	for i, s := range wordlist {
		if changeOnce(begin, s) {
			backtrack(s, end, append(append([]string{}, wordlist[:i]...), wordlist[i+1:]...), count+1, res)
		}
	}
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

func minOrNotExist(res map[int]struct{}) int {
	m := -1
	for i := range res {
		if m == -1 {
			m = i
		}
		if i < m {
			m = i
		}
	}
	return m
}
