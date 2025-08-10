package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateStringLiteral(
	node *ast.StringLiteral,
) (datavalue.Value, error) {
	return datavalue.String(node.Value), nil
}
