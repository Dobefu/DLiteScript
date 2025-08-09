package evaluator

import (
	"math"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateBinaryExpr(node *ast.BinaryExpr) (float64, error) {
	leftEvaluated, err := e.Evaluate(node.Left)

	if err != nil {
		return 0, err
	}

	rightEvaluated, err := e.Evaluate(node.Right)

	if err != nil {
		return 0, err
	}

	if node.Operator.TokenType == token.TokenTypeOperationDiv && rightEvaluated == 0 {
		return 0, errorutil.NewErrorAt(errorutil.ErrorMsgDivByZero, node.Position())
	}

	if node.Operator.TokenType == token.TokenTypeOperationMod && rightEvaluated == 0 {
		return 0, errorutil.NewErrorAt(errorutil.ErrorMsgModByZero, node.Position())
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return leftEvaluated + rightEvaluated, nil

	case token.TokenTypeOperationSub:
		return leftEvaluated - rightEvaluated, nil

	case token.TokenTypeOperationMul:
		return leftEvaluated * rightEvaluated, nil

	case token.TokenTypeOperationDiv:
		return leftEvaluated / rightEvaluated, nil

	case token.TokenTypeOperationMod:
		return math.Mod(leftEvaluated, rightEvaluated), nil

	case token.TokenTypeOperationPow:
		return math.Pow(leftEvaluated, rightEvaluated), nil

	default:
		return 0, errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.Position(),
			node.Operator.Atom,
		)
	}
}
