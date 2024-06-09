package longestpalindromicsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	Input       string
	Expected    string
}

func TestLongestPalindrome(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test Case : babad",
			Input:       "babad",
			Expected:    "bab",
		},
		{
			Description: "Test Case : cbbd",
			Input:       "cbbd",
			Expected:    "bb",
		},
		{
			Description: "Test Case : a",
			Input:       "a",
			Expected:    "a",
		},
		{
			Description: "Test Case : ac",
			Input:       "ac",
			Expected:    "a",
		},
		{
			Description: "Test Case : bb",
			Input:       "bb",
			Expected:    "bb",
		},
		{
			Description: "Test Case : ccc",
			Input:       "ccc",
			Expected:    "ccc",
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := longestPalindrome(v.Input)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
