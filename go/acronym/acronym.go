package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	output := ""
	next := true

	// Loop over runes
	for idx, letter := range s {
		if next && unicode.IsLetter(letter) {
			output += strings.ToUpper(string(s[idx]))
			next = false

			// Skip over space, punctuation, and '
		} else if unicode.IsSpace(letter) || unicode.IsPunct(letter) && letter != 39 {
			next = true
		}
	}

	return output
}
