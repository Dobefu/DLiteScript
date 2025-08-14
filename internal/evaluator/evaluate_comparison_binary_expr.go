package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateComparisonBinaryExpr(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (datavalue.Value, error) {
	leftNumber, rightNumber, err := e.getBinaryExprValueAsNumber(leftValue, rightValue)

	if err != nil {
		return datavalue.Null(), err
	}

	switch node.Operator.TokenType {
	case token.TokenTypeGreaterThan:
		return datavalue.Bool(leftNumber > rightNumber), nil

	case token.TokenTypeGreaterThanOrEqual:
		return datavalue.Bool(leftNumber >= rightNumber), nil

	case token.TokenTypeLessThan:
		return datavalue.Bool(leftNumber < rightNumber), nil

	case token.TokenTypeLessThanOrEqual:
		return datavalue.Bool(leftNumber <= rightNumber), nil

	default:
		return datavalue.Null(), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.Position(),
			node.Operator.Atom,
		)
	}
}
