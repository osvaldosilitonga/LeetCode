package zigzagconversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	result := make([][]string, numRows)

	char := 0
	pointer := 0
	col := 0
	operator := true

	for {
		result[pointer] = append(result[pointer], string(s[char]))
		char++

		if char == len(s) {
			break
		}

		col++
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

	var str string

	for _, v := range result {
		str += strings.Join(v, "")
	}

	return str
}
