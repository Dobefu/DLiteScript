package evaluator

import (
	"io"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
)

func TestEvaluateFunctionDeclaration(t *testing.T) {
	t.Parallel()

	evaluator := NewEvaluator(io.Discard)

	_, err := evaluator.evaluateFunctionDeclaration(
		&ast.FuncDeclarationStatement{
			Name: "abs",
			Args: []ast.FuncParameter{
				{
					Name: "num",
					Type: datatype.DataTypeNumber.AsString(),
				},
			},
			Body: &ast.BlockStatement{
				Statements: []ast.ExprNode{},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 1, Line: 0, Column: 0},
				},
			},
			ReturnValues:    []string{datatype.DataTypeNumber.AsString()},
			NumReturnValues: 1,
			Range: ast.Range{
				Start: ast.Position{Offset: 0, Line: 0, Column: 0},
				End:   ast.Position{Offset: 1, Line: 0, Column: 0},
			},
		},
	)

	if err != nil {
		t.Fatalf("expected no error, got: \"%s\"", err.Error())
	}

	_, hasFunction := evaluator.userFunctions["abs"]

	if !hasFunction {
		t.Fatalf("expected the evaluator to have custom function \"abs\", got none")
	}
}
