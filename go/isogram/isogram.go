package isogram

import (
	"strings"
)

func IsIsogram(s string) bool {
	letterSet := make(map[string]int)

	for _, c := range s {
		if string(c) == "-" || string(c) == " " {
			continue
		}
		if _, ok := letterSet[strings.ToUpper(string(c))]; ok {
			return false
		} else {
			letterSet[strings.ToUpper(string(c))]++
		}
	}
	return true
}

// Alternative from: https://exercism.io/tracks/go/exercises/isogram/solutions/76d4a807ebdf40a685a2efa21f85bb7f
// IsIsogram returns true iff s is an isogram.
func IsIsogram2(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		if _, ok := seen[unicode.ToUpper(r)]; ok {
			return false
		}
		seen[unicode.ToUpper(r)] = true
	}
	return true
}
