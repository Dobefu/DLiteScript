package ast

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/token"
)

// BinaryExpr defines a struct for a binary expression.
type BinaryExpr struct {
	Left     ExprNode
	Right    ExprNode
	Operator token.Token
	Pos      int
}

// Expr returns the expression of the binary expression.
func (e *BinaryExpr) Expr() string {
	return fmt.Sprintf("(%s %s %s)", e.Left.Expr(), e.Operator.Atom, e.Right.Expr())
}

// Position returns the position of the binary expression.
func (e *BinaryExpr) Position() int {
	return e.Pos
}
