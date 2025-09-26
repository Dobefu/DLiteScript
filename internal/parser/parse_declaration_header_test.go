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
				t.Fatalf(
					"expected declared to be \"%s\", got \"%s\"",
					test.expectedName,
					declared,
				)
			}

			if varType != test.expectedType {
				t.Fatalf(
					"expected varType to be \"%s\", got \"%s\"",
					test.expectedType,
					varType,
				)
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
			name:  "unexpected EOF",
			input: []*token.Token{},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "invalid identifier",
			input: []*token.Token{
				token.NewToken("1", token.TokenTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedIdentifier, "1"),
			),
		},
		{
			name: "invalid data type",
			input: []*token.Token{
				token.NewToken("x", token.TokenTypeIdentifier, 0, 0),
				token.NewToken("bogus", token.TokenTypeString, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 7",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "bogus"),
			),
		},
		{
			name: "unexpected EOF after identifier",
			input: []*token.Token{
				token.NewToken("x", token.TokenTypeIdentifier, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
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

func TestParseDataType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "number",
			input: []*token.Token{
				token.NewToken("number", token.TokenTypeTypeNumber, 0, 0),
			},
			expected: "number",
		},
		{
			name: "array",
			input: []*token.Token{
				token.NewToken("[", token.TokenTypeLBracket, 0, 0),
				token.NewToken("]", token.TokenTypeRBracket, 0, 0),
				token.NewToken("number", token.TokenTypeTypeNumber, 0, 0),
			},
			expected: "[]number",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input[1:])
			dataType, err := p.parseDataType(test.input[0])

			if err != nil {
				t.Fatalf("expected no error, got %s", err.Error())
			}

			if dataType != test.expected {
				t.Fatalf(
					"expected dataType to be \"%s\", got \"%s\"",
					test.expected,
					dataType,
				)
			}
		})
	}
}

func TestParseDataTypeErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []*token.Token
		expected string
	}{
		{
			name: "unexpected EOF in array declaration",
			input: []*token.Token{
				token.NewToken("[", token.TokenTypeLBracket, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "unexpected token in array declaration",
			input: []*token.Token{
				token.NewToken("[", token.TokenTypeLBracket, 0, 0),
				token.NewToken("number", token.TokenTypeTypeNumber, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 7",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "number"),
			),
		},
		{
			name: "array declaration without type",
			input: []*token.Token{
				token.NewToken("[", token.TokenTypeLBracket, 0, 0),
				token.NewToken("]", token.TokenTypeRBracket, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 2",
				errorutil.StageParse.String(),
				errorutil.ErrorMsgUnexpectedEOF,
			),
		},
		{
			name: "array declaration with invalid type",
			input: []*token.Token{
				token.NewToken("[", token.TokenTypeLBracket, 0, 0),
				token.NewToken("]", token.TokenTypeRBracket, 0, 0),
				token.NewToken("bogus", token.TokenTypeIdentifier, 0, 0),
			},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 7",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "bogus"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(test.input[1:])
			_, err := p.parseDataType(test.input[0])

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected dataType to be \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
