package palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	reverse := func(s string) string {
		reversed := make([]rune, len(s))
		for i, r := range s {
			reversed[len(s) - i -1] = r
		}
		return string(reversed)
	}
	return str == reverse(str)
}
