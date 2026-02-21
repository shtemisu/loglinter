package rules

import (
	"testing"
)

func TestIsLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "все строчные",
			input:    "starting server",
			expected: true,
		},
		{
			name:     "с заглавной буквы",
			input:    "Starting server",
			expected: false,
		},
		{
			name:     "все заглавные",
			input:    "STARTING SERVER",
			expected: false,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: true,
		},
		{
			name:     "с цифрами",
			input:    "server 123",
			expected: true,
		},
		{
			name:     "с цифрами в начале",
			input:    "123 server",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLower(tt.input); got != tt.expected {
				t.Errorf("IsLower(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestOnlyEnglish(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "только английский",
			input:    "starting server",
			expected: true,
		},
		{
			name:     "с заглавными буквами",
			input:    "Starting Server",
			expected: true,
		},
		{
			name:     "с цифрами",
			input:    "server 123",
			expected: true,
		},
		{
			name:     "с пунктуацией",
			input:    "server, please start!",
			expected: false,
		},
		{
			name:     "с русскими буквами",
			input:    "starting сервер",
			expected: false,
		},
		{
			name:     "только русские",
			input:    "запуск сервера",
			expected: false,
		},
		{
			name:     "смешанный",
			input:    "server запуск",
			expected: false,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: true,
		},
		{
			name:     "только пробелы",
			input:    "   ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := OnlyEnglish(tt.input); got != tt.expected {
				t.Errorf("OnlyEnglish(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHasSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "чистый текст",
			input:    "starting server",
			expected: false,
		},
		{
			name:     "с точкой в конце",
			input:    "starting server.",
			expected: false,
		},
		{
			name:     "с запятой",
			input:    "server, please start",
			expected: false,
		},
		{
			name:     "с дефисом",
			input:    "server-start",
			expected: false,
		},
		{
			name:     "с двоеточием",
			input:    "server: starting",
			expected: true,
		},
		{
			name:     "с апострофом",
			input:    "don't start",
			expected: false,
		},
		{
			name:     "с восклицательным знаком",
			input:    "starting server!",
			expected: true,
		},
		{
			name:     "с двумя восклицательными",
			input:    "starting server!!",
			expected: true,
		},
		{
			name:     "с вопросительным знаком",
			input:    "starting server?",
			expected: true,
		},
		{
			name:     "с многоточием",
			input:    "starting server...",
			expected: true,
		},
		{
			name:     "со звёздочкой",
			input:    "starting *server*",
			expected: true,
		},
		{
			name:     "с символом @",
			input:    "user@server",
			expected: true,
		},
		{
			name:     "с эмодзи",
			input:    "starting server 😀",
			expected: true,
		},
		{
			name:     "с цифрами",
			input:    "server 123",
			expected: false,
		},
		{
			name:     "с русскими буквами",
			input:    "starting сервер",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := HasSpecialChars(tt.input); got != tt.expected {
				t.Errorf("HasSpecialChars(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestHasSensitiveData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "чистый текст",
			input:    "user logged in",
			expected: false,
		},
		{
			name:     "password в тексте",
			input:    "user password changed",
			expected: true,
		},
		{
			name:     "pwd в тексте",
			input:    "invalid pwd",
			expected: true,
		},
		{
			name:     "token в тексте",
			input:    "token expired",
			expected: true,
		},
		{
			name:     "jwt в тексте",
			input:    "jwt validation failed",
			expected: true,
		},
		{
			name:     "api_key в тексте",
			input:    "api_key=12345",
			expected: true,
		},
		{
			name:     "secret в тексте",
			input:    "secret key",
			expected: true,
		},
		{
			name:     "auth внутри слова",
			input:    "authorization failed",
			expected: true,
		},
		{
			name:     "auth отдельно",
			input:    "auth failed",
			expected: true,
		},
		{
			name:     "смешанный регистр",
			input:    "User Password",
			expected: true,
		},
		{
			name:     "несколько слов",
			input:    "invalid authorization token",
			expected: true,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := HasSensitiveData(tt.input); got != tt.expected {
				t.Errorf("HasSensitiveData(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
