package two_sum

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		table[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := table[diff]; ok {
			if v == i {
				continue
			}
			return []int{i, table[diff]}
		}
	}

	return []int{}
}
