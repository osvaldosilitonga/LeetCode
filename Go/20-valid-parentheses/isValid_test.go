package validparentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	Input       string
	Expected    bool
}

func TestIsValid(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test 1",
			Input:       "()",
			Expected:    true,
		},
		{
			Description: "Test 2",
			Input:       "()[]{}",
			Expected:    true,
		},
		{
			Description: "Test 3",
			Input:       "(]",
			Expected:    false,
		},
		{
			Description: "Test 4",
			Input:       "([)]",
			Expected:    false,
		},
		{
			Description: "Test 5",
			Input:       "]",
			Expected:    false,
		},
		{
			Description: "Test 5",
			Input:       "[",
			Expected:    false,
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := isValid(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
