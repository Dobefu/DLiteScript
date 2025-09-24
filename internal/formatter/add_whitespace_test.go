package formatter

import (
	"strings"
	"testing"
)

func TestAddWhitespace(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name:      "depth 0",
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     0,
			expected:  "",
		},
		{
			name:      "depth 1",
			formatter: &Formatter{indentSize: 2, indentChar: " ", maxLineLength: 80},
			depth:     1,
			expected:  "  ",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			builder := &strings.Builder{}
			test.formatter.addWhitespace(builder, test.depth)

			if builder.String() != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, builder.String())
			}
		})
	}
}
