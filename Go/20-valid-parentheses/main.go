package validparentheses

var openParentheses = map[rune]string{
	'{': "cb",
	'(': "rb",
	'[': "sb",
}

var closeParentheses = map[rune]string{
	'}': "cb",
	')': "rb",
	']': "sb",
}

func isValid(s string) bool {
	var open []string

	for _, v := range s {
		if opv, ok := openParentheses[v]; ok {
			open = append(open, opv)
		} else if val, ok := closeParentheses[v]; ok {
			if len(open) == 0 {
				return false
			}

			if val != open[len(open)-1] {
				return false
			} else {
				open = open[:len(open)-1]
			}
		}
	}

	if len(open) != 0 {
		return false
	}

	return true
}
