package longestcommonprefix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	Input       []string
	Expected    string
}

func TestLongestCommonPrefix(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test Case 1",
			Input:       []string{"flower", "flow", "flight"},
			Expected:    "fl",
		},
		{
			Description: "Test Case 2",
			Input:       []string{"dog", "racecar", "car"},
			Expected:    "",
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := longestCommonPrefix(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
