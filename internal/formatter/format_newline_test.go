package formatter

import (
	"strings"
	"testing"
)

func TestFormatNewline(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		formatter *Formatter
		expected  string
	}{
		{
			name:      "newline literal",
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			expected:  "\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.formatNewline(builder)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
