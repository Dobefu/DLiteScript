package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber),
			},
			expected: "1",
		},
		{
			input: []*token.Token{
				token.NewToken("PI", token.TokenTypeIdentifier),
			},
			expected: "PI",
		},
	}

	for _, test := range tests {
		expr, err := NewParser(test.input).parseExpr(test.input[0], nil, 0, 0)

		if err != nil {
			t.Errorf("expected error to be nil, got '%s'", err.Error())
		}

		if expr.Expr() != test.expected {
			t.Errorf("expected expr to be '%s', got '%s'", test.expected, expr.Expr())
		}
	}
}

func TestParseExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input          []*token.Token
		expected       string
		recursionDepth int
	}{
		{
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber),
			},
			recursionDepth: 1_000_000,
			expected:       fmt.Sprintf("maximum recursion depth of (%d) exceeded", maxRecursionDepth),
		},
	}

	for _, test := range tests {
		_, err := NewParser(test.input).parseExpr(test.input[0], nil, 0, test.recursionDepth)

		if err == nil {
			t.Errorf("expected error to be not nil, got nil")
		}

		if err.Error() != test.expected {
			t.Errorf("expected error to be '%s', got '%s'", test.expected, err.Error())
		}
	}
}
