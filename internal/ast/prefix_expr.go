package ast

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/token"
)

// PrefixExpr defines a struct for a prefix expression.
type PrefixExpr struct {
	Operator token.Token
	Operand  ExprNode
	Range    Range
}

// Expr returns the expression of the prefix expression.
func (e *PrefixExpr) Expr() string {
	if e.Operand == nil {
		return ""
	}

	return fmt.Sprintf("(%s %s)", e.Operator.Atom, e.Operand.Expr())
}

// GetRange returns the range of the prefix expression.
func (e *PrefixExpr) GetRange() Range {
	return e.Range
}

// Walk walks the prefix expression and its operand.
func (e *PrefixExpr) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(e)

	if !shouldContinue {
		return
	}

	if e.Operand != nil {
		shouldContinue = fn(e.Operand)

		if !shouldContinue {
			return
		}

		e.Operand.Walk(fn)
	}
}
