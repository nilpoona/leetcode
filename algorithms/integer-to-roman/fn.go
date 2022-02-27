package integer_to_roman

import (
	"fmt"
	"strings"
)

func genRoman(n, digits int) string {
	var baseSymbol, subSymbol string
	switch digits {
	case 1:
		if n == 4 {
			return "IV"
		}
		if n == 9 {
			return "IX"
		}
		if n == 5 {
			return "V"
		}
		baseSymbol = "I"
		subSymbol = "V"
	case 10:
		if n == 4 {
			return "XL"
		}
		if n == 9 {
			return "XC"
		}
		if n == 5 {
			return "L"
		}
		baseSymbol = "X"
		subSymbol = "L"
	case 100:
		if n == 4 {
			return "CD"
		}
		if n == 9 {
			return "CM"
		}
		if n == 5 {
			return "D"
		}
		baseSymbol = "C"
		subSymbol = "D"
	case 1000:
		baseSymbol = "M"
		subSymbol = ""
	default:
		return ""
	}

	if n < 5 {
		return strings.Repeat(baseSymbol, n)
	}

	return fmt.Sprintf("%s%s", subSymbol, strings.Repeat(baseSymbol, n-5))
}

func intToRoman(num int) string {
	digits := 1
	roman := ""
	for {
		n := (num / digits) % 10
		if n == 0 && digits > num {
			break
		}
		r := genRoman(n, digits)
		roman = r + roman
		digits = digits * 10
	}
	return roman
}
