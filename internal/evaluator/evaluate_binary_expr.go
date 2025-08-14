package evaluator

import (
	"math"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateBinaryExpr(
	node *ast.BinaryExpr,
) (datavalue.Value, error) {
	leftValue, err := e.Evaluate(node.Left)

	if err != nil {
		return datavalue.Null(), err
	}

	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return datavalue.Null(), err
	}

	switch node.Operator.TokenType {
	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub,
		token.TokenTypeOperationMul,
		token.TokenTypeOperationDiv,
		token.TokenTypeOperationMod,
		token.TokenTypeOperationPow:
		return e.evaluateArithmeticBinaryExpr(leftValue, rightValue, node)

	case
		token.TokenTypeEqual,
		token.TokenTypeNotEqual:
		return e.evaluateEqualityBinaryExpr(leftValue, rightValue, node)

	case
		token.TokenTypeGreaterThan,
		token.TokenTypeGreaterThanOrEqual,
		token.TokenTypeLessThan,
		token.TokenTypeLessThanOrEqual:
		return e.evaluateComparisonBinaryExpr(leftValue, rightValue, node)

	case
		token.TokenTypeLogicalAnd,
		token.TokenTypeLogicalOr:
		return e.evaluateLogicalBinaryExpr(leftValue, rightValue, node)
	}

	return datavalue.Null(), errorutil.NewErrorAt(
		errorutil.ErrorMsgUnknownOperator,
		node.Position(),
		node.Operator.Atom,
	)
}

func (e *Evaluator) evaluateArithmeticBinaryExpr(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (datavalue.Value, error) {
	leftNumber, rightNumber, err := e.getBinaryExprValueAsNumber(leftValue, rightValue)

	if err != nil {
		return datavalue.Null(), err
	}

	if node.Operator.TokenType == token.TokenTypeOperationDiv && rightNumber == 0 {
		return datavalue.Null(), errorutil.NewErrorAt(errorutil.ErrorMsgDivByZero, node.Position())
	}

	if node.Operator.TokenType == token.TokenTypeOperationMod && rightNumber == 0 {
		return datavalue.Null(), errorutil.NewErrorAt(errorutil.ErrorMsgModByZero, node.Position())
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return datavalue.Number(leftNumber + rightNumber), nil

	case token.TokenTypeOperationSub:
		return datavalue.Number(leftNumber - rightNumber), nil

	case token.TokenTypeOperationMul:
		return datavalue.Number(leftNumber * rightNumber), nil

	case token.TokenTypeOperationDiv:
		return datavalue.Number(leftNumber / rightNumber), nil

	case token.TokenTypeOperationMod:
		return datavalue.Number(math.Mod(leftNumber, rightNumber)), nil

	case token.TokenTypeOperationPow:
		return datavalue.Number(math.Pow(leftNumber, rightNumber)), nil

	default:
		return datavalue.Null(), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.Position(),
			node.Operator.Atom,
		)
	}
}

func (e *Evaluator) getBinaryExprValueAsBool(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
) (bool, bool, error) {
	leftBool, err := leftValue.AsBool()

	if err != nil {
		return false, false, err
	}

	rightBool, err := rightValue.AsBool()

	if err != nil {
		return false, false, err
	}

	return leftBool, rightBool, nil
}

func (e *Evaluator) getBinaryExprValueAsNumber(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
) (float64, float64, error) {
	leftNumber, err := leftValue.AsNumber()

	if err != nil {
		return 0, 0, err
	}

	rightNumber, err := rightValue.AsNumber()

	if err != nil {
		return 0, 0, err
	}

	return leftNumber, rightNumber, nil
}

func (e *Evaluator) getBinaryExprValueAsString(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
) (string, string, error) {
	leftString, err := leftValue.AsString()

	if err != nil {
		return "", "", err
	}

	rightString, err := rightValue.AsString()

	if err != nil {
		return "", "", err
	}

	return leftString, rightString, nil
}
