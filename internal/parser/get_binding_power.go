package parser

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

const (
	bindingPowerParentheses    = 1000
	bindingPowerUnary          = 900
	bindingPowerPower          = 800
	bindingPowerMultiplicative = 700
	bindingPowerAdditive       = 600
	bindingPowerComparison     = 500
	bindingPowerLogicalAnd     = 400
	bindingPowerLogicalOr      = 300
	bindingPowerAssignment     = 10

	// For right-hand associativity, a value of 1 is subtracted from the
	// binding power of the next token.
	// To prevent the binding power from being negative, a value of 1 is
	// added to the default binding power.
	bindingPowerDefault = 1
)

// getBindingPower returns the binding power of the current token.
func (p *Parser) getBindingPower(currentToken *token.Token, isUnary bool) int {
	switch currentToken.TokenType {
	case
		token.TokenTypeLParen,
		token.TokenTypeRParen:
		return bindingPowerParentheses

	case
		token.TokenTypeNot:
		if isUnary {
			return bindingPowerUnary
		}

		return bindingPowerDefault

	case
		token.TokenTypeOperationPow:
		return bindingPowerPower

	case token.TokenTypeOperationMul,
		token.TokenTypeOperationDiv,
		token.TokenTypeOperationMod:
		return bindingPowerMultiplicative

	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub:
		if isUnary {
			return bindingPowerUnary
		}

		return bindingPowerAdditive

	case
		token.TokenTypeEqual,
		token.TokenTypeNotEqual,
		token.TokenTypeGreaterThan,
		token.TokenTypeGreaterThanOrEqual,
		token.TokenTypeLessThan,
		token.TokenTypeLessThanOrEqual:
		return bindingPowerComparison

	case
		token.TokenTypeLogicalAnd:
		return bindingPowerLogicalAnd

	case
		token.TokenTypeLogicalOr:
		return bindingPowerLogicalOr

	case
		token.TokenTypeAssign:
		return bindingPowerAssignment

	default:
		return bindingPowerDefault
	}
}
