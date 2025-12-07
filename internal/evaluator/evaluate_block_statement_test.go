package evaluator

import (
	"errors"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

type discardWriter struct{}

func (d *discardWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type errWriter struct{}

func (e *errWriter) Write(_ []byte) (n int, err error) {
	return 0, errors.New("write error")
}

func (e *errWriter) Error() string {
	return "write error"
}

func TestEvaluateBlockStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.BlockStatement
		expected datavalue.Value
	}{
		{
			name: "single statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "5",
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
			expected: datavalue.Number(5),
		},
		{
			name: "break statement",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.BreakStatement{
						Count: 1,
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
			expected: datavalue.Null(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(&discardWriter{})
			ev.buf.WriteString("test")
			rawResult, err := ev.evaluateBlockStatement(test.input)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err.Error())
			}

			if rawResult.Value.DataType != test.expected.DataType {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.DataType.AsString(),
					rawResult.Value.DataType.AsString(),
				)
			}
		})
	}
}

func TestEvaluateBlockStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.BlockStatement
		expected string
	}{
		{
			name: "write error",
			input: &ast.BlockStatement{
				Statements: []ast.ExprNode{
					&ast.NumberLiteral{
						Value: "5",
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
			expected: "write error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ev := NewEvaluator(&errWriter{})
			ev.buf.WriteString("test")
			_, err := ev.evaluateBlockStatement(test.input)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if err.Error() != test.expected {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err.Error())
			}
		})
	}
}
