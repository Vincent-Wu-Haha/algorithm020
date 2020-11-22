package _89

func rotate(nums []int, k int) {
	position := len(nums) - k%len(nums)
	copy(nums, append(nums[position:], nums[:position]...))
}
