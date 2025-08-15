package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateStringLiteral(
	node *ast.StringLiteral,
) (*controlflow.EvaluationResult, error) {
	return controlflow.NewRegularResult(datavalue.String(node.Value)), nil
}
