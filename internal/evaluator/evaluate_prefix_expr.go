package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluatePrefixExpr(
	node *ast.PrefixExpr,
) (float64, error) {
	result, err := e.Evaluate(node.Operand)

	if err != nil {
		return 0, err
	}

	if node.Operator.TokenType == token.TokenTypeOperationSub {
		return -result, nil
	}

	return result, nil
}
