package lettercombinationsofaphonenumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	Input       string
	Expected    []string
}

func TestLetterCombinations(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test Case 1",
			Input:       "23",
			Expected:    []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			Description: "Test Case 2",
			Input:       "",
			Expected:    []string{},
		},
		{
			Description: "Test Case 3",
			Input:       "2",
			Expected:    []string{"a", "b", "c"},
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := letterCombinations(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
