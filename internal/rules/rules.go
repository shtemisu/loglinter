package rules

import (
	"strings"
	"unicode"
)

var sensitiveKeywords = []string{
	"password", "pwd", "pass", "token", "jwt", "api_key",
	"apikey", "api-key", "secret", "private", "credential",
	"auth", "authorization", "bearer", "cookie",
}

func HasSensitiveData(s string) bool {
	s = strings.ToLower(s)
	for _, kw := range sensitiveKeywords {
		if strings.Contains(s, kw) {
			return true
		}
	}
	return false
}

func IsLower(s string) bool {
	return s == strings.ToLower(s)
}

func OnlyEnglish(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Latin, r) ||
			unicode.IsDigit(r) ||
			r == ' ' ||
			r == '.' ||
			r == ',' ||
			r == '-' ||
			r == ':' ||
			r == '\'' ||
			r == '_' {
			continue
		}
		return false
	}
	return true
}

func HasSpecialChars(s string) bool {
	repeated := []string{"...", "..", "!!!", "!!", "??", "?!"}
	for _, pattern := range repeated {
		if strings.Contains(s, pattern) {
			return true
		}
	}
	for _, r := range s {
		if unicode.IsLetter(r) ||
			unicode.IsDigit(r) ||
			r == ' ' ||
			r == '.' ||
			r == ',' ||
			r == '-' ||
			r == '\'' ||
			r == '_' {
			continue
		}
		return true

	}
	return false
}
