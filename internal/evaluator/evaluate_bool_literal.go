package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateBoolLiteral(
	node *ast.BoolLiteral,
) (*controlflow.EvaluationResult, error) {
	return controlflow.NewRegularResult(datavalue.Bool(node.Value == "true")), nil
}
