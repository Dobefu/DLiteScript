package rules

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/linter/reporter"
)

func TestUnreachableCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected []*reporter.Issue
	}{
		{
			name: "no unreachable code",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: []*reporter.Issue{},
		},
		{
			name: "unreachable code in block statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.ReturnStatement{
						Values: []ast.ExprNode{
							&ast.NumberLiteral{
								Value: "1",
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 0, Line: 0, Column: 0},
								},
							},
						},
						NumValues: 1,
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
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
			expected: []*reporter.Issue{
				{
					Rule:    "unreachable-code",
					Message: "unreachable code",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
					Severity: reporter.SeverityWarning,
				},
			},
		},
		{
			name: "unreachable code in function declaration statement",
			input: &ast.FuncDeclarationStatement{
				Name:            "test",
				Args:            []ast.FuncParameter{},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				Body: &ast.BlockStatement{
					Statements: []ast.ExprNode{
						&ast.ReturnStatement{
							Values: []ast.ExprNode{
								&ast.NumberLiteral{
									Value: "1",
									Range: ast.Range{
										Start: ast.Position{Offset: 0, Line: 0, Column: 0},
										End:   ast.Position{Offset: 0, Line: 0, Column: 0},
									},
								},
							},
							NumValues: 1,
							Range: ast.Range{
								Start: ast.Position{Offset: 0, Line: 0, Column: 0},
								End:   ast.Position{Offset: 0, Line: 0, Column: 0},
							},
						},
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
			},
			expected: []*reporter.Issue{
				{
					Rule:    "unreachable-code",
					Message: "unreachable code",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
					Severity: reporter.SeverityWarning,
				},
			},
		},
		{
			name: "function declaration with nil body",
			input: &ast.FuncDeclarationStatement{
				Name:            "test",
				Args:            []ast.FuncParameter{},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
				Body: nil,
			},
			expected: []*reporter.Issue{},
		},
		{
			name: "unreachable code with comments and newlines after return",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.ReturnStatement{
						Values: []ast.ExprNode{
							&ast.NumberLiteral{
								Value: "1",
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 0, Line: 0, Column: 0},
								},
							},
						},
						NumValues: 1,
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					&ast.CommentLiteral{
						Value: "// This is a comment",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					&ast.NewlineLiteral{
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					&ast.NumberLiteral{
						Value: "2",
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
			expected: []*reporter.Issue{
				{
					Rule:    "unreachable-code",
					Message: "unreachable code after return statement",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
					Severity: reporter.SeverityWarning,
				},
			},
		},
		{
			name: "block with return as last statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					&ast.ReturnStatement{
						Values: []ast.ExprNode{
							&ast.NumberLiteral{
								Value: "2",
								Range: ast.Range{
									Start: ast.Position{Offset: 0, Line: 0, Column: 0},
									End:   ast.Position{Offset: 0, Line: 0, Column: 0},
								},
							},
						},
						NumValues: 1,
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
			expected: []*reporter.Issue{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rule := NewUnreachableCode(reporter.NewReporter(io.Discard))

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
