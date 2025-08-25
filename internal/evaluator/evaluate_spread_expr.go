package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
)

func (e *Evaluator) evaluateSpreadExpr(
	node *ast.SpreadExpr,
) (*controlflow.EvaluationResult, error) {
	return e.Evaluate(node.Expression)
}
