package ast

import (
	"fmt"
)

// ImportStatement represents an import statement in the AST.
type ImportStatement struct {
	Path      *StringLiteral
	Namespace string
	Alias     string
	Range     Range
}

// Expr returns the expression of the import statement.
func (i *ImportStatement) Expr() string {
	if i.Alias != "" {
		return fmt.Sprintf("import %s as %s", i.Path.Expr(), i.Alias)
	}

	return fmt.Sprintf("import %s", i.Path.Expr())
}

// GetRange returns the range of the import statement.
func (i *ImportStatement) GetRange() Range {
	return i.Range
}

// Walk walks the import statement.
func (i *ImportStatement) Walk(fn func(node ExprNode) bool) {
	fn(i)
}
