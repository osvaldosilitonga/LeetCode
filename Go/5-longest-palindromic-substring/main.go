package longestpalindromicsubstring

func longestPalindrome(s string) string {
	var result string
	var tmp string

	for start := 0; start < len(s); start++ {
		for end := start; end < len(s); end++ {
			tmp = s[start : end+1]
			if IsPalindrome(tmp) {
				if len(tmp) > len(result) {
					result = tmp
				}
			}
		}
	}

	return result
}

func IsPalindrome(str string) bool {
	for start, end := 0, len(str)-1; start < end; start++ {
		if str[start] != str[end] {
			return false
		}

		end--
	}

	return true
}
