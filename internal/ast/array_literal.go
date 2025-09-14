package ast

import (
	"fmt"
	"strings"
)

// ArrayLiteral defines a struct for a literal array value.
type ArrayLiteral struct {
	Values   []ExprNode
	StartPos int
	EndPos   int
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

// StartPosition returns the start position of the array literal.
func (e *ArrayLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the array literal.
func (e *ArrayLiteral) EndPosition() int {
	return e.EndPos
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
