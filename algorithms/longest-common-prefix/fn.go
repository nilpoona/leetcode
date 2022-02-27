package longest_common_prefix

import "math"

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	return getLongestCommonPrefix(strs, 0, len(strs)-1)
}

func getLongestCommonPrefix(strs []string, start, end  int) string {
	if start == end {
		return strs[start]
	}
	mid := (start + end) / 2
	left := getLongestCommonPrefix(strs, start, mid)
	right := getLongestCommonPrefix(strs, mid + 1, end)
	return getPrefix(left, right)
}

func getPrefix(left, right string) string {
	min := int(math.Min(float64(len(left)), float64(len(right))))
	for i := 0; i < min; i++ {
		if string(left[i]) != string(right[i]) {
			return left[0:i]
		}
	}
	return left[0:min]
}