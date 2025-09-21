package ast

// CommentLiteral defines a struct for a comment literal.
type CommentLiteral struct {
	Value    string
	StartPos int
	EndPos   int
}

// Expr returns the expression of the comment literal.
func (e *CommentLiteral) Expr() string {
	return e.Value
}

// StartPosition returns the start position of the comment literal.
func (e *CommentLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the comment literal.
func (e *CommentLiteral) EndPosition() int {
	return e.EndPos
}

// Walk walks the comment literal.
func (e *CommentLiteral) Walk(_ func(node ExprNode) bool) {
	// noop
}
