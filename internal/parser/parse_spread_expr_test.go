package parser

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseSpreadExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input: []*token.Token{
				token.NewToken("...", token.TokenTypeOperationSpread, 0, 0),
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: "...1",
		},
	}

	for _, test := range tests {
		t.Run(test.input[0].Atom, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			expr, err := p.Parse()

			if err != nil {
				t.Errorf("Error parsing spread expression: %v", err)
			}

			if expr.Expr() != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, expr.Expr())
			}
		})
	}
}

func TestParseSpreadExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input: []*token.Token{
				token.NewToken("...", token.TokenTypeOperationSpread, 0, 0),
			},
			expected: "unexpected end of expression at position 1",
		},
	}

	for _, test := range tests {
		t.Run(test.input[0].Atom, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.Parse()

			if err == nil {
				t.Fatalf("Expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("Expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
