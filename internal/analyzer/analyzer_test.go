package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, New(), "a")
}

func TestStripFormatVerbs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "double percent",
			input: "%%",
			want:  "%",
		},
		{
			name:  "format for digit",
			input: "%d",
			want:  "",
		},
		{
			name:  "format for digit 2",
			input: "got %d items",
			want:  "got  items",
		},
		{
			name:  "format with width, precision",
			input: "%08.3f",
			want:  "",
		},
		{
			name:  "no verb",
			input: "no verb",
			want:  "no verb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := stripFormatVerbs(tt.input); got != tt.want {
				t.Errorf("stripFormatVerbs(%q) = %q, want %q", tt.input, got, tt.want)

			}
		})
	}
}
