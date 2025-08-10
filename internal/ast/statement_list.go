package ast

import "strings"

// StatementList represents multiple statements separated by newlines.
type StatementList struct {
	Statements []ExprNode
	Pos        int
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

// Position returns the position of the statement list.
func (sl *StatementList) Position() int {
	return sl.Pos
}
