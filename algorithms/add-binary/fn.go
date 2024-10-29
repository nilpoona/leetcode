package add_binary

func addBinary(a string, b string) string {
	maxLen := max(len(a), len(b))
	a = padZeroes(a, maxLen)
	b = padZeroes(b, maxLen)

	result := make([]byte, maxLen+1)
	carry := 0

	for i := maxLen - 1; i >= 0; i-- {
		sum := int(a[i]-'0') + int(b[i]-'0') + carry
		result[i+1] = byte(sum%2 + '0')
		carry = sum / 2
	}

	if carry == 1 {
		result[0] = '1'
		return string(result)
	} else {
		return string(result[1:])
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func padZeroes(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}
	return s
}
