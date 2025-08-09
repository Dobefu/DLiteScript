package evaluator

import (
	"strconv"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (e *Evaluator) evaluateNumberLiteral(
	node *ast.NumberLiteral,
) (float64, error) {
	return strconv.ParseFloat(node.Value, 64)
}
