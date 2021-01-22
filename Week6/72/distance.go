package _2

func minDistance(word1 string, word2 string) int {
	cache := map[string]int{}
	return dp(word1, word2, cache)
}

func dp(word1 string, word2 string, cache map[string]int) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	if word1[len(word1)-1] == word2[len(word2)-1] {
		return dp(word1[:len(word1)-1], word2[:len(word2)-1], cache)
	}
	if v, ok := cache[gemKey(word1, word2)]; ok {
		return v
	}

	del := dp(word1[:len(word1)-1], word2, cache) + 1
	cache[gemKey(word1[:len(word1)-1], word2)] = del - 1

	insert := dp(word1, word2[:len(word2)-1], cache) + 1
	cache[gemKey(word1, word2[:len(word2)-1])] = insert - 1

	replace := dp(word1[:len(word1)-1], word2[:len(word2)-1], cache) + 1
	cache[gemKey(word1[:len(word1)-1], word2[:len(word2)-1])] = replace - 1

	return min(del, insert, replace)
}
func min(arr ...int) int {
	m := arr[0]
	for _, v := range arr {
		if m > v {
			m = v
		}
	}
	return m
}

func gemKey(w1, w2 string) string {
	return w1 + ":" + w2
}

// 带位置的最大公共子序列
