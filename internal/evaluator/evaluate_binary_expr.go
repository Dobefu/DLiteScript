package evaluator

import (
	"math"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateBinaryExpr(node *ast.BinaryExpr) (datavalue.Value, error) {
	leftValue, err := e.Evaluate(node.Left)

	if err != nil {
		return datavalue.Null(), err
	}

	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return datavalue.Null(), err
	}

	leftEvaluated, err := leftValue.AsNumber()

	if err != nil {
		return datavalue.Null(), err
	}

	rightEvaluated, err := rightValue.AsNumber()

	if err != nil {
		return datavalue.Null(), err
	}

	if node.Operator.TokenType == token.TokenTypeOperationDiv && rightEvaluated == 0 {
		return datavalue.Null(), errorutil.NewErrorAt(errorutil.ErrorMsgDivByZero, node.Position())
	}

	if node.Operator.TokenType == token.TokenTypeOperationMod && rightEvaluated == 0 {
		return datavalue.Null(), errorutil.NewErrorAt(errorutil.ErrorMsgModByZero, node.Position())
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return datavalue.Number(leftEvaluated + rightEvaluated), nil

	case token.TokenTypeOperationSub:
		return datavalue.Number(leftEvaluated - rightEvaluated), nil

	case token.TokenTypeOperationMul:
		return datavalue.Number(leftEvaluated * rightEvaluated), nil

	case token.TokenTypeOperationDiv:
		return datavalue.Number(leftEvaluated / rightEvaluated), nil

	case token.TokenTypeOperationMod:
		return datavalue.Number(math.Mod(leftEvaluated, rightEvaluated)), nil

	case token.TokenTypeOperationPow:
		return datavalue.Number(math.Pow(leftEvaluated, rightEvaluated)), nil

	default:
		return datavalue.Null(), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.Position(),
			node.Operator.Atom,
		)
	}
}
