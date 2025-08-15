package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateNullLiteral() (*controlflow.EvaluationResult, error) {
	return controlflow.NewRegularResult(datavalue.Null()), nil
}
