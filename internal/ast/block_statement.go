package ast

import (
	"fmt"
	"strings"
)

// BlockStatement defines a struct for a block statement.
type BlockStatement struct {
	Statements []ExprNode
	StartPos   int
	EndPos     int
}

// Expr returns the expression of the block statement.
func (e *BlockStatement) Expr() string {
	if len(e.Statements) == 0 {
		return "()"
	}

	statements := []string{}

	for _, statement := range e.Statements {
		if statement == nil {
			continue
		}

		statements = append(statements, statement.Expr())
	}

	return fmt.Sprintf("(%s)", strings.Join(statements, " "))
}

// StartPosition returns the start position of the block statement.
func (e *BlockStatement) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the block statement.
func (e *BlockStatement) EndPosition() int {
	return e.EndPos
}

// Walk walks the block statement and its statements.
func (e *BlockStatement) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(e)

	if !shouldContinue {
		return
	}

	for _, statement := range e.Statements {
		shouldContinue = fn(statement)

		if !shouldContinue {
			return
		}

		statement.Walk(fn)
	}
}
