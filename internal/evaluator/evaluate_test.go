package evaluator

import (
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

type unknownNode struct {
	StartPos int
	EndPos   int
}

func (n *unknownNode) GetRange() ast.Range {
	return ast.Range{
		Start: ast.Position{Offset: n.StartPos, Line: 0, Column: 0},
		End:   ast.Position{Offset: n.EndPos, Line: 0, Column: 0},
	}
}

func (n *unknownNode) Expr() string {
	return "unknown"
}

func (n *unknownNode) Walk(fn func(_ ast.ExprNode) bool) {
	fn(n)
}

func TestEvaluate(t *testing.T) {
	t.Parallel()

	funcDeclaration := &ast.FuncDeclarationStatement{
		Name: "someFunction",
		Args: []ast.FuncParameter{},
		Body: &ast.BlockStatement{
			Statements: []ast.ExprNode{},
			Range: ast.Range{
				Start: ast.Position{Offset: 0, Line: 0, Column: 0},
				End:   ast.Position{Offset: 1, Line: 0, Column: 0},
			},
		},
		ReturnValues:    []string{},
		NumReturnValues: 0,
		Range: ast.Range{
			Start: ast.Position{Offset: 0, Line: 0, Column: 0},
			End:   ast.Position{Offset: 1, Line: 0, Column: 0},
		},
	}

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected *controlflow.EvaluationResult
	}{
		{
			name: "binary expression",
			input: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.BinaryExpr{
					Left: &ast.NumberLiteral{
						Value: "5",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
					Right: &ast.PrefixExpr{
						Operator: token.Token{
							Atom:      "-",
							TokenType: token.TokenTypeOperationSub,
							StartPos:  0,
							EndPos:    0,
						},
						Operand: &ast.FunctionCall{
							Namespace:    "math",
							FunctionName: "abs",
							Arguments: []ast.ExprNode{
								&ast.Identifier{
									Value: "PI",
									Range: ast.Range{
										Start: ast.Position{Offset: 4, Line: 0, Column: 0},
										End:   ast.Position{Offset: 5, Line: 0, Column: 0},
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
					Operator: token.Token{
						Atom:      "+",
						TokenType: token.TokenTypeOperationAdd,
						StartPos:  0,
						EndPos:    0,
					},
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "+",
					TokenType: token.TokenTypeOperationAdd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(
				datavalue.Number(5 + math.Abs(-5+math.Pi)),
			),
		},
		{
			name: "index assignment statement",
			input: &ast.IndexAssignmentStatement{
				Array: &ast.Identifier{
					Value: "someArray",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Index: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "1",
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
			expected: controlflow.NewRegularResult(
				datavalue.Array(datavalue.Number(1)),
			),
		},
		{
			name:  "function declaration statement",
			input: funcDeclaration,
			expected: controlflow.NewRegularResult(
				datavalue.Function(funcDeclaration),
			),
		},
		{
			name: "comment literal",
			input: &ast.CommentLiteral{
				Value: "test",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Null()),
		},
		{
			name: "spread expression",
			input: &ast.SpreadExpr{
				Expression: &ast.NumberLiteral{
					Value: "1",
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
			expected: controlflow.NewRegularResult(datavalue.Number(1)),
		},
		{
			name: "index expression",
			input: &ast.IndexExpr{
				Array: &ast.Identifier{
					Value: "someArray",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Index: &ast.NumberLiteral{
					Value: "0",
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
			expected: controlflow.NewRegularResult(datavalue.Number(0)),
		},
		{
			name: "import statement",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "../../examples/09_imports/test.dl",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Null()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			ev.outerScope["someArray"] = &Variable{
				Value: datavalue.Array(datavalue.Number(0)),
				Type:  "array",
			}

			result, err := ev.Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating \"%s\": %s", test.input.Expr(), err.Error())
			}

			if result.Value.DataType != test.expected.Value.DataType {
				t.Errorf(
					"expected \"%T\", got \"%T\"",
					test.expected.Value.DataType,
					result.Value.DataType,
				)
			}

			if !result.Value.Equals(test.expected.Value) {
				t.Errorf(
					"expected \"%s\", got \"%s\"",
					test.expected.Value.ToString(),
					result.Value.ToString(),
				)
			}
		})
	}
}

func TestEvaluateErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name:  "unknown node type",
			input: &unknownNode{StartPos: 0, EndPos: 1},
			expected: fmt.Sprintf(
				"%s: %s line 1 at position 1",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgUnknownNodeType,
					&unknownNode{}, //nolint:exhaustruct
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).Evaluate(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestOutput(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "function call",
			input: &ast.FunctionCall{
				Namespace:    "",
				FunctionName: "printf",
				Arguments: []ast.ExprNode{
					&ast.StringLiteral{
						Value: "test",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
			expected: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			evaluator := NewEvaluator(io.Discard)

			_, err := evaluator.Evaluate(test.input)

			if err != nil {
				t.Fatalf("error evaluating \"%s\": %s", test.input.Expr(), err.Error())
			}

			if evaluator.Output() != test.expected {
				t.Errorf("expected \"%v\", got \"%v\"", test.expected, evaluator.Output())
			}
		})
	}
}

func BenchmarkEvaluate(b *testing.B) {
	for b.Loop() {
		_, _ = NewEvaluator(io.Discard).Evaluate(
			&ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.BinaryExpr{
					Left: &ast.PrefixExpr{
						Operator: *token.NewToken("-", token.TokenTypeOperationSub, 0, 0),
						Operand: &ast.NumberLiteral{
							Value: "2",
							Range: ast.Range{
								Start: ast.Position{Offset: 2, Line: 0, Column: 0},
								End:   ast.Position{Offset: 3, Line: 0, Column: 0},
							},
						},
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					Right: &ast.BinaryExpr{
						Left: &ast.NumberLiteral{
							Value: "3",
							Range: ast.Range{
								Start: ast.Position{Offset: 4, Line: 0, Column: 0},
								End:   ast.Position{Offset: 5, Line: 0, Column: 0},
							},
						},
						Right: &ast.NumberLiteral{
							Value: "4",
							Range: ast.Range{
								Start: ast.Position{Offset: 6, Line: 0, Column: 0},
								End:   ast.Position{Offset: 7, Line: 0, Column: 0},
							},
						},
						Operator: *token.NewToken("/", token.TokenTypeOperationDiv, 0, 0),
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 0, Line: 0, Column: 0},
						},
					},
					Operator: *token.NewToken("*", token.TokenTypeOperationMul, 0, 0),
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 0, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("+", token.TokenTypeOperationAdd, 0, 0),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 0, Line: 0, Column: 0},
				},
			},
		)
	}
}
