package tokenizer

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func tokenizeTestGetNumberToken(atom string) *token.Token {
	return token.NewToken(atom, token.TokenTypeNumber)
}

func TestTokenize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected []*token.Token
	}{
		{
			input: "1\n2",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("1"),
				{Atom: "\n", TokenType: token.TokenTypeNewline},
				tokenizeTestGetNumberToken("2"),
			},
		},
		{
			input:    "1",
			expected: []*token.Token{tokenizeTestGetNumberToken("1")},
		},
		{
			input:    "1e0",
			expected: []*token.Token{tokenizeTestGetNumberToken("1e0")},
		},
		{
			input:    "1e5",
			expected: []*token.Token{tokenizeTestGetNumberToken("1e5")},
		},
		{
			input:    "1e+6",
			expected: []*token.Token{tokenizeTestGetNumberToken("1e6")},
		},
		{
			input:    "1.2E-8",
			expected: []*token.Token{tokenizeTestGetNumberToken("1.2E-8")},
		},
		{
			input: "1 + 1",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("1"),
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				tokenizeTestGetNumberToken("1"),
			},
		},
		{
			input: "2 ** 2",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("2"),
				{Atom: "**", TokenType: token.TokenTypeOperationPow},
				tokenizeTestGetNumberToken("2"),
			},
		},
		{
			input: "10 % 3",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("10"),
				{Atom: "%", TokenType: token.TokenTypeOperationMod},
				tokenizeTestGetNumberToken("3"),
			},
		},
		{
			input: "1 + 2 * 3 / 4",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("1"),
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				tokenizeTestGetNumberToken("2"),
				{Atom: "*", TokenType: token.TokenTypeOperationMul},
				tokenizeTestGetNumberToken("3"),
				{Atom: "/", TokenType: token.TokenTypeOperationDiv},
				tokenizeTestGetNumberToken("4"),
			},
		},
		{
			input: "4 - 5",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("4"),
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				tokenizeTestGetNumberToken("5"),
			},
		},
		{
			input: "()",
			expected: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
		},
		{
			input: "min(1, 2)",
			expected: []*token.Token{
				{Atom: "min", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				tokenizeTestGetNumberToken("1"),
				{Atom: ",", TokenType: token.TokenTypeComma},
				tokenizeTestGetNumberToken("2"),
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
		},
	}

	for _, test := range tests {
		tokens, err := NewTokenizer(test.input).Tokenize()

		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tokens, test.expected) {
			t.Fatalf("expected %v, got %v", test.expected, tokens)
		}
	}
}

func TestTokenizeErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1e",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e"),
		},
		{
			input:    "1e-",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e-"),
		},
		{
			input:    "1e-r",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e-r"),
		},
		{
			input:    "1e6e6",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberMultipleExponentSigns, "1e6e6"),
		},
		{
			input:    "1e6er",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e6er"),
		},
		{
			input:    "💔",
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedChar, "💔"),
		},
		{
			input:    "*",
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			input:    "1_e\x80",
			expected: string(errorutil.ErrorMsgInvalidUTF8Char),
		},
	}

	for _, test := range tests {
		_, err := NewTokenizer(test.input).Tokenize()

		if err == nil {
			t.Fatalf("expected error, got none for input %s", test.input)
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

func BenchmarkTokenize(b *testing.B) {
	for b.Loop() {
		_, _ = NewTokenizer("1 + -2 * 3 / 4").Tokenize()
	}
}
