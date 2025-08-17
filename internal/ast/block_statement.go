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
	statements := []string{}

	for _, statement := range e.Statements {
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
