package _1

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				if i == 1 {
					dp[i] = 1
				} else {
					dp[i] = dp[i-2]
				}
			} else {
				return 0
			}
		} else {
			if s[i-1:i+1] > "10" && s[i-1:i+1] <= "26" && s[i-1:i+1] != "20" {
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] += dp[i-2] + dp[i-1]
				}
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[len(s)-1]
}
