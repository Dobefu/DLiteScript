package parser

import (
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestParseDeclarationHeader(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        []*token.Token
		expectedName string
		expectedType string
	}{
		{
			name: "number declaration",
			input: []*token.Token{
				token.NewToken("x", token.TokenTypeIdentifier, 0, 0),
				token.NewToken("number", token.TokenTypeTypeNumber, 0, 0),
			},
			expectedName: "x",
			expectedType: "number",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			declared, varType, err := p.parseDeclarationHeader()

			if err != nil {
				t.Fatalf("expected error to be nil, got %s", err.Error())
			}

			if declared != test.expectedName {
				t.Fatalf("expected declared to be \"%s\", got \"%s\"", test.expectedName, declared)
			}

			if varType != test.expectedType {
				t.Fatalf("expected varType to be \"%s\", got \"%s\"", test.expectedType, varType)
			}
		})
	}
}

func TestParseDeclarationHeaderErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name:     "unexpected EOF",
			input:    []*token.Token{},
			expected: errorutil.ErrorMsgUnexpectedEOF + " at position 0",
		},
		{
			name: "invalid identifier",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedIdentifier, "1") + " at position 1",
		},
		{
			name: "invalid data type",
			input: []*token.Token{
				token.NewToken("x", token.TokenTypeIdentifier, 0, 0),
				token.NewToken("bogus", token.TokenTypeString, 0, 0),
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "bogus") + " at position 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input)
			_, _, err := p.parseDeclarationHeader()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
