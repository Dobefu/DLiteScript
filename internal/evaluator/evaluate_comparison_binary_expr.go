package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateComparisonBinaryExpr(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	leftNumber, rightNumber, err := e.getBinaryExprValueAsNumber(leftValue, rightValue)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	switch node.Operator.TokenType {
	case token.TokenTypeGreaterThan:
		return controlflow.NewRegularResult(datavalue.Bool(leftNumber > rightNumber)), nil

	case token.TokenTypeGreaterThanOrEqual:
		return controlflow.NewRegularResult(datavalue.Bool(leftNumber >= rightNumber)), nil

	case token.TokenTypeLessThan:
		return controlflow.NewRegularResult(datavalue.Bool(leftNumber < rightNumber)), nil

	case token.TokenTypeLessThanOrEqual:
		return controlflow.NewRegularResult(datavalue.Bool(leftNumber <= rightNumber)), nil

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.Position(),
			node.Operator.Atom,
		)
	}
}
