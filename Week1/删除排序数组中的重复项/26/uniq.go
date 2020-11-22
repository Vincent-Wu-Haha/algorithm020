package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	first := 0
	for second := 1; second < len(nums); second++ {
		for second < len(nums) && nums[first] == nums[second] {
			second++
		}
		if second >= len(nums) {
			break
		}
		first++
		nums[first] = nums[second]
	}
	return first + 1
}
