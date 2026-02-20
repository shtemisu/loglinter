package rules

import (
	"strings"
	"unicode"
)

var sensitiveKeywords = []string{
	"password",
	"pwd",
	"pass",
	"token",
	"jwt",
	"api_key",
	"apikey",
	"api-key",
	"secret",
	"private",
	"credential",
	"auth",
	"authorization",
	"bearer",
	"cookie",
}

func HasSensetiveData(s string) bool {
	for _, kwords := range sensitiveKeywords {
		if strings.Contains(s, kwords) {
			return true
		}
	}
	return false
}

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

func HasSpecialChars(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) ||
			unicode.IsDigit(r) ||
			r == ' ' ||
			r == '.' ||
			r == ',' ||
			r == '-' ||
			r == ':' ||
			r == '!') {
			return true
		}
	}
	return false
}
