package rules

import (
	"testing"
)

// Тесты повторяют примеры из RULES.md.

func TestIsLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "все строчные",
			input:    "server started",
			expected: true,
		},
		{
			name:     "с цифрами в начале",
			input:    "8080 port busy",
			expected: true,
		},
		{
			name:     "аббревиатура",
			input:    "HTTP request failed",
			expected: true,
		},
		{
			name:     "аббревиатура с цифрой",
			input:    "S3 upload failed",
			expected: true,
		},
		{
			name:     "аббревиатура без продолжения",
			input:    "HTTP",
			expected: true,
		},
		{
			name:     "аббревиатура с двоеточием",
			input:    "HTTP: request failed",
			expected: true,
		},
		{
			name:     "имя собственное из списка",
			input:    "Redis connection lost",
			expected: true,
		},
		{
			name:     "имя собственное с двоеточием",
			input:    "PostgreSQL: query failed",
			expected: true,
		},
		{
			name:     "пустая строка не проверяется",
			input:    "",
			expected: true,
		},
		{
			name:     "с заглавной буквы",
			input:    "Server started",
			expected: false,
		},
		{
			name:     "одно слово с заглавной",
			input:    "Error",
			expected: false,
		},
		{
			name:     "одна заглавная буква не аббревиатура",
			input:    "A request failed",
			expected: false,
		},
		{
			name:     "слово не целиком заглавное",
			input:    "HTTPServer started",
			expected: false,
		},
		{
			name:     "кириллица с заглавной",
			input:    "Ошибка",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
			input:    "server started",
			expected: true,
		},
		{
			name:     "с цифрами",
			input:    "port 8080",
			expected: true,
		},
		{
			name:     "пунктуация и эмодзи не проверяются",
			input:    "server started!🚀",
			expected: true,
		},
		{
			name:     "спецсимволы не проверяются",
			input:    "user_id: {value} 100%",
			expected: true,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: true,
		},
		{
			name:     "только русские",
			input:    "ошибка подключения",
			expected: false,
		},
		{
			name:     "смешанный",
			input:    "server запущен",
			expected: false,
		},
		{
			name:     "другая письменность",
			input:    "server 起動",
			expected: false,
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
			input:    "server started",
			expected: false,
		},
		{
			name:     "с двоеточием",
			input:    "failed to connect: timeout",
			expected: false,
		},
		{
			name:     "с апострофом",
			input:    "can't connect",
			expected: false,
		},
		{
			name:     "с одной точкой",
			input:    "failed to read config.yaml",
			expected: false,
		},
		{
			name:     "с запятой",
			input:    "retry 3, giving up",
			expected: false,
		},
		{
			name:     "с квадратными скобками",
			input:    "[worker] started",
			expected: false,
		},
		{
			name:     "с цифрами",
			input:    "port 8080",
			expected: false,
		},
		{
			name:     "кириллица не проверяется",
			input:    "сервер запущен",
			expected: false,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: false,
		},
		{
			name:     "две точки",
			input:    "loading..",
			expected: true,
		},
		{
			name:     "многоточие",
			input:    "something went wrong...",
			expected: true,
		},
		{
			name:     "восклицательный знак",
			input:    "connection failed!",
			expected: true,
		},
		{
			name:     "вопросительный знак",
			input:    "connection failed?",
			expected: true,
		},
		{
			name:     "подчёркивание",
			input:    "user_id not found",
			expected: true,
		},
		{
			name:     "дефис",
			input:    "non-empty value",
			expected: true,
		},
		{
			name:     "фигурные скобки",
			input:    "value {id}",
			expected: true,
		},
		{
			name:     "слэш",
			input:    "path a/b",
			expected: true,
		},
		{
			name:     "эмодзи",
			input:    "server started🚀",
			expected: true,
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
			name:     "pass внутри слова",
			input:    "operation passed",
			expected: false,
		},
		{
			name:     "auth внутри слова",
			input:    "author not found",
			expected: false,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: false,
		},
		{
			name:     "auth отдельно",
			input:    "auth failed",
			expected: true,
		},
		{
			name:     "слово в конце",
			input:    "invalid pwd",
			expected: true,
		},
		{
			name:     "единственное слово",
			input:    "password",
			expected: true,
		},
		{
			name:     "через подчёркивание",
			input:    "user_password",
			expected: true,
		},
		{
			name:     "смешанный регистр",
			input:    "User Password",
			expected: true,
		},
		{
			name:     "api_key через подчёркивание",
			input:    "api_key=12345",
			expected: true,
		},
		{
			name:     "api-key через дефис",
			input:    "invalid api-key",
			expected: true,
		},
		{
			name:     "apikey слитно",
			input:    "apikey expired",
			expected: true,
		},
		{
			name:     "token с двоеточием",
			input:    "token: ",
			expected: true,
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
