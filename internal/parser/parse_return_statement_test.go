package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseReturnStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         []*token.Token
		expectedValue ast.ExprNode
	}{
		{
			name: "no values",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      "\n",
					TokenType: token.TokenTypeNewline,
					StartPos:  6,
					EndPos:    7,
				},
			},
			expectedValue: &ast.ReturnStatement{
				Values:    []ast.ExprNode{},
				NumValues: 0,
				StartPos:  0,
				EndPos:    6,
			},
		},
		{
			name: "single value",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      "1",
					TokenType: token.TokenTypeNumber,
					StartPos:  6,
					EndPos:    7,
				},
				{
					Atom:      "\n",
					TokenType: token.TokenTypeNewline,
					StartPos:  7,
					EndPos:    8,
				},
			},
			expectedValue: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{
						Value:    "1",
						StartPos: 6,
						EndPos:   7,
					},
				},
				NumValues: 1,
				StartPos:  0,
				EndPos:    7,
			},
		},
		{
			name: "multiple values",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      "1",
					TokenType: token.TokenTypeNumber,
					StartPos:  6,
					EndPos:    7,
				},
				{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
					StartPos:  7,
					EndPos:    8,
				},
				{
					Atom:      "2",
					TokenType: token.TokenTypeNumber,
					StartPos:  8,
					EndPos:    9,
				},
				{
					Atom:      "\n",
					TokenType: token.TokenTypeNewline,
					StartPos:  9,
					EndPos:    10,
				},
			},
			expectedValue: &ast.ReturnStatement{
				Values: []ast.ExprNode{
					&ast.NumberLiteral{
						Value:    "1",
						StartPos: 6,
						EndPos:   7,
					},
					&ast.NumberLiteral{
						Value:    "2",
						StartPos: 8,
						EndPos:   9,
					},
				},
				NumValues: 2,
				StartPos:  0,
				EndPos:    10,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			result, err := parser.Parse()

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if result.Expr() != test.expectedValue.Expr() {
				t.Fatalf(
					"expected value to be \"%s\", got \"%s\"",
					test.expectedValue.Expr(),
					result.Expr(),
				)
			}
		})
	}
}

func TestParseReturnStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "no values",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid return value",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  6,
					EndPos:    12,
				},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "return"),
			),
		},
		{
			name: "comma only",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
					StartPos:  6,
					EndPos:    7,
				},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, ","),
			),
		},
		{
			name: "trailing comma",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      "1",
					TokenType: token.TokenTypeNumber,
					StartPos:  6,
					EndPos:    7,
				},
				{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
					StartPos:  7,
					EndPos:    8,
				},
				{
					Atom:      "\n",
					TokenType: token.TokenTypeReturn,
					StartPos:  8,
					EndPos:    9,
				},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 4",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "\n"),
			),
		},
		{
			name: "multiple values without comma",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      "1",
					TokenType: token.TokenTypeNumber,
					StartPos:  6,
					EndPos:    7,
				},
				{
					Atom:      "2",
					TokenType: token.TokenTypeNumber,
					StartPos:  7,
					EndPos:    8,
				},
				{
					Atom:      "\n",
					TokenType: token.TokenTypeReturn,
					StartPos:  8,
					EndPos:    9,
				},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 4",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "\n"),
			),
		},
		{
			name: "multiple commas",
			input: []*token.Token{
				{
					Atom:      "return",
					TokenType: token.TokenTypeReturn,
					StartPos:  0,
					EndPos:    6,
				},
				{
					Atom:      "1",
					TokenType: token.TokenTypeNumber,
					StartPos:  6,
					EndPos:    7,
				},
				{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
					StartPos:  7,
					EndPos:    8,
				},
				{
					Atom:      ",",
					TokenType: token.TokenTypeComma,
					StartPos:  8,
					EndPos:    9,
				},
				{
					Atom:      "\n",
					TokenType: token.TokenTypeReturn,
					StartPos:  8,
					EndPos:    9,
				},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 5",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "\n"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			_, err := parser.Parse()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
