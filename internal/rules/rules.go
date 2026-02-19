package rules

import (
	"strings"
	"unicode"
)

func IsLower(s string) bool {
	return s == strings.ToLower(s)
}

func OnlyEnglishAndWithoutSpecChar(s string) bool {
	for _, r := range s {
		if !(unicode.Is(unicode.Latin, r) ||
			unicode.IsPunct(r) ||
			unicode.IsSpace(r) ||
			unicode.IsDigit(r)) {
			return false
		}
	}
	return true
}
