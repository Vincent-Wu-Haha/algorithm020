package _5

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	return help(nums)
}

func help(nums []int) bool {
	l := len(nums)
	min := -1
	flag := false
	for i := l - 2; i >= 0; i-- {
		if nums[i] >= l-i-1 {
			flag = true
			min = i
		}
	}
	if flag == true {
		if min == 0 {
			return true
		}
		return help(nums[:min+1])
	}
	return false
}
