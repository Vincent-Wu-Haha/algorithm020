package _53

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] >  nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
