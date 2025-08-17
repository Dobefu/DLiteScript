package ast

import "strings"

// StatementList represents multiple statements separated by newlines.
type StatementList struct {
	Statements []ExprNode
	StartPos   int
	EndPos     int
}

// Expr returns the expression in the statement list.
func (sl *StatementList) Expr() string {
	var statements strings.Builder

	for i, statement := range sl.Statements {
		statements.WriteString(statement.Expr())

		if i < len(sl.Statements)-1 {
			statements.WriteString("\n")
		}
	}

	return statements.String()
}

// StartPosition returns the start position of the statement list.
func (sl *StatementList) StartPosition() int {
	return sl.StartPos
}

// EndPosition returns the end position of the statement list.
func (sl *StatementList) EndPosition() int {
	return sl.EndPos
}
