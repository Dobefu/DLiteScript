package rules

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/linter/reporter"
)

// UnusedVariables checks for unused variable declarations.
type UnusedVariables struct {
	name        string
	description string
	reporter    *reporter.Reporter
}

// NewUnusedVariables creates a new unused variables rule.
func NewUnusedVariables(reporter *reporter.Reporter) *UnusedVariables {
	return &UnusedVariables{
		name:        "unused-variables",
		description: "Detects variables that are declared but never used",
		reporter:    reporter,
	}
}

// Name returns the name of the rule.
func (r *UnusedVariables) Name() string {
	return r.name
}

// Description returns the description of the rule.
func (r *UnusedVariables) Description() string {
	return r.description
}

// Analyze analyzes the AST for unused variables.
func (r *UnusedVariables) Analyze(node ast.ExprNode) {
	variables := make(map[string]*ast.VariableDeclaration)
	constants := make(map[string]*ast.ConstantDeclaration)

	node.Walk(func(n ast.ExprNode) bool {
		switch decl := n.(type) {
		case *ast.VariableDeclaration:
			variables[decl.Name] = decl

		case *ast.ConstantDeclaration:
			constants[decl.Name] = decl
		}

		return true
	})

	usage := r.collectUsage(node)

	for name, decl := range variables {
		if usage[name] {
			continue
		}

		r.reporter.AddIssue(&reporter.Issue{
			Rule:     r.name,
			Message:  fmt.Sprintf("variable '%s' is declared but never used", name),
			Range:    decl.GetRange(),
			Severity: reporter.SeverityWarning,
		})
	}

	for name, decl := range constants {
		if usage[name] {
			continue
		}

		r.reporter.AddIssue(&reporter.Issue{
			Rule:     r.name,
			Message:  fmt.Sprintf("constant '%s' is declared but never used", name),
			Range:    decl.GetRange(),
			Severity: reporter.SeverityWarning,
		})
	}
}

func (r *UnusedVariables) collectUsage(node ast.ExprNode) map[string]bool {
	usage := make(map[string]bool)

	node.Walk(func(n ast.ExprNode) bool {
		switch node := n.(type) {
		case *ast.Identifier:
			usage[node.Value] = true

		case *ast.AssignmentStatement:
			if node.Left != nil {
				usage[node.Left.Value] = true
			}

		case *ast.ShorthandAssignmentExpr:
			if node.Left == nil {
				return false
			}

			identifier, hasIdentifier := node.Left.(*ast.Identifier)

			if hasIdentifier {
				usage[identifier.Value] = true
			}
		}

		return true
	})

	return usage
}
