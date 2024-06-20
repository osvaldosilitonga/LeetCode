package validparentheses

func isValid(s string) bool {
	pairs := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := []rune{}

	for _, v := range s {
		if _, ok := pairs[v]; ok {
			stack = append(stack, v)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
