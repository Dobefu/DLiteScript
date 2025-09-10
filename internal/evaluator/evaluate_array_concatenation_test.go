package evaluator

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func TestEvaluateArrayConcatenation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		inputLeft  []datavalue.Value
		inputRight []datavalue.Value
		expected   datavalue.Value
	}{
		{
			name: "array concatenation",
			inputLeft: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
			},
			inputRight: []datavalue.Value{
				datavalue.Number(3),
				datavalue.Number(4),
			},
			expected: datavalue.Array(
				datavalue.Number(1),
				datavalue.Number(2),
				datavalue.Number(3),
				datavalue.Number(4),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			result, err := ev.evaluateArrayConcatenation(
				test.inputLeft,
				test.inputRight,
				nil,
			)

			if err != nil {
				t.Fatalf("expected no error, got: %s", err.Error())
			}

			if result.Value.DataType() != test.expected.DataType() {
				t.Fatalf(
					"expected type to be \"%s\", got \"%s\"",
					test.expected.DataType().AsString(),
					result.Value.DataType().AsString(),
				)
			}

			if !result.Value.Equals(test.expected) {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.ToString(),
					result.Value.ToString(),
				)
			}
		})
	}
}

func TestEvaluateArrayConcatenationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		inputLeft  []datavalue.Value
		inputRight []datavalue.Value
		expected   string
	}{
		{
			name: "type mismatch",
			inputLeft: []datavalue.Value{
				datavalue.Number(1),
				datavalue.Number(2),
			},
			inputRight: []datavalue.Value{datavalue.String("3")},
			expected:   fmt.Sprintf(errorutil.ErrorMsgTypeMismatch, "number", "string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(io.Discard)

			_, err := ev.evaluateArrayConcatenation(
				test.inputLeft,
				test.inputRight,
				&ast.BinaryExpr{
					Left:  &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					Right: &ast.NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
					Operator: token.Token{
						Atom:      "+",
						TokenType: token.TokenTypeOperationAdd,
						StartPos:  0,
						EndPos:    0,
					},
					StartPos: 0,
					EndPos:   0,
				},
			)

			if err == nil {
				t.Fatalf("expected error, got: nil")
			}

			if errors.Unwrap(err).Error() != test.expected {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected,
					errors.Unwrap(err).Error(),
				)
			}
		})
	}
}
