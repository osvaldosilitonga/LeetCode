package reverseinteger

import (
	"strconv"
)

func reverse(x int) int {
	str := strconv.Itoa(x)

	isNegative := false

	if x < 0 {
		str = string(str[1:])
		isNegative = true
	}

	runeStr := []rune(str)
	// reverse
	for i, j := 0, len(runeStr)-1; i < j; i++ {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]

		j--
	}

	if isNegative {
		str = "-" + string(runeStr)

		result, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return 0
		}

		return int(result)
	}

	result, err := strconv.ParseInt(string(runeStr), 10, 32)
	if err != nil {
		return 0
	}

	return int(result)
}
