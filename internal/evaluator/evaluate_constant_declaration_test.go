package evaluator

import (
	"fmt"
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateConstantDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input *ast.ConstantDeclaration
	}{
		{
			input: &ast.ConstantDeclaration{
				Name:     "x",
				Value:    &ast.NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Type:     datatype.DataTypeNumber.AsString(),
				StartPos: 0,
				EndPos:   1,
			},
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator(io.Discard).evaluateConstantDeclaration(test.input)

		if err != nil {
			t.Fatalf("error evaluating %s: %s", test.input.Expr(), err)
		}
	}
}

func TestEvaluateConstantDeclarationErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.ConstantDeclaration
		expected string
	}{
		{
			input: &ast.ConstantDeclaration{
				Name:     "x",
				Value:    &ast.StringLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Type:     datatype.DataTypeNumber.AsString(),
				StartPos: 0,
				EndPos:   1,
			},
			expected: fmt.Sprintf(
				errorutil.ErrorMsgTypeMismatch,
				datatype.DataTypeNumber.AsString(),
				datatype.DataTypeString.AsString(),
			),
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator(io.Discard).evaluateConstantDeclaration(test.input)

		if err == nil {
			t.Fatalf("expected error evaluating %s, got nil", test.input.Expr())
		}

		if err.Error() != test.expected {
			t.Fatalf("expected \"%s\", got \"%s\"", test.expected, err.Error())
		}
	}
}
