package removeelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	InputNums   []int
	InputVal    int
	Output      int
	Nums        []int
}

func TestRemoveElement(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test 1",
			InputNums:   []int{3, 2, 2, 3},
			InputVal:    3,
			Output:      2,
			Nums:        []int{2, 2},
		},
		{
			Description: "Test 2",
			InputNums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			InputVal:    2,
			Output:      5,
			Nums:        []int{0, 1, 4, 0, 3},
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			nums, output := removeElement(v.InputNums, v.InputVal)

			assert.Equal(t, v.Output, output, fmt.Sprintf("Expected: %v, Result: %v", v.Output, output))
			assert.ElementsMatch(t, v.Nums, nums, fmt.Sprintf("Expected: %v, Result: %v", v.Nums, nums))
		})
	}

}
