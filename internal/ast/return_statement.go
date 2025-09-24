package ast

import (
	"fmt"
	"strings"
)

// ReturnStatement represents a return statement.
type ReturnStatement struct {
	Values    []ExprNode
	NumValues int
	Range     Range
}

// Expr returns the expression of the return statement.
func (b *ReturnStatement) Expr() string {
	if b.NumValues == 0 {
		return "return"
	}

	values := make([]string, 0, b.NumValues)

	for _, value := range b.Values {
		if value == nil {
			continue
		}

		values = append(values, value.Expr())
	}

	return fmt.Sprintf("return %s", strings.Join(values, ", "))
}

// GetRange returns the range of the return statement.
func (b *ReturnStatement) GetRange() Range {
	return b.Range
}

// Walk walks the return statement.
func (b *ReturnStatement) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(b)

	if !shouldContinue {
		return
	}

	for _, value := range b.Values {
		shouldContinue = fn(value)

		if !shouldContinue {
			return
		}

		value.Walk(fn)
	}
}
