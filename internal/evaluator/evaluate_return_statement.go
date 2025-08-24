package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateReturnStatement(
	node *ast.ReturnStatement,
) (*controlflow.EvaluationResult, error) {
	if node.NumValues == 0 {
		return controlflow.NewReturnResult(datavalue.Null()), nil
	}

	if len(node.Values) == 1 {
		value, err := e.Evaluate(node.Values[0])

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		return controlflow.NewReturnResult(value.Value), nil
	}

	values := make([]datavalue.Value, len(node.Values))

	for i, val := range node.Values {
		result, err := e.Evaluate(val)

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		values[i] = result.Value
	}

	if len(values) == 1 {
		return controlflow.NewReturnResult(values[0]), nil
	}

	return controlflow.NewReturnResult(datavalue.Null()), nil
}
