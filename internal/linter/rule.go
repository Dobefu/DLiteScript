package linter

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

// Rule represents a linting rule that can analyze AST nodes.
type Rule interface {
	// Name returns the name of the rule.
	Name() string
	// Description returns a description of what the rule checks.
	Description() string
	// Analyze analyzes the given AST node and returns any issues found.
	Analyze(node ast.ExprNode)
}
