package ast

import (
	"fmt"
	"strings"
)

// FuncStatement represents a function statement.
type FuncStatement struct {
	Name            string
	Params          []string
	Body            []ExprNode
	ReturnValues    []ExprNode
	NumReturnValues int
	StartPos        int
	EndPos          int
}

// Expr returns the expression of the function statement.
func (b *FuncStatement) Expr() string {
	if b.NumReturnValues == 0 {
		return fmt.Sprintf("func %s(%s)", b.Name, strings.Join(b.Params, ", "))
	}

	returnValues := make([]string, 0, b.NumReturnValues)

	for _, value := range b.ReturnValues {
		returnValues = append(returnValues, value.Expr())
	}

	return fmt.Sprintf(
		"func %s(%s) %s",
		b.Name,
		strings.Join(b.Params, ", "),
		strings.Join(returnValues, ", "),
	)
}

// StartPosition returns the start position of the function statement.
func (b *FuncStatement) StartPosition() int {
	return b.StartPos
}

// EndPosition returns the end position of the function statement.
func (b *FuncStatement) EndPosition() int {
	return b.EndPos
}

// Walk walks the function statement.
func (b *FuncStatement) Walk(fn func(node ExprNode) bool) {
	fn(b)
}
