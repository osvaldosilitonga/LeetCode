package longestsubstring

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var result int
	sub := []string{}

	for _, char := range s {
		if !strings.Contains(strings.Join(sub, ""), string(char)) {
			sub = append(sub, string(char))

			fmt.Printf("Char: %v, Sub: %v \n", string(char), sub)

			result = max(result, len(sub))
		} else {
			cut := strings.Index(strings.Join(sub, ""), string(char))

			fmt.Printf("Char: %v, Sub: %v, Cut: %v \n", string(char), sub, cut)

			sub = append(sub[cut+1:], string(char))
		}
	}

	return result
}
