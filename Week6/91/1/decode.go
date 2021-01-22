package _1

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	cur := 1
	pre := 0
	tmp := 0

	for i := 1; i < len(s); i++ {
		tmp = cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				if i == 1 {
					cur = 1
				} else {
					cur = pre
				}
			} else {
				return 0
			}
		} else {
			if s[i-1:i+1] > "10" && s[i-1:i+1] <= "26" && s[i-1:i+1] != "20" {
				if i == 1 {
					cur = 2
				} else {
					cur += pre
				}
			} else {
				cur = cur
			}
		}
		pre = tmp
	}
	return cur
}
