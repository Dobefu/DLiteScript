package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateAssignmentStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected any
	}{
		{
			name: "assignment statement",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.VariableDeclaration{
						Name:     "x",
						Type:     "number",
						Value:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
						StartPos: 0,
						EndPos:   1,
					},
					&ast.AssignmentStatement{
						Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
						Right:    &ast.NumberLiteral{Value: "1", StartPos: 4, EndPos: 7},
						StartPos: 0,
						EndPos:   7,
					},
				},
				StartPos: 0,
				EndPos:   7,
			},
			expected: 1.0,
		},
		{
			name: "assignment statement in block scope",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.VariableDeclaration{
						Name:     "x",
						Type:     "number",
						Value:    &ast.NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
						StartPos: 0,
						EndPos:   1,
					},
					&ast.BlockStatement{
						Statements: []ast.ExprNode{
							&ast.VariableDeclaration{
								Name:     "x",
								Type:     "number",
								Value:    &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
								StartPos: 0,
								EndPos:   1,
							},
							&ast.AssignmentStatement{
								Left:     &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
								Right:    &ast.NumberLiteral{Value: "42", StartPos: 4, EndPos: 6},
								StartPos: 0,
								EndPos:   6,
							},
						},
						StartPos: 0,
						EndPos:   6,
					},
				},
				StartPos: 0,
				EndPos:   6,
			},
			expected: 42.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			eval := NewEvaluator(io.Discard)

			result, err := eval.Evaluate(test.input)

			if err != nil {
				t.Fatalf("failed to evaluate assignment statement: %s", err.Error())
			}

			if result.Value.DataType() != datatype.DataTypeNumber {
				t.Fatalf("expected number result, got %v", result.Value.DataType())
			}

			num, err := result.Value.AsNumber()

			if err != nil {
				t.Fatalf("expected number result: %s", err.Error())
			}

			if num != test.expected {
				t.Errorf("expected %f, got %f", test.expected, num)
			}
		})
	}
}

func TestEvaluateAssignmentStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    ast.ExprNode
		expected string
	}{
		{
			name: "assignment to undefined variable",
			input: &ast.AssignmentStatement{
				Left:     &ast.Identifier{Value: "undefined_var", StartPos: 0, EndPos: 13},
				Right:    &ast.NumberLiteral{Value: "1", StartPos: 16, EndPos: 17},
				StartPos: 0,
				EndPos:   17,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "undefined_var"),
			),
		},
		{
			name: "assignment to constant",
			input: &ast.StatementList{
				Statements: []ast.ExprNode{
					&ast.ConstantDeclaration{
						Name:     "const_var",
						Type:     "number",
						Value:    &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
						StartPos: 0,
						EndPos:   1,
					},
					&ast.AssignmentStatement{
						Left:     &ast.Identifier{Value: "const_var", StartPos: 0, EndPos: 9},
						Right:    &ast.NumberLiteral{Value: "10", StartPos: 12, EndPos: 14},
						StartPos: 0,
						EndPos:   14,
					},
				},
				StartPos: 0,
				EndPos:   14,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgReassignmentToConstant, "const_var"),
			),
		},
		{
			name: "assignment to constant in block scope",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.ConstantDeclaration{
						Name:     "block_const",
						Type:     "number",
						Value:    &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
						StartPos: 0,
						EndPos:   1,
					},
					&ast.AssignmentStatement{
						Left:     &ast.Identifier{Value: "block_const", StartPos: 0, EndPos: 11},
						Right:    &ast.NumberLiteral{Value: "10", StartPos: 14, EndPos: 16},
						StartPos: 0,
						EndPos:   16,
					},
				},
				StartPos: 0,
				EndPos:   16,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgReassignmentToConstant, "block_const"),
			),
		},
		{
			name: "assignment with right side evaluation error",
			input: &ast.AssignmentStatement{
				Left: &ast.Identifier{Value: "x", StartPos: 0, EndPos: 1},
				Right: &ast.FunctionCall{
					Namespace:    "",
					FunctionName: "undefined_func",
					Arguments:    []ast.ExprNode{},
					StartPos:     4,
					EndPos:       18,
				},
				StartPos: 0,
				EndPos:   18,
			},
			expected: fmt.Sprintf(
				"%s: %s at position 4",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedFunction, "undefined_func"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			eval := NewEvaluator(io.Discard)

			_, err := eval.Evaluate(test.input)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected error \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}
