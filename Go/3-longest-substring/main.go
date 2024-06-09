package longestsubstring

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var result int
	sub := []string{}

	for _, char := range s {
		if !strings.Contains(strings.Join(sub, ""), string(char)) {
			sub = append(sub, string(char))

			result = max(result, len(sub))
		} else {
			cut := strings.Index(strings.Join(sub, ""), string(char))

			sub = append(sub[cut+1:], string(char))
		}
	}

	return result
}
