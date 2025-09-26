package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseConstantDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "constant declaration",
			input: []*token.Token{
				{Atom: "const", TokenType: token.TokenTypeConst},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "const x number = 1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			result, err := p.Parse()

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if result.Expr() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, result.Expr())
			}
		})
	}
}

func TestParseConstantDeclarationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "unexpected EOF after const",
			input: []*token.Token{
				{Atom: "const", TokenType: token.TokenTypeConst},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 6",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected EOF after declaration",
			input: []*token.Token{
				{Atom: "const", TokenType: token.TokenTypeConst},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 13",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "no assignment operator after declaration",
			input: []*token.Token{
				{Atom: "const", TokenType: token.TokenTypeConst},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "\n", TokenType: token.TokenTypeNewline},
			},
			expected: fmt.Sprintf(
				"%s: %s line 2 at position 1",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgConstantDeclarationWithNoValue, "x"),
			),
		},
		{
			name: "unexpected EOF after value",
			input: []*token.Token{
				{Atom: "const", TokenType: token.TokenTypeConst},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 14",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected value",
			input: []*token.Token{
				{Atom: "const", TokenType: token.TokenTypeConst},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 15",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "}"),
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
