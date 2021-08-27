package bob

import (
	"strings"
	"unicode"
)

func hasLetter(remark string) bool {
	for _, r := range remark {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if strings.HasSuffix(remark, "?") {
		if strings.ToUpper(remark) == remark && hasLetter(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else {
		if strings.ToUpper(remark) == remark && hasLetter(remark) {
			return "Whoa, chill out!"
		}
		if len(remark) == 0 {
			return "Fine. Be that way!"
		}
		return "Whatever."
	}
}
