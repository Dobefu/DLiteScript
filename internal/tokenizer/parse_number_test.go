package tokenizer

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func parseNumberTestGetNumberToken(atom string) *token.Token {
	return token.NewToken(atom, token.TokenTypeNumber, 0, 0)
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
				token.NewToken("+", token.TokenTypeOperationAdd, 0, 0),
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

		if len(tokens) != len(test.expected) {
			t.Fatalf("expected %d tokens, got %d", len(test.expected), len(tokens))
		}

		for i, token := range tokens {
			if token.Atom != test.expected[i].Atom {
				t.Fatalf(
					"expected token %d to be %s, got %s",
					i,
					test.expected[i].Atom,
					token.Atom,
				)
			}
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

		_, _ = t.parseNumber('1', 0)
	}
}
