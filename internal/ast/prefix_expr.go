package ast

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/token"
)

// PrefixExpr defines a struct for a prefix expression.
type PrefixExpr struct {
	Operator token.Token
	Operand  ExprNode
	Pos      int
}

// Expr returns the expression of the prefix expression.
func (e *PrefixExpr) Expr() string {
	return fmt.Sprintf("(%s %s)", e.Operator.Atom, e.Operand.Expr())
}

// Position returns the position of the prefix expression.
func (e *PrefixExpr) Position() int {
	return e.Pos
}
