package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateLogicalBinaryExpr(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	leftBool, rightBool, err := e.getBinaryExprValueAsBool(leftValue, rightValue)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	switch node.Operator.TokenType {
	case token.TokenTypeLogicalAnd:
		return controlflow.NewRegularResult(
			datavalue.Bool(leftBool && rightBool),
		), nil

	case token.TokenTypeLogicalOr:
		return controlflow.NewRegularResult(
			datavalue.Bool(leftBool || rightBool),
		), nil

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.Position(),
			node.Operator.Atom,
		)
	}
}
