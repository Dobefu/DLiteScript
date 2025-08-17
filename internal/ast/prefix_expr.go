package ast

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/token"
)

// PrefixExpr defines a struct for a prefix expression.
type PrefixExpr struct {
	Operator token.Token
	Operand  ExprNode
	StartPos int
	EndPos   int
}

// Expr returns the expression of the prefix expression.
func (e *PrefixExpr) Expr() string {
	return fmt.Sprintf("(%s %s)", e.Operator.Atom, e.Operand.Expr())
}

// StartPosition returns the start position of the prefix expression.
func (e *PrefixExpr) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the prefix expression.
func (e *PrefixExpr) EndPosition() int {
	return e.EndPos
}
