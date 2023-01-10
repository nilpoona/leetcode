package xor_operation_in_an_array

// xorOperation is 1486. XOR Operation in an Array
// https://leetcode.com/problems/xor-operation-in-an-array/
func xorOperation(n int, start int) int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = start + (2 * i)
	}
	if len(nums) == 1 {
		return nums[0]
	}
	result := nums[0]
	for _, num := range nums[1:] {
		result = result ^ num
	}
	return result
}
