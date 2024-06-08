package longestsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	Input       string
	Expected    int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCase := []Data{
		{
			Description: "Testing 1",
			Input:       "abcabcbb",
			Expected:    3,
		},
		{
			Description: "Testing 2",
			Input:       "bbbbb",
			Expected:    1,
		},
		{
			Description: "Testing 3",
			Input:       "pwwkew",
			Expected:    3,
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := lengthOfLongestSubstring(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
