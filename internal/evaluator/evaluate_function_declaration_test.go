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
				StartPos:   0,
				EndPos:     1,
			},
			ReturnValues:    []string{datatype.DataTypeNumber.AsString()},
			NumReturnValues: 1,
			StartPos:        0,
			EndPos:          1,
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
