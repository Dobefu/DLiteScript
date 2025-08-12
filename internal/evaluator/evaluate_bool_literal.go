package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateBoolLiteral(
	node *ast.BoolLiteral,
) (datavalue.Value, error) {
	return datavalue.Bool(node.Value == "true"), nil
}
