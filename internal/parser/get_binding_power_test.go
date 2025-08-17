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
			input:    token.NewToken("1", token.TokenTypeNumber, 0, 0),
			expected: bindingPowerDefault,
		},
		{
			input:    token.NewToken("+", token.TokenTypeOperationAdd, 0, 0),
			expected: bindingPowerAdditive,
		},
		{
			input:    token.NewToken("-", token.TokenTypeOperationSub, 0, 0),
			expected: bindingPowerAdditive,
		},
		{
			input:    token.NewToken("*", token.TokenTypeOperationMul, 0, 0),
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("/", token.TokenTypeOperationDiv, 0, 0),
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("%", token.TokenTypeOperationMod, 0, 0),
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("**", token.TokenTypeOperationPow, 0, 0),
			expected: bindingPowerPower,
		},
		{
			input:    token.NewToken("(", token.TokenTypeLParen, 0, 0),
			expected: bindingPowerParentheses,
		},
		{
			input:    token.NewToken(")", token.TokenTypeRParen, 0, 0),
			expected: bindingPowerParentheses,
		},
	}

	for _, test := range tests {
		bindingPower := NewParser([]*token.Token{test.input}).getBindingPower(test.input, false)

		if bindingPower != test.expected {
			t.Errorf(
				"expected binding power to be %d, got %d",
				test.expected,
				bindingPower,
			)
		}
	}
}
