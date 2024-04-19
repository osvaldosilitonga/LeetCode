package twoSum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Description string
	Nums        []int
	Target      int
	Expected    any
}

func TestTwoSum(t *testing.T) {
	testCase := []Case{
		{
			Description: "Test case 1",
			Nums:        []int{2, 7, 11, 15},
			Target:      9,
			Expected:    []int{0, 1},
		},
		{
			Description: "Test case 2",
			Nums:        []int{3, 2, 4},
			Target:      6,
			Expected:    []int{1, 2},
		},
		{
			Description: "Test case 3",
			Nums:        []int{3, 3},
			Target:      6,
			Expected:    []int{0, 1},
		},
		{
			Description: "Test case 4",
			Nums:        []int{3, 2, 3},
			Target:      6,
			Expected:    []int{0, 2},
		},
	}

	for _, v := range testCase {
		result := TwoSum(v.Nums, v.Target)

		assert.ElementsMatch(t, v.Expected, result, fmt.Sprintf("Description: %v. Expected: %v, Actual: %v", v.Description, v.Expected, result))
	}
}
