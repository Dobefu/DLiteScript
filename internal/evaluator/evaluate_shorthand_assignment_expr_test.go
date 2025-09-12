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
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(11)),
		},
		{
			name: "subtraction",
			input: &ast.ShorthandAssignmentExpr{
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("-=", token.TokenTypeOperationSubAssign, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(9)),
		},
		{
			name: "multiplication",
			input: &ast.ShorthandAssignmentExpr{
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("*=", token.TokenTypeOperationMulAssign, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(20)),
		},
		{
			name: "division",
			input: &ast.ShorthandAssignmentExpr{
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("/=", token.TokenTypeOperationDivAssign, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(5)),
		},
		{
			name: "modulo",
			input: &ast.ShorthandAssignmentExpr{
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("%=", token.TokenTypeOperationModAssign, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(0)),
		},
		{
			name: "power",
			input: &ast.ShorthandAssignmentExpr{
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("**=", token.TokenTypeOperationPowAssign, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: controlflow.NewRegularResult(datavalue.Number(100)),
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
			name: "evaluation error",
			input: &ast.ShorthandAssignmentExpr{
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.NumberLiteral{Value: "bogus", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAdd, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: "invalid syntax",
		},
		{
			name: "undefined identifier for right",
			input: &ast.ShorthandAssignmentExpr{
				Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right:    &ast.Identifier{Value: "bogus", StartPos: 0, EndPos: 1},
				Operator: *token.NewToken("+=", token.TokenTypeOperationAddAssign, 0, 1),
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"),
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
				Array:    &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Index:    &ast.NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   1,
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
				Value: datavalue.Array(datavalue.Number(0), datavalue.Number(1), datavalue.Number(2)),
				Type:  "array",
			}

			result, err := ev.assignArrayIndex(test.input, datavalue.Number(2))

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if !result.Value.Equals(test.expected.Value) {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected.Value.ToString(), result.Value.ToString())
			}
		})
	}
}
