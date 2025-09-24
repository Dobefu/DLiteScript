package ast

// ExprNode defines a common interface signature for expression structs.
type ExprNode interface {
	Expr() string
	GetRange() Range
	Walk(func(node ExprNode) bool)
}
