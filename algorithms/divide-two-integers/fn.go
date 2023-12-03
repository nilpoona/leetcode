package divide_two_integers

var (
	maximum = 1<<31 - 1
	minimum = -1 << 31
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

	result := sign * quotient

	if result > maximum {
		result = maximum
	}

	if result < minimum {
		result = minimum
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
