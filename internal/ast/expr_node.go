package ast

// ExprNode defines a common interface signature for expression structs.
type ExprNode interface {
	Expr() string
	StartPosition() int
	EndPosition() int
}
