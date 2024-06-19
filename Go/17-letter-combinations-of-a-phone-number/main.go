package lettercombinationsofaphonenumber

var phoneNumber = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	var result []string

	recursive(&result, digits, "", 0)

	return result
}

func recursive(result *[]string, digits string, current string, index int) {
	if len(digits) == index {
		*result = append(*result, current)
		return
	}

	letters := phoneNumber[rune(digits[index])]
	for _, letter := range letters {
		recursive(result, digits, current+string(letter), index+1)
	}
}
