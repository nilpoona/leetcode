package divide_two_integers

// divide is 29. Divide Two Integers
// SEE: https://leetcode.com/problems/divide-two-integers/
// https://enockey.com/shift-operation-binary-number-multiplication-and-division/
func divide(dividend, divisor int) int {
	if divisor == 0 {
		return -1
	}

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

	return sign * quotient
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
