package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluateLogicalBinaryExpr(t *testing.T) {
	t.Parallel()

	evaluator := NewEvaluator(io.Discard)

	tests := []struct {
		name       string
		inputLeft  datavalue.Value
		inputRight datavalue.Value
		inputNode  *ast.BinaryExpr
		expected   datavalue.Value
	}{
		{
			name:       "true && false",
			inputLeft:  datavalue.Bool(true),
			inputRight: datavalue.Bool(false),
			inputNode: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.BoolLiteral{
					Value: "false",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "&&",
					TokenType: token.TokenTypeLogicalAnd,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Bool(false),
		},
		{
			name:       "true || false",
			inputLeft:  datavalue.Bool(true),
			inputRight: datavalue.Bool(false),
			inputNode: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.BoolLiteral{
					Value: "false",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "||",
					TokenType: token.TokenTypeLogicalOr,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Bool(true),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			rawResult, err := evaluator.evaluateLogicalBinaryExpr(
				test.inputLeft,
				test.inputRight,
				test.inputNode,
			)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.inputNode.Expr(), err)
			}

			if rawResult.Value.DataType != test.expected.DataType {
				t.Fatalf(
					"expected %s, got %s",
					test.expected.DataType.AsString(),
					rawResult.Value.DataType.AsString(),
				)
			}
		})
	}
}

func TestEvaluateLogicalBinaryExprErr(t *testing.T) {
	t.Parallel()

	evaluator := NewEvaluator(io.Discard)

	tests := []struct {
		name       string
		inputLeft  datavalue.Value
		inputRight datavalue.Value
		inputNode  *ast.BinaryExpr
		expected   string
	}{
		{
			name:       "number and string",
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.String("5"),
			inputNode: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      ">",
					TokenType: token.TokenTypeGreaterThan,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"could not get binary expr value as bool: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "number"),
			),
		},
		{
			name:       "number and number",
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.Number(5),
			inputNode: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "==",
					TokenType: token.TokenTypeEqual,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"could not get binary expr value as bool: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "number"),
			),
		},
		{
			name:       "unexpected operator",
			inputLeft:  datavalue.Bool(true),
			inputRight: datavalue.Bool(true),
			inputNode: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 2, Line: 0, Column: 0},
						End:   ast.Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Operator: token.Token{
					Atom:      "**",
					TokenType: token.TokenTypeOperationPow,
					StartPos:  0,
					EndPos:    0,
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnknownOperator, "**"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := evaluator.evaluateLogicalBinaryExpr(
				test.inputLeft,
				test.inputRight,
				test.inputNode,
			)

			if err == nil {
				t.Fatalf(
					"expected error evaluating \"%s\", got nil",
					test.inputNode.Expr(),
				)
			}

			if err.Error() != test.expected {
				t.Errorf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
