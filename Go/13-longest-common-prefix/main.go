package longestcommonprefix

import "strings"

func longestCommonPrefix(strs []string) string {
	var result, tmp string

	for _, c := range strs[0] {
		tmp += string(c)

		for _, val := range strs[1:] {
			if !strings.HasPrefix(val, tmp) {
				return result
			}
		}

		result = tmp
	}

	return result
}
