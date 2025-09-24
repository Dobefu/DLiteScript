package ast

import "strings"

// StatementList represents multiple statements separated by newlines.
type StatementList struct {
	Statements []ExprNode
	Range      Range
}

// Expr returns the expression in the statement list.
func (sl *StatementList) Expr() string {
	if len(sl.Statements) == 0 {
		return ""
	}

	var statements strings.Builder

	for i, statement := range sl.Statements {
		if statement == nil {
			continue
		}

		statements.WriteString(statement.Expr())

		if i < len(sl.Statements)-1 {
			statements.WriteString("\n")
		}
	}

	return statements.String()
}

// GetRange returns the range of the statement list.
func (sl *StatementList) GetRange() Range {
	return sl.Range
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
