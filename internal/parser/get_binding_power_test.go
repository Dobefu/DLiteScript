package parser

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestGetBindingPower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *token.Token
		expected int
	}{
		{
			input:    token.NewToken("1", token.TokenTypeNumber),
			expected: bindingPowerDefault,
		},
		{
			input:    token.NewToken("+", token.TokenTypeOperationAdd),
			expected: bindingPowerAdditive,
		},
		{
			input:    token.NewToken("-", token.TokenTypeOperationSub),
			expected: bindingPowerAdditive,
		},
		{
			input:    token.NewToken("*", token.TokenTypeOperationMul),
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("/", token.TokenTypeOperationDiv),
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("%", token.TokenTypeOperationMod),
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("**", token.TokenTypeOperationPow),
			expected: bindingPowerPower,
		},
		{
			input:    token.NewToken("(", token.TokenTypeLParen),
			expected: bindingPowerParentheses,
		},
		{
			input:    token.NewToken(")", token.TokenTypeRParen),
			expected: bindingPowerParentheses,
		},
	}

	for _, test := range tests {
		bindingPower := NewParser([]*token.Token{test.input}).getBindingPower(test.input, false)

		if bindingPower != test.expected {
			t.Errorf("expected binding power to be %d, got %d", test.expected, bindingPower)
		}
	}
}
