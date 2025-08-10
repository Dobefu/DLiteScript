package evaluator

import (
	"strconv"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateNumberLiteral(
	node *ast.NumberLiteral,
) (datavalue.Value, error) {
	value, err := strconv.ParseFloat(node.Value, 64)

	if err != nil {
		return datavalue.Null(), err
	}

	return datavalue.Number(value), nil
}
