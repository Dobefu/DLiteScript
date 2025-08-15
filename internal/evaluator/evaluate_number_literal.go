package evaluator

import (
	"strconv"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateNumberLiteral(
	node *ast.NumberLiteral,
) (*controlflow.EvaluationResult, error) {
	value, err := strconv.ParseFloat(node.Value, 64)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	return controlflow.NewRegularResult(datavalue.Number(value)), nil
}
