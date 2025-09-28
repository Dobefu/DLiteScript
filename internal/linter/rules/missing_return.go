package rules

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/linter/reporter"
)

// MissingReturn checks for functions that don't return values when they should.
type MissingReturn struct {
	name        string
	description string
	reporter    *reporter.Reporter
}

// NewMissingReturn creates a new missing return rule.
func NewMissingReturn(reporter *reporter.Reporter) *MissingReturn {
	return &MissingReturn{
		name:        "missing-return",
		description: "Detects functions that should return values but don't have return statements",
		reporter:    reporter,
	}
}

// Name returns the name of the rule.
func (r *MissingReturn) Name() string {
	return r.name
}

// Description returns the description of the rule.
func (r *MissingReturn) Description() string {
	return r.description
}

// Analyze analyzes the AST for missing return statements.
func (r *MissingReturn) Analyze(node ast.ExprNode) {
	checkedFunctions := make(map[*ast.FuncDeclarationStatement]bool)

	node.Walk(func(n ast.ExprNode) bool {
		funcDecl, hasFuncDecl := n.(*ast.FuncDeclarationStatement)

		if hasFuncDecl && !checkedFunctions[funcDecl] {
			r.checkFunctionForMissingReturn(funcDecl)
			checkedFunctions[funcDecl] = true
		}

		return true
	})
}

func (r *MissingReturn) checkFunctionForMissingReturn(
	funcDecl *ast.FuncDeclarationStatement,
) {
	if funcDecl.NumReturnValues == 0 {
		return
	}

	hasReturn := false

	funcDecl.Body.Walk(func(n ast.ExprNode) bool {
		if _, isReturn := n.(*ast.ReturnStatement); isReturn {
			hasReturn = true

			return false
		}

		return true
	})

	if !hasReturn {
		r.reporter.AddIssue(
			&reporter.Issue{
				Rule: r.name,
				Message: fmt.Sprintf(
					"function \"%s\" should return \"%s\" but has no return statement",
					funcDecl.Name,
					r.formatReturnTypes(funcDecl.ReturnValues),
				),
				Range:    funcDecl.GetRange(),
				Severity: reporter.SeverityError,
			},
		)
	}
}

func (r *MissingReturn) formatReturnTypes(returnTypes []string) string {
	if len(returnTypes) == 0 {
		return "values"
	}

	if len(returnTypes) == 1 {
		return returnTypes[0]
	}

	result := ""

	for i, returnType := range returnTypes {
		if i > 0 {
			result += ", "
		}

		result += returnType
	}

	return result
}
