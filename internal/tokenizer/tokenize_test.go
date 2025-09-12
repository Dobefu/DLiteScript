package tokenizer

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func tokenizeTestGetNumberToken(atom string) *token.Token {
	return token.NewToken(atom, token.TokenTypeNumber, 0, 0)
}

func TestTokenize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []*token.Token
	}{
		{
			name:  "number with newline",
			input: "1\n2",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("1"),
				{Atom: "\n", TokenType: token.TokenTypeNewline},
				tokenizeTestGetNumberToken("2"),
			},
		},
		{
			name:     "number",
			input:    "1",
			expected: []*token.Token{tokenizeTestGetNumberToken("1")},
		},
		{
			name:     "number with zero exponent",
			input:    "1e0",
			expected: []*token.Token{tokenizeTestGetNumberToken("1e0")},
		},
		{
			name:     "number with exponent",
			input:    "1e5",
			expected: []*token.Token{tokenizeTestGetNumberToken("1e5")},
		},
		{
			name:     "number with positive exponent",
			input:    "1e+6",
			expected: []*token.Token{tokenizeTestGetNumberToken("1e6")},
		},
		{
			name:     "number with negative exponent",
			input:    "1.2E-8",
			expected: []*token.Token{tokenizeTestGetNumberToken("1.2E-8")},
		},
		{
			name:  "number with addition",
			input: "1 + 1",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("1"),
				{Atom: "+", TokenType: token.TokenTypeOperationAdd},
				tokenizeTestGetNumberToken("1"),
			},
		},
		{
			name:  "number with power",
			input: "2 ** 2",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("2"),
				{Atom: "**", TokenType: token.TokenTypeOperationPow},
				tokenizeTestGetNumberToken("2"),
			},
		},
		{
			name:  "number with modulo",
			input: "10 % 3",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("10"),
				{Atom: "%", TokenType: token.TokenTypeOperationMod},
				tokenizeTestGetNumberToken("3"),
			},
		},
		{
			name:  "number with division",
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
			name:  "number with subtraction",
			input: "4 - 5",
			expected: []*token.Token{
				tokenizeTestGetNumberToken("4"),
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				tokenizeTestGetNumberToken("5"),
			},
		},
		{
			name:  "number with parentheses",
			input: "()",
			expected: []*token.Token{
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
		},
		{
			name:  "number with function call",
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
		{
			name:  "string",
			input: `"test"`,
			expected: []*token.Token{
				{Atom: "test", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with escape",
			input: `"te\"st"`,
			expected: []*token.Token{
				{Atom: "te\"st", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with newline",
			input: `"\n"`,
			expected: []*token.Token{
				{Atom: "\n", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with tab",
			input: `"\t"`,
			expected: []*token.Token{
				{Atom: "\t", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with carriage return",
			input: `"\r"`,
			expected: []*token.Token{
				{Atom: "\r", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with null",
			input: `"\0"`,
			expected: []*token.Token{
				{Atom: "\000", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with backspace",
			input: `"\b"`,
			expected: []*token.Token{
				{Atom: "\b", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with form feed",
			input: `"\f"`,
			expected: []*token.Token{
				{Atom: "\f", TokenType: token.TokenTypeString},
			},
		},
		{
			name:  "string with vertical tab",
			input: `"\v"`,
			expected: []*token.Token{
				{Atom: "\v", TokenType: token.TokenTypeString},
			},
		},
		{
			name:     "string with newline",
			input:    "//\n",
			expected: []*token.Token{},
		},
		{
			name:     "string with comment",
			input:    "// Comment",
			expected: []*token.Token{},
		},
		{
			name:  "variable declaration",
			input: "var x number = 1",
			expected: []*token.Token{
				{Atom: "var", TokenType: token.TokenTypeVar},
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "number", TokenType: token.TokenTypeTypeNumber},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
		},
		{
			name:  "boolean not",
			input: "!true",
			expected: []*token.Token{
				{Atom: "!", TokenType: token.TokenTypeNot},
				{Atom: "true", TokenType: token.TokenTypeBool},
			},
		},
		{
			name:  "boolean less than",
			input: "1 < 1",
			expected: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: "<", TokenType: token.TokenTypeLessThan},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
		},
		{
			name:  "boolean greater than",
			input: "1 > 1",
			expected: []*token.Token{
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ">", TokenType: token.TokenTypeGreaterThan},
				{Atom: "1", TokenType: token.TokenTypeNumber},
			},
		},
		{
			name:  "boolean logical and",
			input: "true && true",
			expected: []*token.Token{
				{Atom: "true", TokenType: token.TokenTypeBool},
				{Atom: "&&", TokenType: token.TokenTypeLogicalAnd},
				{Atom: "true", TokenType: token.TokenTypeBool},
			},
		},
		{
			name:  "boolean logical or",
			input: "true || true",
			expected: []*token.Token{
				{Atom: "true", TokenType: token.TokenTypeBool},
				{Atom: "||", TokenType: token.TokenTypeLogicalOr},
				{Atom: "true", TokenType: token.TokenTypeBool},
			},
		},
		{
			name:  "boolean brace",
			input: "{true}",
			expected: []*token.Token{
				{Atom: "{", TokenType: token.TokenTypeLBrace},
				{Atom: "true", TokenType: token.TokenTypeBool},
				{Atom: "}", TokenType: token.TokenTypeRBrace},
			},
		},
		{
			name:  "function call",
			input: "math.abs(-1)",
			expected: []*token.Token{
				{Atom: "math", TokenType: token.TokenTypeIdentifier},
				{Atom: ".", TokenType: token.TokenTypeDot},
				{Atom: "abs", TokenType: token.TokenTypeIdentifier},
				{Atom: "(", TokenType: token.TokenTypeLParen},
				{Atom: "-", TokenType: token.TokenTypeOperationSub},
				{Atom: "1", TokenType: token.TokenTypeNumber},
				{Atom: ")", TokenType: token.TokenTypeRParen},
			},
		},
		{
			name:  "array assignment",
			input: "x = []",
			expected: []*token.Token{
				{Atom: "x", TokenType: token.TokenTypeIdentifier},
				{Atom: "=", TokenType: token.TokenTypeAssign},
				{Atom: "[", TokenType: token.TokenTypeLBracket},
				{Atom: "]", TokenType: token.TokenTypeRBracket},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
		})
	}
}

func TestTokenizeErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "incomplete exponent",
			input:    "1e",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e"),
		},
		{
			name:     "incomplete exponent with negative sign",
			input:    "1e-",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e-"),
		},
		{
			name:     "incomplete exponent with trailing character",
			input:    "1e-r",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e-r"),
		},
		{
			name:     "multiple exponent signs",
			input:    "1e6e6",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberMultipleExponentSigns, "1e6e6"),
		},
		{
			name:     "incomplete exponent with trailing character",
			input:    "1e6er",
			expected: fmt.Sprintf(errorutil.ErrorMsgNumberTrailingChar, "1e6er"),
		},
		{
			name:     "unexpected character",
			input:    "💔",
			expected: fmt.Sprintf(errorutil.ErrorMsgUnexpectedChar, "💔"),
		},
		{
			name:     "unexpected end of expression",
			input:    "*",
			expected: errorutil.ErrorMsgUnexpectedEOF,
		},
		{
			name:     "invalid UTF-8 character",
			input:    "1_e\x80",
			expected: string(errorutil.ErrorMsgInvalidUTF8Char),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
		})
	}
}

func BenchmarkTokenize(b *testing.B) {
	for b.Loop() {
		_, _ = NewTokenizer("1 + -2 * 3 / 4").Tokenize()
	}
}
