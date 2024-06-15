package stringtointeger

import (
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	var tmp string

	for i, v := range s {
		if i == 0 && (string(v) == "-" || string(v) == "+") {
			tmp += string(v)
			continue
		}

		if _, err := strconv.Atoi(string(v)); err != nil {
			break
		}

		tmp += string(v)
	}

	result, _ := strconv.ParseInt(tmp, 10, 32)

	return int(result)
}
