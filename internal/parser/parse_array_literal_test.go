package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseArrayLiteral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "empty array",
			input: []*token.Token{
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "]", TokenType: token.TokenTypeRBracket},
			},
			expected: "[]",
		},
		{
			name: "array with one element",
			input: []*token.Token{
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "]", TokenType: token.TokenTypeRBracket},
			},
			expected: "[1]",
		},
		{
			name: "array with multiple elements",
			input: []*token.Token{
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ",", TokenType: token.TokenTypeComma},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "]", TokenType: token.TokenTypeRBracket},
			},
			expected: "[1, 2]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			expr, err := parser.Parse()

			if err != nil {
				t.Errorf("expected no error, got: \"%s\"", err.Error())
			}

			if expr.Expr() != test.expected {
				t.Errorf("expected: \"%s\", got: \"%s\"", test.expected, expr.Expr())
			}
		})
	}
}

func TestParseArrayLiteralErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "unclosed array",
			input: []*token.Token{
				{Atom: "[", TokenType: token.TokenTypeLBracket},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "missing closing bracket",
			input: []*token.Token{
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "1", TokenType: token.TokenTypeIdentifier},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 3",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token",
			input: []*token.Token{
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 3",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "1"),
			),
		},
		{
			name: "evaluation error",
			input: []*token.Token{
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "}", TokenType: token.TokenTypeLBrace},
				{Atom: "]", TokenType: token.TokenTypeRBracket},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 3",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "}"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			_, err := parser.Parse()

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Errorf("expected: \"%s\", got: \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestParseArrayLiteralExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:  "no tokens",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			parser := NewParser(test.input)
			_, err := parser.parseArrayLiteralExpr(0)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
