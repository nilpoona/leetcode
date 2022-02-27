package roman_to_integer

/**
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
Iは、V（5）とX（10）の前に置かれ、4と9を作ることができる。
XはL（50）とC（100）の前に置いて、40と90にすることができる。
CはD(500)とM(1000)の前に置くと、400と900になる。
*/
func romanToInt(s string) int {
	result := 0
	prev := ""
	for _, n := range s {
		switch string(n) {
		case "M":
			if prev == "C" {
				result += 800
				break
			}
			result += 1000
		case "D":
			if prev == "C" {
				result += 300
				break
			}
			result += 500
		case "C":
			if prev == "X" {
				result += 80
				break
			}
			result += 100
		case "L":
			if prev == "X" {
				result += 30
				break
			}
			result += 50
		case "X":
			if prev == "I" {
				result += 8
				break
			}
			result += 10
		case "V":
			if prev == "I" {
				result += 3
				break
			}
			result += 5
		case "I":
			result += 1
		default:
			result += 0
		}
		prev = string(n)
	}

	return result
}
