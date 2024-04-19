package palindromenumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	type Case struct {
		Description string
		Input       int
		Expected    bool
	}

	testCase := []Case{
		{
			Description: "Test 1 - Success",
			Input:       121,
			Expected:    true,
		},
		{
			Description: "Test 2 - Fail",
			Input:       -121,
			Expected:    false,
		},
		{
			Description: "Test 3 - Fail",
			Input:       10,
			Expected:    false,
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := isPalindrome(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
