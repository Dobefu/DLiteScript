package rules

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/linter/reporter"
)

func TestMissingReturn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected []*reporter.Issue
	}{
		{
			name: "no missing return",
			input: &ast.FuncDeclarationStatement{
				Name:            "test",
				Args:            []ast.FuncParameter{},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: []*reporter.Issue{},
		},
		{
			name: "missing return",
			input: &ast.FuncDeclarationStatement{
				Name:            "test",
				Args:            []ast.FuncParameter{},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: []*reporter.Issue{
				{
					Rule:    "missing-return",
					Message: "function \"test\" should return \"number\" but has no return statement",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
					Severity: reporter.SeverityError,
				},
			},
		},
		{
			name: "function with return statement should not trigger missing return",
			input: &ast.FuncDeclarationStatement{
				Name:            "test",
				Args:            []ast.FuncParameter{},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{
									Value: "42",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 2, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 2, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 2, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 2, Line: 0, Column: 0},
				},
			},
			expected: []*reporter.Issue{},
		},
		{
			name: "function with empty return types should format as 'values'",
			input: &ast.FuncDeclarationStatement{
				Name:            "test",
				Args:            []ast.FuncParameter{},
				ReturnValues:    []string{},
				NumReturnValues: 1,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: []*reporter.Issue{
				{
					Rule:    "missing-return",
					Message: "function \"test\" should return \"values\" but has no return statement",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
					Severity: reporter.SeverityError,
				},
			},
		},
		{
			name: "function with multiple return types should format with commas",
			input: &ast.FuncDeclarationStatement{
				Name:            "test",
				Args:            []ast.FuncParameter{},
				ReturnValues:    []string{"number", "string", "boolean"},
				NumReturnValues: 3,
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.NumberLiteral{
							Value: "1",
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 1, Line: 0, Column: 0},
							},
						},
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: []*reporter.Issue{
				{
					Rule:    "missing-return",
					Message: "function \"test\" should return \"number, string, boolean\" but has no return statement",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
					Severity: reporter.SeverityError,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rule := NewMissingReturn(reporter.NewReporter(io.Discard))

			if len(rule.Name()) == 0 {
				t.Fatalf("expected name, got none")
			}

			if len(rule.Description()) == 0 {
				t.Fatalf("expected description, got none")
			}

			rule.Analyze(test.input)

			if len(rule.reporter.GetIssues()) != len(test.expected) {
				t.Fatalf(
					"expected %d issue(s), got %d",
					len(test.expected),
					len(rule.reporter.GetIssues()),
				)
			}
		})
	}
}
