package parser

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: "1",
		},
		{
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "\n", TokenType: token.TokenTypeNewline},
			},
			expected: "1",
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "10", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "5", TokenType: token.TokenTypeNumber},
				{Atom: "**", TokenType: token.TokenTypeOperationPow},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				{Atom: "5", TokenType: token.TokenTypeNumber},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "123", TokenType: token.TokenTypeNumber},
				{Atom: "%", TokenType: token.TokenTypeOperationMod},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "5", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
			expected: "((1 + (10 / ((5 ** 2) * 5))) - ((- 123) % (- 5)))",
		},
	}

	for _, test := range tests {
		parser := NewParser(test.input)
		result, err := parser.Parse()

		if err != nil {
			t.Errorf("expected no error, got '%s'", err.Error())

			continue
		}

		if result.Expr() != test.expected {
			t.Errorf("expected '%s', got '%s'", test.expected, result.Expr())
		}
	}
}

func TestParseErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []*token.Token
		expected string
	}{
		{
			input:    []*token.Token{},
			expected: errorutil.ErrorMsgEmptyExpression,
		},
		{
			input: []*token.Token{
				{Atom: "_", TokenType: token.TokenTypeNumber},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
			},
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input: []*token.Token{
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "/"),
		},
		{
			input: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "2", TokenType: token.TokenTypeNumber},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "2"),
		},
	}

	for _, test := range tests {
		p := NewParser(test.input)

		// Test the EOF error.
		if len(test.input) == 1 && test.input[0].Atom == "_" {
			p.isEOF = true
		}

		_, err := p.Parse()

		if err == nil {
			t.Fatalf("expected error for \"%v\", got none", test.input)
		}

		if errors.Unwrap(err).Error() != test.expected {
			t.Errorf(
				"expected error \"%s\", got \"%s\"",
				test.expected,
				errors.Unwrap(err).Error(),
			)
		}
	}
}

func BenchmarkParse(b *testing.B) {
	for b.Loop() {
		p := NewParser(
			[]*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "2", TokenType: token.TokenTypeNumber},
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				{Atom: "3", TokenType: token.TokenTypeNumber},
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				{Atom: "4", TokenType: token.TokenTypeNumber},
			},
		)
		_, _ = p.Parse()
	}
}
