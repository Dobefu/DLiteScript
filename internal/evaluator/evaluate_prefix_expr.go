package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluatePrefixExpr(
	node *ast.PrefixExpr,
) (datavalue.Value, error) {
	rawResult, err := e.Evaluate(node.Operand)

	if err != nil {
		return datavalue.Null(), err
	}

	if node.Operator.TokenType == token.TokenTypeOperationSub {
		number, err := rawResult.AsNumber()

		if err != nil {
			return datavalue.Null(), err
		}

		return datavalue.Number(-number), nil
	}

	if node.Operator.TokenType == token.TokenTypeOperationAdd {
		number, err := rawResult.AsNumber()

		if err != nil {
			return datavalue.Null(), err
		}

		return datavalue.Number(number), nil
	}

	if node.Operator.TokenType == token.TokenTypeNot {
		boolean, err := rawResult.AsBool()

		if err != nil {
			return datavalue.Null(), err
		}

		return datavalue.Bool(!boolean), nil
	}

	return rawResult, nil
}
