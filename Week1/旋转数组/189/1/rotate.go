package _89

func rotate(nums []int, k int) {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		copy(nums, append(nums[len(nums)-1:], nums[:len(nums)-1]...))
	}
}
