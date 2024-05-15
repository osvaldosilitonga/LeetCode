package validanagram

import (
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRune := []rune(s)
	tRune := []rune(t)

	sr := sortSlice(sRune)
	tr := sortSlice(tRune)

	for i, v := range sr {
		if v != tr[i] {
			return false
		}
	}

	return true
}

func sortSlice(r []rune) []rune {
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return r
}
