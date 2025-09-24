package ast

import (
	"fmt"
	"strings"
)

// ArrayLiteral defines a struct for a literal array value.
type ArrayLiteral struct {
	Values []ExprNode
	Range  Range
}

// Expr returns the expression of the array literal.
func (e *ArrayLiteral) Expr() string {
	if len(e.Values) == 0 {
		return "[]"
	}

	var values strings.Builder

	for i, value := range e.Values {
		if value == nil {
			continue
		}

		values.WriteString(value.Expr())

		if i < len(e.Values)-1 {
			values.WriteString(", ")
		}
	}

	return fmt.Sprintf("[%s]", values.String())
}

// GetRange returns the range of the array literal.
func (e *ArrayLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the array literal and its values.
func (e *ArrayLiteral) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(e)

	if !shouldContinue {
		return
	}

	for _, value := range e.Values {
		shouldContinue = fn(value)

		if !shouldContinue {
			return
		}

		value.Walk(fn)
	}
}
