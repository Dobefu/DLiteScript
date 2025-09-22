package formatter

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func TestFormatNumberLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     *ast.NumberLiteral
		formatter *Formatter
		depth     int
		expected  string
	}{
		{
			name:      "number literal 1",
			input:     &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "1\n",
		},
		{
			name:      "number literal 1.1",
			input:     &ast.NumberLiteral{Value: "1.1", StartPos: 0, EndPos: 3},
			formatter: &Formatter{indentSize: 2, indentChar: " "},
			depth:     0,
			expected:  "1.1\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			formatter := New()
			result := formatter.Format(test.input)

			if result != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}
