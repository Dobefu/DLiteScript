package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateAnyLiteral(
	node *ast.AnyLiteral,
) (*controlflow.EvaluationResult, error) {
	return controlflow.NewRegularResult(datavalue.Any(node.Value)), nil
}
