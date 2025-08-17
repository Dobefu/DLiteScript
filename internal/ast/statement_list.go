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

// Walk walks the statement list and its statements.
func (sl *StatementList) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(sl)

	if !shouldContinue {
		return
	}

	for _, statement := range sl.Statements {
		shouldContinue = fn(statement)

		if !shouldContinue {
			return
		}

		statement.Walk(fn)
	}
}
