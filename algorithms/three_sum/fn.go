package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	length := len(nums) - 2
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sums := nums[i] + nums[left] + nums[right]
			if sums < 0 {
				left += 1
			} else if sums > 0 {
				right -= 1
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left > right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			}
		}
	}

	return result
}
