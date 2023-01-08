package find_the_index_of_the_first_occurrence_in_a_string

func find(haystack, needle []rune, index int) int {
	if len(needle) > len(haystack) || len(needle) == 0 {
		return -1
	}

	if len(haystack) == index || index+len(needle) > len(haystack) {
		return -1
	}

	found := true
	count := 0
	for _, code := range haystack[index : index+len(needle)] {
		if code != needle[count] {
			found = false
			break
		}
		count++
	}

	if found {
		return index
	}

	return find(haystack, needle, index+1)
}

// strStr is 28. Find the Index of the First Occurrence in a String
// SEE: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	return find([]rune(haystack), []rune(needle), 0)
}
