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
	Range           Range
}

// Expr returns the expression of the function declaration statement.
func (b *FuncDeclarationStatement) Expr() string {
	argStrings := make([]string, len(b.Args))

	for i, arg := range b.Args {
		argStrings[i] = fmt.Sprintf("%s %s", arg.Name, arg.Type)
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

// GetRange returns the range of the function declaration statement.
func (b *FuncDeclarationStatement) GetRange() Range {
	return b.Range
}

// Walk walks the function declaration statement.
func (b *FuncDeclarationStatement) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(b)

	if !shouldContinue {
		return
	}

	if b.Body != nil {
		shouldContinue = fn(b.Body)

		if !shouldContinue {
			return
		}

		b.Body.Walk(fn)
	}
}
