package ast

import (
	"fmt"
	"strings"
)

// FuncDeclarationStatement represents a function declaration statement.
type FuncDeclarationStatement struct {
	Name            string
	Args            []FuncParameter
	Body            ExprNode
	ReturnValues    []string
	NumReturnValues int
	StartPos        int
	EndPos          int
}

// Expr returns the expression of the function declaration statement.
func (b *FuncDeclarationStatement) Expr() string {
	argStrings := make([]string, len(b.Args))

	for i, arg := range b.Args {
		argStrings[i] = arg.Name + " " + arg.Type
	}

	if b.NumReturnValues == 0 {
		return fmt.Sprintf("func %s(%s)", b.Name, strings.Join(argStrings, ", "))
	}

	return fmt.Sprintf(
		"func %s(%s) %s",
		b.Name,
		strings.Join(argStrings, ", "),
		strings.Join(b.ReturnValues, ", "),
	)
}

// StartPosition returns the start position of the function declaration statement.
func (b *FuncDeclarationStatement) StartPosition() int {
	return b.StartPos
}

// EndPosition returns the end position of the function declaration statement.
func (b *FuncDeclarationStatement) EndPosition() int {
	return b.EndPos
}

// Walk walks the function declaration statement.
func (b *FuncDeclarationStatement) Walk(fn func(node ExprNode) bool) {
	fn(b)
}
