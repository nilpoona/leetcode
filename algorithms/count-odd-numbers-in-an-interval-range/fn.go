package count_odd_numbers_in_an_interval_range

func countOdds(low int, high int) int {
	count := 0
	lowIsOddNumber := false
	if low == high && low%2 != 0 {
		return 1
	}

	if low%2 != 0 {
		count++
		lowIsOddNumber = true
	}
	if high%2 != 0 {
		count++
	}

	diff := high - low - 1
	n := diff / 2
	if diff > 1 && diff%2 != 0 && !lowIsOddNumber {
		n++
	}
	if diff == 1 && !lowIsOddNumber {
		n++
	}

	return count + n
}
