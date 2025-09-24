package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluateShorthandAssignmentExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected *controlflow.EvaluationResult
	}{
		{
			name: "addition",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
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
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Number(11)),
		},
		{
			name: "subtraction",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
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
				Operator: *token.NewToken("-=", token.TokenTypeOperationSubAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Number(9)),
		},
		{
			name: "multiplication",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "2",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("*=", token.TokenTypeOperationMulAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Number(20)),
		},
		{
			name: "division",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "2",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("/=", token.TokenTypeOperationDivAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Number(5)),
		},
		{
			name: "modulo",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "2",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("%=", token.TokenTypeOperationModAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Number(0)),
		},
		{
			name: "power",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "2",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("**=", token.TokenTypeOperationPowAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Number(100)),
		},
		{
			name: "array index addition",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.IndexExpr{
					Array: &ast.Identifier{
						Value: "arr",
						Range: ast.Range{
							Start: ast.Position{Offset: 0, Line: 0, Column: 0},
							End:   ast.Position{Offset: 1, Line: 0, Column: 0},
						},
					},
					Index: &ast.NumberLiteral{
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
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Array(
				datavalue.Number(0),
				datavalue.Number(6),
			)),
		},
		{
			name: "unsupported left operand type",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.NumberLiteral{
					Value: "10",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "5",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Null()),
		},
		{
			name: "unknown operator type",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "3",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("??=", token.Type(999), 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: controlflow.NewRegularResult(datavalue.Number(13)),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			ev.outerScope["x"] = &Variable{
				Value: datavalue.Number(10),
				Type:  "number",
			}

			ev.outerScope["arr"] = &Variable{
				Value: datavalue.Array(datavalue.Number(0), datavalue.Number(1)),
				Type:  "array",
			}

			result, err := ev.Evaluate(test.input)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
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

func TestEvaluateShorthandAssignmentExprErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "evaluation error on right-hand side",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "bogus",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAdd, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "invalid syntax",
		},
		{
			name: "evaluation error on left-hand side",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
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
				Operator: *token.NewToken("+=", token.TokenTypeOperationAdd, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "x"),
		},
		{
			name: "undefined identifier for right-hand side",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.Identifier{
					Value: "bogus",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"),
		},
		{
			name: "arithmetic binary expression error",
			input: &ast.ShorthandAssignmentExpr{
				Left: &ast.NumberLiteral{
					Value: "10",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Right: &ast.NumberLiteral{
					Value: "0",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Operator: *token.NewToken("/=", token.TokenTypeOperationDivAssign, 0, 1),
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			expected: "division by zero",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			_, err := ev.Evaluate(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Errorf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}

func TestAssignArrayIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.IndexExpr
		expected *controlflow.EvaluationResult
	}{
		{
			name: "assignment to array variable index",
			input: &ast.IndexExpr{
				Array: &ast.Identifier{
					Value: "x",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Index: &ast.NumberLiteral{
					Value: "2",
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
				datavalue.Array(
					datavalue.Number(0),
					datavalue.Number(1),
					datavalue.Number(2),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			ev.outerScope["x"] = &Variable{
				Value: datavalue.Array(
					datavalue.Number(0),
					datavalue.Number(1),
					datavalue.Number(2),
				),
				Type: "array",
			}

			result, err := ev.assignArrayIndex(test.input, datavalue.Number(2))

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
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

func TestAssignArrayIndexErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.IndexExpr
		expected string
	}{
		{
			name: "evaluation error on array expression",
			input: &ast.IndexExpr{
				Array: &ast.Identifier{
					Value: "bogus",
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
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"),
		},
		{
			name: "evaluation error on index expression",
			input: &ast.IndexExpr{
				Array: &ast.Identifier{
					Value: "arr",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Index: &ast.Identifier{
					Value: "bogus",
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
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"),
		},
		{
			name: "evaluation error on array",
			input: &ast.IndexExpr{
				Array: &ast.NumberLiteral{
					Value: "10",
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
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "array", "number"),
		},
		{
			name: "index value cannot be converted to number",
			input: &ast.IndexExpr{
				Array: &ast.Identifier{
					Value: "arr",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Index: &ast.StringLiteral{
					Value: "nan",
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
			expected: fmt.Sprintf(errorutil.ErrorMsgTypeExpected, "number", "string"),
		},
		{
			name: "index out of bounds",
			input: &ast.IndexExpr{
				Array: &ast.Identifier{
					Value: "arr",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				Index: &ast.NumberLiteral{
					Value: "-1",
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
			expected: fmt.Sprintf(errorutil.ErrorMsgArrayIndexOutOfBounds, "-1"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)
			ev.outerScope["arr"] = &Variable{
				Value: datavalue.Array(datavalue.Number(0), datavalue.Number(1)),
				Type:  "array",
			}

			_, err := ev.assignArrayIndex(test.input, datavalue.Number(2))

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Errorf(
					"expected error \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}

func TestAssignArrayIndexNonIdentifier(t *testing.T) {
	t.Parallel()

	expr := &ast.IndexExpr{
		Array: &ast.ArrayLiteral{
			Values: []ast.ExprNode{
				&ast.NumberLiteral{
					Value: "1",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				&ast.NumberLiteral{
					Value: "2",
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
	}

	ev := NewEvaluator(io.Discard)
	result, err := ev.assignArrayIndex(expr, datavalue.Number(99))

	if err != nil {
		t.Fatalf("expected no error, got \"%s\"", err.Error())
	}

	expected := controlflow.NewRegularResult(datavalue.Number(99))

	if !result.Value.Equals(expected.Value) {
		t.Errorf(
			"expected \"%s\", got \"%s\"",
			expected.Value.ToString(),
			result.Value.ToString(),
		)
	}
}
