package valid_parentheses

var brackets = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func isValid(s string) bool {
	if s == "" || len(s) % 2 != 0 {
		return false
	}

	expected := make([]string, 0)
	return validBracket(s, expected)
}

func validBracket(s string, expectedBrackets []string) bool{
	if s == "" && len(expectedBrackets) != 0{
		return false
	}
	if s == "" {
		return true
	}
	h := string(s[0])
	closed, ok := brackets[h]
	if ok {
		expectedBrackets = append([]string{closed}, expectedBrackets...)
		return validBracket(s[1:], expectedBrackets)
	} else {
		if len(expectedBrackets) == 0 {
			return false
		}
	}
	closed = expectedBrackets[0]
	if len(expectedBrackets) > 0 {
		expectedBrackets = expectedBrackets[1:]
	}
	if h != closed {
		return false
	}
	return validBracket(s[1:], expectedBrackets)
}