package zigzagconversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	result := make([]string, numRows)

	char, pointer := 0, 0
	operator := true // the pointer increases if true, decreases if false

	for col := 0; char != len(s); col++ {
		result[pointer] += string(s[char])
		char++

		if operator {
			pointer++
		} else {
			pointer--
		}

		if pointer == 0 {
			operator = true
		}
		if pointer == numRows-1 {
			operator = false
		}
	}

	return strings.Join(result, "")
}
