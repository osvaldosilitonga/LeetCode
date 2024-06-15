package stringtointeger

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

func TestStringToInteger(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test 1",
			Input:       "42",
			Expected:    42,
		},
		{
			Description: "Test 2",
			Input:       " -042",
			Expected:    -42,
		},
		{
			Description: "Test 3",
			Input:       "1337c0d3",
			Expected:    1337,
		},
		{
			Description: "Test 4",
			Input:       "0-1",
			Expected:    0,
		},
		{
			Description: "Test 5",
			Input:       "words and 987",
			Expected:    0,
		},
		{
			Description: "Test 6",
			Input:       "-91283472332",
			Expected:    -2147483648,
		},
		{
			Description: "Test 7",
			Input:       "+1",
			Expected:    1,
		},
		{
			Description: "Test 8",
			Input:       "-13+8",
			Expected:    -13,
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := myAtoi(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
