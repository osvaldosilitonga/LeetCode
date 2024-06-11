package reverseinteger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	Input       int
	Expected    int
}

func TestReverse(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test 1",
			Input:       123,
			Expected:    321,
		},
		{
			Description: "Test 2",
			Input:       -123,
			Expected:    -321,
		},
		{
			Description: "Test 3",
			Input:       120,
			Expected:    21,
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := reverse(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
