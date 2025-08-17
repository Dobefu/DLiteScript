package evaluator

import (
	"errors"
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateIdentifier(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.Identifier
		expected float64
	}{
		{
			input:    &ast.Identifier{Value: "PI", StartPos: 0, EndPos: 1},
			expected: math.Pi,
		},
		{
			input:    &ast.Identifier{Value: "TAU", StartPos: 0, EndPos: 1},
			expected: math.Pi * 2,
		},
		{
			input:    &ast.Identifier{Value: "E", StartPos: 0, EndPos: 1},
			expected: math.E,
		},
		{
			input:    &ast.Identifier{Value: "PHI", StartPos: 0, EndPos: 1},
			expected: math.Phi,
		},
		{
			input:    &ast.Identifier{Value: "LN2", StartPos: 0, EndPos: 1},
			expected: math.Ln2,
		},
		{
			input:    &ast.Identifier{Value: "LN10", StartPos: 0, EndPos: 1},
			expected: math.Ln10,
		},
	}

	for _, test := range tests {
		rawResult, err := NewEvaluator(io.Discard).evaluateIdentifier(test.input)

		if err != nil {
			t.Errorf("error evaluating '%s': %s", test.input.Expr(), err.Error())
		}

		result, err := rawResult.Value.AsNumber()

		if err != nil {
			t.Fatalf("expected number, got type error: %s", err.Error())
		}

		if result != test.expected {
			t.Errorf("expected %f, got %f", test.expected, result)
		}
	}
}

func TestEvaluateIdentifierErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    *ast.Identifier
		expected string
	}{
		{
			input:    &ast.Identifier{Value: "bogus", StartPos: 0, EndPos: 1},
			expected: fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "bogus"),
		},
	}

	for _, test := range tests {
		_, err := NewEvaluator(io.Discard).evaluateIdentifier(test.input)

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
	}
}
