package threesum

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type Data struct {
	Description string
	Input       []int
	Expected    [][]int
}

func TestThreeSum(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test Case 1",
			Input:       []int{-1, 0, 1, 2, -1, -4},
			Expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			Description: "Test Case 2",
			Input:       []int{0, 1, 1},
			Expected:    [][]int{},
		},
		{
			Description: "Test Case 2",
			Input:       []int{0, 0, 0},
			Expected: [][]int{
				{0, 0, 0},
			},
		},
	}

	// Test for `threeSum` function
	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := threeSum(v.Input)

			resultCheck(t, v.Description, result, v.Expected)
		})
	}

	// Test for `threeSum2` function
	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := threeSum2(v.Input)

			resultCheck(t, v.Description, result, v.Expected)
		})
	}
}

func resultCheck(t *testing.T, description string, result, expected [][]int) {
	if len(result) != len(expected) {
		t.Error(fmt.Sprintf("FAIL - %v", description))
	}

	for _, rv := range result {
		isContains := false
		for _, ev := range expected {
			sort.Ints(rv)
			sort.Ints(ev)

			if reflect.DeepEqual(rv, ev) {
				isContains = true
			}
		}

		if isContains == false {
			t.Error(fmt.Sprintf("FAIL - %v", description))
		}
	}
}
