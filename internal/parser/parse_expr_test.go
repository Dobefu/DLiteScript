package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "number",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: "1",
		},
		{
			name: "identifier",
			input: []*token.Token{
				token.NewToken("PI", token.TokenTypeIdentifier, 0, 0),
			},
			expected: "PI",
		},
		{
			name: "power",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("**", token.TokenTypeOperationPow, 0, 0),
				token.NewToken("2", token.TokenTypeNumber, 0, 0),
			},
			expected: "(1 ** 2)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input[1:])
			expr, err := p.parseExpr(test.input[0], nil, 0, 0)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if expr.Expr() != test.expected {
				t.Fatalf(
					"expected expr to be \"%s\", got \"%s\"",
					test.expected,
					expr.Expr(),
				)
			}
		})
	}
}

func TestParseExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []*token.Token
		recursionDepth int
		expected       string
	}{
		{
			name: "maximum recursion depth exceeded",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			recursionDepth: 1_000_000,
			expected: fmt.Sprintf(
				"maximum recursion depth of (%d) exceeded",
				maxRecursionDepth,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input[1:])
			_, err := p.parseExpr(test.input[0], nil, 0, test.recursionDepth)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Errorf(
					"expected error to be \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}

func TestHandleBasicOperatorTokensErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "unexpected EOF",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected EOF after basic operator",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("+", token.TokenTypeOperationAdd, 0, 1),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token after basic operator",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 1),
				token.NewToken("+", token.TokenTypeOperationAdd, 1, 3),
				token.NewToken("bogus", token.TokenTypeIdentifier, 3, 8),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "+"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input[1:])
			_, err := p.handleBasicOperatorTokens(test.input[0], nil, 0, 0)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected error to be \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}

func TestHandlePowTokenErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "unexpected EOF",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected EOF after power",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("**", token.TokenTypeOperationPow, 0, 1),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token after power",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 1),
				token.NewToken("**", token.TokenTypeOperationPow, 1, 3),
				token.NewToken("bogus", token.TokenTypeIdentifier, 3, 8),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "**"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input[1:])
			_, err := p.handlePowToken(nil, 0, 0)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected error to be \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
func TestHandleArrayToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         []*token.Token
		leftExpr      ast.ExprNode
		minPrecedence int
		expected      string
	}{
		{
			name: "array with lower precedence",
			input: []*token.Token{
				token.NewToken("[", token.TokenTypeLBracket, 0, 0),
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("]", token.TokenTypeRBracket, 0, 0),
			},
			leftExpr: &ast.Identifier{
				Value:    "x",
				StartPos: 0,
				EndPos:   0,
			},
			minPrecedence: bindingPowerAssignment,
			expected:      "x",
		},
		{
			name: "array with higher precedence",
			input: []*token.Token{
				token.NewToken("[", token.TokenTypeLBracket, 0, 0),
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("]", token.TokenTypeRBracket, 0, 0),
			},
			leftExpr: &ast.Identifier{
				Value:    "x",
				StartPos: 0,
				EndPos:   0,
			},
			minPrecedence: bindingPowerDefault,
			expected:      "x[1]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			expr, err := p.handleArrayToken(
				test.input[1],
				test.leftExpr,
				test.minPrecedence,
				0,
			)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if expr.Expr() != test.expected {
				t.Fatalf(
					"expected expr to be \"%s\", got \"%s\"",
					test.expected,
					expr.Expr(),
				)
			}
		})
	}
}

func TestHandleArrayTokenErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{

		{
			name: "unexpected EOF",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected EOF after array",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("[", token.TokenTypeLBracket, 0, 1),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "error parsing expression",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("[", token.TokenTypeLBracket, 0, 1),
				token.NewToken("]", token.TokenTypeRBracket, 1, 2),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "]"),
			),
		},
		{
			name: "unexpected EOF after array",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("[", token.TokenTypeLBracket, 0, 1),
				token.NewToken("bogus", token.TokenTypeIdentifier, 1, 6),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token after array",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
				token.NewToken("[", token.TokenTypeLBracket, 0, 1),
				token.NewToken("0", token.TokenTypeNumber, 1, 2),
				token.NewToken("bogus", token.TokenTypeIdentifier, 1, 6),
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgExpectedCloseBracket, "bogus"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input[1:])
			_, err := p.handleArrayToken(test.input[0], nil, 0, 0)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected error to be \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
