package majorityelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	type Case struct {
		Name     string
		Input    []int
		Expected int
	}

	testCase := []Case{
		{
			Name:     "Test 1",
			Input:    []int{3, 2, 3},
			Expected: 3,
		},
		{
			Name:     "Test 2",
			Input:    []int{2, 2, 1, 1, 1, 2, 2},
			Expected: 2,
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			result := majorityElement(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
