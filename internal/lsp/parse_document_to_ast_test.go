package lsp

import (
	"testing"
)

func TestParseDocumentToAstErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "invalid token",
			input:    "return",
			expected: "unexpected end of expression at position 1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := parseDocumentToAst(test.input)

			if err == nil {
				t.Errorf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
