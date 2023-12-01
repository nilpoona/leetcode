package check_if_two_string_arrays_are_equivalent

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
