package decode_xored_array

// decoded is 1720. Decode XORed Array
// SEE: https://leetcode.com/problems/decode-xored-array/
func decode(encoded []int, first int) []int {
	before := first
	result := make([]int, len(encoded)+1)
	result[0] = first
	for i, n := range encoded {
		raw := n ^ before
		result[i+1] = raw
		before = raw
	}
	return result
}
