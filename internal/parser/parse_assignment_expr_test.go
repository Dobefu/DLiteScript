package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseAssignmentExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "assignment expression",
			input: []*token.Token{
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "x = 1",
		},
		{
			name: "array assignment expression",
			input: []*token.Token{
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "]", TokenType: token.TokenTypeRBracket},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "2", TokenType: token.TokenTypeNumber},
			},
			expected: "x[1] = 2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			expr, err := p.Parse()

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if expr.Expr() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, expr.Expr())
			}
		})
	}
}

func TestParseAssignmentExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "unexpected token",
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 1",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "1"),
			),
		},
		{
			name: "no next token after assignment",
			input: []*token.Token{
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "=", TokenType: token.TokenTypeAssign},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 2",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid expression after assignment",
			input: []*token.Token{
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 4",
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
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, err.Error())
			}
		})
	}
}
