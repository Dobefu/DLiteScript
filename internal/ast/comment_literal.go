package ast

// CommentLiteral defines a struct for a comment literal.
type CommentLiteral struct {
	Value string
	Range Range
}

// Expr returns the expression of the comment literal.
func (e *CommentLiteral) Expr() string {
	return e.Value
}

// GetRange returns the range of the comment literal.
func (e *CommentLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the comment literal.
func (e *CommentLiteral) Walk(_ func(node ExprNode) bool) {
	// noop
}
