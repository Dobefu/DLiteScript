package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseSpreadExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "spread expression",
			input: []*token.Token{
				token.NewToken("...", token.TokenTypeOperationSpread, 0, 0),
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: "...1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			expr, err := p.Parse()

			if err != nil {
				t.Fatalf("error parsing spread expression: %v", err)
			}

			if expr.Expr() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, expr.Expr())
			}
		})
	}
}

func TestParseSpreadExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "spread operator without expression",
			input: []*token.Token{
				token.NewToken("...", token.TokenTypeOperationSpread, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 4",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "spread operator with invalid expression",
			input: []*token.Token{
				token.NewToken("...", token.TokenTypeOperationSpread, 0, 0),
				token.NewToken("+", token.TokenTypeOperationAdd, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 5",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, err := p.Parse()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
