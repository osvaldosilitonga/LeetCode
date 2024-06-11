package zigzagconversion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data struct {
	Description string
	InputString string
	NumRows     int
	Expected    string
}

func TestZigzagConversion(t *testing.T) {
	testCase := []Data{
		{
			Description: "Test 1",
			InputString: "PAYPALISHIRING",
			NumRows:     3,
			Expected:    "PAHNAPLSIIGYIR",
		},
		{
			Description: "Test 2",
			InputString: "PAYPALISHIRING",
			NumRows:     4,
			Expected:    "PINALSIGYAHRPI",
		},
		{
			Description: "Test 3",
			InputString: "A",
			NumRows:     1,
			Expected:    "A",
		},
		{
			Description: "Test 4",
			InputString: "AB",
			NumRows:     1,
			Expected:    "AB",
		},
	}

	for _, v := range testCase {
		t.Run(v.Description, func(t *testing.T) {
			result := convert(v.InputString, v.NumRows)

			assert.Equal(t, v.Expected, result, fmt.Sprintf("Expected : %v, Actual : %v", v.Expected, result))
		})
	}
}
