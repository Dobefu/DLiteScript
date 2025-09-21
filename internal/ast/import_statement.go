package ast

import (
	"fmt"
)

// ImportStatement represents an import statement in the AST.
type ImportStatement struct {
	Path      *StringLiteral
	Namespace string
	StartPos  int
	EndPos    int
}

// Expr returns the expression of the import statement.
func (i *ImportStatement) Expr() string {
	return fmt.Sprintf("import %s", i.Path.Expr())
}

// StartPosition returns the start position of the import statement.
func (i *ImportStatement) StartPosition() int {
	return i.StartPos
}

// EndPosition returns the end position of the import statement.
func (i *ImportStatement) EndPosition() int {
	return i.EndPos
}

// Walk walks the import statement.
func (i *ImportStatement) Walk(fn func(node ExprNode) bool) {
	fn(i)
}
