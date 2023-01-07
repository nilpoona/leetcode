package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	before := nums[0]
	length := len(nums)
	k := 1
	for i := 1; i < length; i++ {
		if before != nums[i] {
			before = nums[i]
			nums = append(append(nums[:k], before), nums[k+1:]...)
			k++
		}
	}
	return k
}
