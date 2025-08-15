package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
)

func (e *Evaluator) evaluateBreakStatement(
	node *ast.BreakStatement,
) (*controlflow.EvaluationResult, error) {
	return controlflow.NewBreakResult(node.Count), nil
}
