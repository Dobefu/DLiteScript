package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
)

func (e *Evaluator) evaluateContinueStatement(
	node *ast.ContinueStatement,
) (*controlflow.EvaluationResult, error) {
	return controlflow.NewContinueResult(node.Count), nil
}
