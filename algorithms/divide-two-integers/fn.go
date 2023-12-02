package divide_two_integers

import (
	"math"
)

var (
	maximum = math.Pow(float64(2), float64(31)) - 1
	minimum = math.Pow(float64(-2), float64(31)) - 1
)

// divide is 29. Divide Two Integers
// SEE: https://leetcode.com/problems/divide-two-integers/
// https://enockey.com/shift-operation-binary-number-multiplication-and-division/
func divide(dividend, divisor int) int {
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	quotient := 0
	for dividend >= divisor {
		shifts := 0
		for dividend >= (divisor << shifts) {
			shifts++
		}
		quotient += 1 << (shifts - 1)
		dividend -= divisor << (shifts - 1)
	}

	if quotient > int(maximum) {
		quotient = int(maximum)
	}

	if quotient < int(minimum) {
		quotient = int(minimum)
	}

	return sign * quotient
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
