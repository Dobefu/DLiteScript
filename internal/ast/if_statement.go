package ast

import (
	"fmt"
)

// IfStatement defines a struct for an if statement.
type IfStatement struct {
	Condition ExprNode
	ThenBlock *BlockStatement
	ElseBlock *BlockStatement
	Range     Range
}

// Expr returns the expression of the if statement.
func (e *IfStatement) Expr() string {
	if e.Condition == nil || e.ThenBlock == nil {
		return ""
	}

	if e.ElseBlock == nil {
		return fmt.Sprintf(
			"if %s { %s }",
			e.Condition.Expr(),
			e.ThenBlock.Expr(),
		)
	}

	return fmt.Sprintf(
		"if %s { %s } else { %s }",
		e.Condition.Expr(),
		e.ThenBlock.Expr(),
		e.ElseBlock.Expr(),
	)
}

// GetRange returns the range of the if statement.
func (e *IfStatement) GetRange() Range {
	return e.Range
}

// Walk walks the if statement and its condition and body.
func (e *IfStatement) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(e)

	if !shouldContinue {
		return
	}

	if e.Condition != nil {
		shouldContinue = fn(e.Condition)

		if !shouldContinue {
			return
		}

		e.Condition.Walk(fn)
	}

	if e.ThenBlock != nil {
		shouldContinue = fn(e.ThenBlock)

		if !shouldContinue {
			return
		}

		e.ThenBlock.Walk(fn)
	}

	if e.ElseBlock != nil {
		shouldContinue = fn(e.ElseBlock)

		if !shouldContinue {
			return
		}

		e.ElseBlock.Walk(fn)
	}
}
