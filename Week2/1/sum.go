package _1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, i}
		} else {
			m[num] = i
		}
	}
	return nil
}

// 暴力，两遍loop
// 1遍，放到map,找减数
