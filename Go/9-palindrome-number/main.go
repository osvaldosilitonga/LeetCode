package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xCopy := x
	xReversed := 0
	for xCopy > 0 {
		xReversed = xReversed*10 + xCopy%10
		xCopy /= 10
	}

	return x == xReversed
}
