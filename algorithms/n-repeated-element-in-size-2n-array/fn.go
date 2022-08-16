package n_repeated_element_in_size_2n_array

func repeatedNTimes(nums []int) int {
	numMap := make(map[int]int)
	for _, n := range nums {
		if _, exists := numMap[n]; exists {
			return n
		}
		numMap[n] = 1
	}
	return 0
}
