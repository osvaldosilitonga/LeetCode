package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	// iterasi selama kedua list tidak kosong
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	// jika salah satu list masih memiliki elemen, tambahkan semuanya ke hasil
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	// return node berikutnya dari dummy, yang merupakan head dari
	return dummy.Next
}
