package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateFunctionDeclaration(
	node *ast.FuncDeclarationStatement,
) (*controlflow.EvaluationResult, error) {
	e.userFunctions[node.Name] = node

	return controlflow.NewRegularResult(datavalue.Function(node)), nil
}
