package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateArrayLiteral(
	node *ast.ArrayLiteral,
) (*controlflow.EvaluationResult, error) {
	var values []datavalue.Value

	for _, value := range node.Values {
		evaluatedValue, err := e.Evaluate(value)

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		values = append(values, evaluatedValue.Value)
	}

	return controlflow.NewRegularResult(datavalue.Array(values...)), nil
}
