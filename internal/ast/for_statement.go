package ast

import (
	"fmt"
)

// ForStatement represents a for statement.
type ForStatement struct {
	Condition        ExprNode
	Body             *BlockStatement
	StartPos         int
	EndPos           int
	DeclaredVariable string
	RangeVariable    string
	RangeFrom        ExprNode
	RangeTo          ExprNode
	IsRange          bool
}

// Expr returns the expression of the for statement.
func (f *ForStatement) Expr() string {
	if f.Body == nil {
		return "for { }"
	}

	if f.IsRange {
		if f.RangeFrom != nil && f.RangeTo != nil {
			if f.DeclaredVariable != "" {
				return fmt.Sprintf(
					"for var %s from %s to %s { %s }",
					f.DeclaredVariable,
					f.RangeFrom.Expr(),
					f.RangeTo.Expr(),
					f.Body.Expr(),
				)
			}

			return fmt.Sprintf(
				"for from %s to %s { %s }",
				f.RangeFrom.Expr(),
				f.RangeTo.Expr(),
				f.Body.Expr(),
			)
		}

		if f.DeclaredVariable != "" {
			return fmt.Sprintf(
				"for var %s to %s { %s }",
				f.DeclaredVariable,
				f.RangeTo.Expr(),
				f.Body.Expr(),
			)
		}

		return fmt.Sprintf(
			"for from 0 to %s { %s }",
			f.RangeTo.Expr(),
			f.Body.Expr(),
		)
	}

	if f.Condition == nil {
		return fmt.Sprintf("for { %s }", f.Body.Expr())
	}

	if f.DeclaredVariable != "" {
		return fmt.Sprintf(
			"for var %s %s { %s }",
			f.DeclaredVariable,
			f.Condition.Expr(),
			f.Body.Expr(),
		)
	}

	return fmt.Sprintf("for %s { %s }", f.Condition.Expr(), f.Body.Expr())
}

// StartPosition returns the start position of the for statement.
func (f *ForStatement) StartPosition() int {
	return f.StartPos
}

// EndPosition returns the end position of the for statement.
func (f *ForStatement) EndPosition() int {
	return f.EndPos
}

// Walk walks the for statement and its condition and body.
func (f *ForStatement) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(f)

	if !shouldContinue {
		return
	}

	if f.Condition != nil {
		shouldContinue = fn(f.Condition)

		if !shouldContinue {
			return
		}

		f.Condition.Walk(fn)
	}

	if f.RangeFrom != nil {
		shouldContinue = fn(f.RangeFrom)

		if !shouldContinue {
			return
		}

		f.RangeFrom.Walk(fn)
	}

	if f.RangeTo != nil {
		shouldContinue = fn(f.RangeTo)

		if !shouldContinue {
			return
		}

		f.RangeTo.Walk(fn)
	}

	if f.Body != nil {
		shouldContinue = fn(f.Body)

		if !shouldContinue {
			return
		}

		f.Body.Walk(fn)
	}
}
