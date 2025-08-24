package ast

import (
	"fmt"
	"strings"
)

// ReturnStatement represents a return statement.
type ReturnStatement struct {
	Values    []ExprNode
	NumValues int
	StartPos  int
	EndPos    int
}

// Expr returns the expression of the return statement.
func (b *ReturnStatement) Expr() string {
	if b.NumValues == 0 {
		return "return"
	}

	values := make([]string, 0, b.NumValues)

	for _, value := range b.Values {
		values = append(values, value.Expr())
	}

	return fmt.Sprintf("return %s", strings.Join(values, ", "))
}

// StartPosition returns the start position of the return statement.
func (b *ReturnStatement) StartPosition() int {
	return b.StartPos
}

// EndPosition returns the end position of the return statement.
func (b *ReturnStatement) EndPosition() int {
	return b.EndPos
}

// Walk walks the return statement.
func (b *ReturnStatement) Walk(fn func(node ExprNode) bool) {
	fn(b)
}
