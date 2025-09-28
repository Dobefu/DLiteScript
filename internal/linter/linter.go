// Package linter provides a linter.
package linter

import (
	"io"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/linter/reporter"
	"github.com/Dobefu/DLiteScript/internal/linter/rules"
)

// Linter represents the main linter.
type Linter struct {
	reporter *reporter.Reporter
	rules    []Rule
	outFile  io.Writer
}

// New creates a new linter instance with default rules.
func New(outFile io.Writer) *Linter {
	reporter := reporter.NewReporter(outFile)

	return &Linter{
		reporter: reporter,
		rules: []Rule{
			rules.NewUnusedVariables(reporter),
			rules.NewUnreachableCode(reporter),
			rules.NewMissingReturn(reporter),
		},
		outFile: outFile,
	}
}

// Lint analyzes the given AST and returns a list of issues.
func (l *Linter) Lint(node ast.ExprNode) {
	for _, rule := range l.rules {
		rule.Analyze(node)
	}
}

// PrintIssues prints all issues found by the linter.
func (l *Linter) PrintIssues(filename string) {
	l.reporter.PrintIssues(filename)
}

// HasIssues returns true if there are any issues.
func (l *Linter) HasIssues() bool {
	return l.reporter.HasIssues()
}
