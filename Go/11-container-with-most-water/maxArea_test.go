package containerwithmostwater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test 1",
			Input:       input1,
			Expected:    49,
		},
		{
			Description: "Test 2",
			Input:       input2,
			Expected:    1,
		},
		{
			Description: "Test 3",
			Input:       input3,
			Expected:    705634720,
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := maxArea2(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}

}
