package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluateEqualityBinaryExpr(t *testing.T) {
	t.Parallel()

	function := &ast.FuncDeclarationStatement{
		Name: "printf",
		Args: []ast.FuncParameter{},
		Body: &ast.StringLiteral{
			Value: "1",
			Range: ast.Range{
				Start: ast.Position{Offset: 0, Line: 0, Column: 0},
				End:   ast.Position{Offset: 3, Line: 0, Column: 0},
			},
		},
		ReturnValues: []string{
			"string",
		},
		NumReturnValues: 1,
		Range: ast.Range{
			Start: ast.Position{Offset: 0, Line: 0, Column: 0},
			End:   ast.Position{Offset: 18, Line: 0, Column: 0},
		},
	}

	tests := []struct {
		name       string
		inputLeft  datavalue.Value
		inputRight datavalue.Value
		inputNode  *ast.BinaryExpr
		expected   datavalue.Value
	}{
		{
			name:       "number",
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "string",
			inputLeft:  datavalue.String("5"),
			inputRight: datavalue.String("5"),
			inputNode: &ast.BinaryExpr{
				Left: &ast.StringLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.StringLiteral{
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "boolean",
			inputLeft:  datavalue.Bool(true),
			inputRight: datavalue.Bool(true),
			inputNode: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.BoolLiteral{
					Value: "true",
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "any",
			inputLeft:  datavalue.Any(1),
			inputRight: datavalue.Any(1),
			inputNode: &ast.BinaryExpr{
				Left: &ast.AnyLiteral{
					Value: &ast.NumberLiteral{
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
				Right: &ast.AnyLiteral{
					Value: &ast.NumberLiteral{
						Value: "1",
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "null",
			inputLeft:  datavalue.Null(),
			inputRight: datavalue.Null(),
			inputNode: &ast.BinaryExpr{
				Left: &ast.NullLiteral{
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NullLiteral{
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "function",
			inputLeft:  datavalue.Function(function),
			inputRight: datavalue.Function(function),
			inputNode: &ast.BinaryExpr{
				Left:  function,
				Right: function,
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "tuple",
			inputLeft:  datavalue.Tuple(datavalue.Number(1), datavalue.Number(2)),
			inputRight: datavalue.Tuple(datavalue.Number(1), datavalue.Number(2)),
			inputNode: &ast.BinaryExpr{
				Left:  nil,
				Right: nil,
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "array",
			inputLeft:  datavalue.Array(datavalue.Number(1), datavalue.Number(2)),
			inputRight: datavalue.Array(datavalue.Number(1), datavalue.Number(2)),
			inputNode: &ast.BinaryExpr{
				Left:  nil,
				Right: nil,
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
			expected: datavalue.Bool(true),
		},
		{
			name:       "different data types",
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
				Right: &ast.StringLiteral{
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
			expected: datavalue.Bool(false),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := NewEvaluator(io.Discard).evaluateEqualityBinaryExpr(
				test.inputLeft,
				test.inputRight,
				test.inputNode,
			)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.inputNode.Expr(), err)
			}

			if !result.Value.Equals(test.expected) {
				t.Fatalf(
					"expected \"%v\", got \"%v\"",
					test.expected,
					result.Value,
				)
			}
		})
	}
}

func TestEvaluateEqualityBinaryExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		inputLeft  datavalue.Value
		inputRight datavalue.Value
		inputNode  *ast.BinaryExpr
		expected   string
	}{
		{
			name:       "right value is not a number",
			inputLeft:  datavalue.Number(5),
			inputRight: datavalue.Any(nil),
			inputNode: &ast.BinaryExpr{
				Left: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.AnyLiteral{
					Value: &ast.NullLiteral{
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
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
				"could not get binary expr value as number: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "any"),
			),
		},
		{
			name:       "right value is not a string",
			inputLeft:  datavalue.String("5"),
			inputRight: datavalue.Any(nil),
			inputNode: &ast.BinaryExpr{
				Left: &ast.StringLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.AnyLiteral{
					Value: &ast.NullLiteral{
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
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
				"could not get binary expr value as string: %s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "string", "any"),
			),
		},
		{
			name:       "right value is not a boolean",
			inputLeft:  datavalue.Bool(true),
			inputRight: datavalue.Any(nil),
			inputNode: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Value: "true",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.AnyLiteral{
					Value: &ast.NullLiteral{
						Range: ast.Range{
							Start: ast.Position{Offset: 2, Line: 0, Column: 0},
							End:   ast.Position{Offset: 3, Line: 0, Column: 0},
						},
					},
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
				fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "bool", "any"),
			),
		},
		{
			name: "unknown left data type",
			inputLeft: datavalue.Value{ //nolint:exhaustruct
				DataType: datatype.DataType(-1),
			},
			inputRight: datavalue.Value{ //nolint:exhaustruct
				DataType: datatype.DataType(-1),
			},
			inputNode: &ast.BinaryExpr{
				Left: &ast.AnyLiteral{
					Value: &ast.NullLiteral{
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
				Right: &ast.BoolLiteral{
					Value: "true",
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
				"%s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgTypeUnknownDataType, "unknown"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewEvaluator(io.Discard).evaluateEqualityBinaryExpr(
				test.inputLeft,
				test.inputRight,
				test.inputNode,
			)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
