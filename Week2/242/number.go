package _42

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, i := range s {
		m[i]++
	}
	for _, i := range t {
		m[i]--
		if m[i] < 0 {
			return false
		} else if m[i] == 0 {
			delete(m, i)
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// map,k letter, v count
// 暴力，遍历S，删除T，T为空，S遍历完成
