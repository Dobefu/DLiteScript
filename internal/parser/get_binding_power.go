package parser

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

const (
	bindingPowerParentheses    = 1000
	bindingPowerPower          = 400
	bindingPowerUnary          = 300
	bindingPowerMultiplicative = 200
	bindingPowerAdditive       = 100

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

	default:
		return bindingPowerDefault
	}
}
