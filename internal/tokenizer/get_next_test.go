package tokenizer

import (
	"errors"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestGetNext(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected rune
	}{
		{
			input:    "1 + 1",
			expected: '1',
		},
		{
			input:    "🙂",
			expected: '🙂',
		},
	}

	for _, test := range tests {
		tokenizer := NewTokenizer(test.input)
		token, err := tokenizer.GetNext()

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if token != test.expected {
			t.Errorf("expected %v, got %v", test.expected, token)
		}
	}
}

func TestGetNextErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "\xFF",
			expected: errorutil.ErrorMsgInvalidUTF8Char,
		},
		{
			input:    "",
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
	}

	for _, test := range tests {
		tokenizer := NewTokenizer(test.input)
		_, err := tokenizer.GetNext()

		if err == nil {
			t.Fatalf("expected error, got nil")
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
