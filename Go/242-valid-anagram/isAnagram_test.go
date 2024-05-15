package validanagram

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAnagram(t *testing.T) {
	type Case struct {
		Name     string
		Input1   string
		Input2   string
		Expected bool
	}

	testCase := []Case{
		{
			Name:     "Test 1",
			Input1:   "anagram",
			Input2:   "nagaram",
			Expected: true,
		},
		{
			Name:     "Test 2",
			Input1:   "rat",
			Input2:   "car",
			Expected: false,
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			result := isAnagram(v.Input1, v.Input2)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
