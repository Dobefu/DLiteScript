package ast

import (
	"fmt"
	"strings"
)

// BlockStatement defines a struct for a block statement.
type BlockStatement struct {
	Statements []ExprNode
	Pos        int
}

// Expr returns the expression of the block statement.
func (e *BlockStatement) Expr() string {
	statements := []string{}

	for _, statement := range e.Statements {
		statements = append(statements, statement.Expr())
	}

	return fmt.Sprintf("(%s)", strings.Join(statements, " "))
}

// Position returns the position of the block statement.
func (e *BlockStatement) Position() int {
	return e.Pos
}
