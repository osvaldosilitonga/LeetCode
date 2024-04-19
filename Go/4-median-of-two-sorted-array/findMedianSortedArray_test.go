package medianoftwosortedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArray(t *testing.T) {
	type Data struct {
		Name     string
		Nums1    []int
		Nums2    []int
		Expected float64
	}

	testCase := []Data{
		{
			Name:     "Test 1",
			Nums1:    []int{1, 3},
			Nums2:    []int{2},
			Expected: 2,
		},
		{
			Name:     "Test 2",
			Nums1:    []int{1, 2},
			Nums2:    []int{3, 4},
			Expected: 2.5,
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			result := findMedianSortedArray(v.Nums1, v.Nums2)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
