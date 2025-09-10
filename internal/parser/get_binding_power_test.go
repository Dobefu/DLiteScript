package parser

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestGetBindingPower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *token.Token
		isUnary  bool
		expected int
	}{
		{
			input:    token.NewToken("1", token.TokenTypeNumber, 0, 0),
			isUnary:  false,
			expected: bindingPowerDefault,
		},
		{
			input:    token.NewToken("+", token.TokenTypeOperationAdd, 0, 0),
			isUnary:  false,
			expected: bindingPowerAdditive,
		},
		{
			input:    token.NewToken("-", token.TokenTypeOperationSub, 0, 0),
			isUnary:  false,
			expected: bindingPowerAdditive,
		},
		{
			input:    token.NewToken("*", token.TokenTypeOperationMul, 0, 0),
			isUnary:  false,
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("/", token.TokenTypeOperationDiv, 0, 0),
			isUnary:  false,
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("%", token.TokenTypeOperationMod, 0, 0),
			isUnary:  false,
			expected: bindingPowerMultiplicative,
		},
		{
			input:    token.NewToken("[", token.TokenTypeLBracket, 0, 0),
			isUnary:  false,
			expected: bindingPowerArray,
		},
		{
			input:    token.NewToken("**", token.TokenTypeOperationPow, 0, 0),
			isUnary:  false,
			expected: bindingPowerPower,
		},
		{
			input:    token.NewToken("(", token.TokenTypeLParen, 0, 0),
			isUnary:  false,
			expected: bindingPowerParentheses,
		},
		{
			input:    token.NewToken(")", token.TokenTypeRParen, 0, 0),
			isUnary:  false,
			expected: bindingPowerParentheses,
		},
		{
			input:    token.NewToken("!", token.TokenTypeNot, 0, 0),
			isUnary:  false,
			expected: bindingPowerDefault,
		},
		{
			input:    token.NewToken("!", token.TokenTypeNot, 0, 0),
			isUnary:  true,
			expected: bindingPowerUnary,
		},
		{
			input:    token.NewToken("=", token.TokenTypeAssign, 0, 0),
			isUnary:  false,
			expected: bindingPowerAssignment,
		},
		{
			input:    token.NewToken("==", token.TokenTypeEqual, 0, 0),
			isUnary:  false,
			expected: bindingPowerComparison,
		},
		{
			input:    token.NewToken("&&", token.TokenTypeLogicalAnd, 0, 0),
			isUnary:  false,
			expected: bindingPowerLogicalAnd,
		},
		{
			input:    token.NewToken("||", token.TokenTypeLogicalOr, 0, 0),
			isUnary:  false,
			expected: bindingPowerLogicalOr,
		},
	}

	for _, test := range tests {
		t.Run(test.input.Atom, func(t *testing.T) {
			t.Parallel()

			bindingPower := NewParser([]*token.Token{test.input}).getBindingPower(
				test.input,
				test.isUnary,
			)

			if bindingPower != test.expected {
				t.Errorf(
					"expected binding power to be %d, got %d",
					test.expected,
					bindingPower,
				)
			}
		})
	}
}
