package string_to_integer_atoi

import (
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimLeft(s, "　")
	if s == "" {
		return 0
	}

	base := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if len(s) < 1 {
			return 0
		}
	}

	n := 0
	th := 2147483648
	for _, c := range s {
		c -= '0'
		if c < 0 || c > 9 {
			break
		}
		n = n*10 + int(c)
		if n >= th {
			break
		}
	}
	if base[0] == '-' {
		n = -n
	}

	if n >= th {
		n = th - 1
	}

	if n < -th {
		n = -th
	}

	return n
}
