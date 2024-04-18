package addTwoNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Description string
	List1       []ListNode
	List2       []ListNode
	Expected    []ListNode
}

func TestMain(t *testing.T) {
	// Helper function to create a linked list from an array of integers
	createList := func(arr []int) *ListNode {
		dummyHead := &ListNode{}
		current := dummyHead
		for _, val := range arr {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		return dummyHead.Next
	}

	// Helper function to convert a linked list to an array
	listToArray := func(head *ListNode) []int {
		var arr []int
		for head != nil {
			arr = append(arr, head.Val)
			head = head.Next
		}
		return arr
	}

	tests := []struct {
		name   string
		l1     []int
		l2     []int
		output []int
	}{
		{"Test 1", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"Test 2", []int{0}, []int{0}, []int{0}},
		{"Test 3", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l1 := createList(test.l1)
			l2 := createList(test.l2)
			expected := test.output

			result := addTwoNumbers(l1, l2)
			actual := listToArray(result)

			assert.Equal(t, expected, actual)
		})
	}
}
