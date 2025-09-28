package rules

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/linter/reporter"
)

// UnreachableCode checks for unreachable code after return statements.
type UnreachableCode struct {
	name        string
	description string
	reporter    *reporter.Reporter
}

// NewUnreachableCode creates a new unreachable code rule.
func NewUnreachableCode(reporter *reporter.Reporter) *UnreachableCode {
	return &UnreachableCode{
		name:        "unreachable-code",
		description: "Detects code that cannot be reached after return statements",
		reporter:    reporter,
	}
}

// Name returns the name of the rule.
func (r *UnreachableCode) Name() string {
	return r.name
}

// Description returns the description of the rule.
func (r *UnreachableCode) Description() string {
	return r.description
}

// Analyze analyzes the AST for unreachable code.
func (r *UnreachableCode) Analyze(node ast.ExprNode) {
	checkedBlocks := make(map[*ast.BlockStatement]bool)

	node.Walk(func(n ast.ExprNode) bool {
		switch node := n.(type) {
		case *ast.BlockStatement:
			if !checkedBlocks[node] {
				r.checkBlockForUnreachableCode(node)
				checkedBlocks[node] = true
			}

		case *ast.FuncDeclarationStatement:
			if node.Body == nil {
				return true
			}

			block, hasBlock := node.Body.(*ast.BlockStatement)

			if hasBlock && !checkedBlocks[block] {
				r.checkBlockForUnreachableCode(block)
				checkedBlocks[block] = true
			}
		}

		return true
	})
}

func (r *UnreachableCode) checkBlockForUnreachableCode(
	block *ast.BlockStatement,
) {
	if block == nil || len(block.Statements) == 0 {
		return
	}

	lastReturnIndex := -1

	for i, stmt := range block.Statements {
		_, isReturn := stmt.(*ast.ReturnStatement)

		if isReturn {
			lastReturnIndex = i
		}
	}

	if lastReturnIndex < 0 || lastReturnIndex >= len(block.Statements)-1 {
		return
	}

	for i := lastReturnIndex + 1; i < len(block.Statements); i++ {
		stmt := block.Statements[i]

		_, isComment := stmt.(*ast.CommentLiteral)
		_, isNewline := stmt.(*ast.NewlineLiteral)

		if isComment || isNewline {
			continue
		}

		r.reporter.AddIssue(
			&reporter.Issue{
				Rule:     r.name,
				Message:  "unreachable code after return statement",
				Range:    stmt.GetRange(),
				Severity: reporter.SeverityWarning,
			},
		)
	}
}
