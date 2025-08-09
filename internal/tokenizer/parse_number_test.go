package tokenizer

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func parseNumberTestGetNumberToken(atom string) *token.Token {
	return token.NewToken(atom, token.TokenTypeNumber)
}

func TestParseNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected []*token.Token
	}{
		{
			input:    "1",
			expected: []*token.Token{parseNumberTestGetNumberToken("1")},
		},
		{
			input:    "1.1",
			expected: []*token.Token{parseNumberTestGetNumberToken("1.1")},
		},
		{
			input:    ".1",
			expected: []*token.Token{parseNumberTestGetNumberToken(".1")},
		},
		{
			input:    "1e1",
			expected: []*token.Token{parseNumberTestGetNumberToken("1e1")},
		},
		{
			input:    "1e+1",
			expected: []*token.Token{parseNumberTestGetNumberToken("1e1")},
		},
		{
			input:    "1e-1",
			expected: []*token.Token{parseNumberTestGetNumberToken("1e-1")},
		},
		{
			input: "1+1",
			expected: []*token.Token{
				parseNumberTestGetNumberToken("1"),
				token.NewToken("+", token.TokenTypeOperationAdd),
				parseNumberTestGetNumberToken("1"),
			},
		},
		{
			input:    "1_000_000",
			expected: []*token.Token{parseNumberTestGetNumberToken("1000000")},
		},
		{
			input:    "1.1_000_000",
			expected: []*token.Token{parseNumberTestGetNumberToken("1.1000000")},
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

func TestParseNumberErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1__2",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberMultipleUnderscores, "1__2"),
		},
		{
			input:    "1..1",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberMultipleDecimalPoints, "1..1"),
		},
		{
			input:    "1.",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1."),
		},
		{
			input:    "1e++",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e++"),
		},
	}

	for _, test := range tests {
		_, err := NewTokenizer(test.input).Tokenize()

		if err == nil {
			t.Fatalf("expected error for %s, got none", test.input)
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

func BenchmarkParseNumber(b *testing.B) {
	for b.Loop() {
		t := NewTokenizer("1 + -2 * 3 / 4")

		_, _ = t.parseNumber('1')
	}
}
