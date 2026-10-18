package rules

import (
	"slices"
	"strings"
	"unicode"
)

// sensitiveKeywords — ключевые слова, уже разбитые на слова через splitWords.
// Составные ключи (api_key) хранятся как последовательность слов.
var sensitiveKeywords = func() [][]string {
	keywords := []string{
		"password", "passwd", "pwd", "pass", "token", "jwt",
		"apikey", "api_key", "secret", "private", "credential",
		"credentials", "auth", "authorization", "bearer", "cookie",
	}
	result := make([][]string, 0, len(keywords))
	for _, kw := range keywords {
		result = append(result, splitWords(kw))
	}
	return result
}()

// properNames — имена собственные, с которых может начинаться сообщение.
// Аббревиатуры (HTTP, S3) сюда добавлять не нужно, их распознаёт isAcronym.
var properNames = []string{
	"Redis", "PostgreSQL",
}

// allowedSpecChars — символы, кроме букв и цифр, разрешённые в сообщении.
var allowedSpecChars = []rune{' ', ':', '\'', '.', ',', '[', ']', '%'}

// HasSensitiveData reports whether s contains a sensitive keyword as a whole word.
// Words are runs of letters and digits; comparison is case-insensitive.
func HasSensitiveData(s string) bool {
	words := splitWords(s)
	for _, kw := range sensitiveKeywords {
		if containsSequence(words, kw) {
			return true
		}
	}
	return false
}

func splitWords(s string) []string {
	return strings.FieldsFunc(strings.ToLower(s), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

func containsSequence(words, sequence []string) bool {
	for i := 0; i <= len(words)-len(sequence); i++ {
		if slices.Equal(words[i:i+len(sequence)], sequence) {
			return true
		}
	}
	return false
}

// IsLower reports whether s does not start with an uppercase letter.
// Allowed exceptions for the first word: an acronym (HTTP, S3) or a name
// from properNames. An empty message is not checked.
func IsLower(s string) bool {
	if s == "" {
		return true
	}
	runes := []rune(s)
	if !unicode.IsUpper(runes[0]) {
		return true
	}
	word := firstWord(runes)
	return isAcronym(word) || slices.Contains(properNames, string(word))
}

func firstWord(runes []rune) []rune {
	for i, r := range runes {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return runes[:i]
		}
	}
	return runes
}

// isAcronym reports whether word consists of uppercase Latin letters and digits
// and is longer than one character.
func isAcronym(word []rune) bool {
	if len(word) < 2 {
		return false
	}
	for _, r := range word {
		if !(r >= 'A' && r <= 'Z') && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// OnlyEnglish reports whether s contains no non-Latin letters.
// Digits, punctuation and emoji are not checked here (see HasSpecialChars).
func OnlyEnglish(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return false
		}
	}
	return true
}

// HasSpecialChars reports whether s contains a character outside the allowed set
// (letters, digits and allowedSpecChars) or two or more dots in a row.
func HasSpecialChars(s string) bool {
	if strings.Contains(s, "..") {
		return true
	}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || slices.Contains(allowedSpecChars, r) {
			continue
		}
		return true
	}
	return false
}
