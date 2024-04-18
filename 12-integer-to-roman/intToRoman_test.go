package integertoroman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	type Case struct {
		Name     string
		Input    int
		Expected string
	}

	testCase := []Case{
		{
			Name:     "Test 1",
			Input:    3,
			Expected: "III",
		},
		{
			Name:     "Test 2",
			Input:    58,
			Expected: "LVIII",
		},
		{
			Name:     "Test 3",
			Input:    1994,
			Expected: "MCMXCIV",
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			result := intToRoman(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
