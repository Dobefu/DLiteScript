package evaluator

import (
	"math"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateBinaryExpr(
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	leftValue, err := e.Evaluate(node.Left)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	switch node.Operator.TokenType {
	case
		token.TokenTypeOperationAdd,
		token.TokenTypeOperationSub,
		token.TokenTypeOperationMul,
		token.TokenTypeOperationDiv,
		token.TokenTypeOperationMod,
		token.TokenTypeOperationPow:
		return e.evaluateArithmeticBinaryExpr(leftValue.Value, rightValue.Value, node)

	case
		token.TokenTypeEqual,
		token.TokenTypeNotEqual:
		return e.evaluateEqualityBinaryExpr(leftValue.Value, rightValue.Value, node)

	case
		token.TokenTypeGreaterThan,
		token.TokenTypeGreaterThanOrEqual,
		token.TokenTypeLessThan,
		token.TokenTypeLessThanOrEqual:
		return e.evaluateComparisonBinaryExpr(leftValue.Value, rightValue.Value, node)

	case
		token.TokenTypeLogicalAnd,
		token.TokenTypeLogicalOr:
		return e.evaluateLogicalBinaryExpr(leftValue.Value, rightValue.Value, node)
	}

	return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
		errorutil.ErrorMsgUnknownOperator,
		node.StartPosition(),
		node.Operator.Atom,
	)
}

func (e *Evaluator) evaluateArithmeticBinaryExpr(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	leftNumber, rightNumber, err := e.getBinaryExprValueAsNumber(leftValue, rightValue)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if node.Operator.TokenType == token.TokenTypeOperationDiv && rightNumber == 0 {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(errorutil.ErrorMsgDivByZero, node.StartPosition())
	}

	if node.Operator.TokenType == token.TokenTypeOperationMod && rightNumber == 0 {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(errorutil.ErrorMsgModByZero, node.StartPosition())
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return controlflow.NewRegularResult(datavalue.Number(leftNumber + rightNumber)), nil

	case token.TokenTypeOperationSub:
		return controlflow.NewRegularResult(datavalue.Number(leftNumber - rightNumber)), nil

	case token.TokenTypeOperationMul:
		return controlflow.NewRegularResult(datavalue.Number(leftNumber * rightNumber)), nil

	case token.TokenTypeOperationDiv:
		return controlflow.NewRegularResult(datavalue.Number(leftNumber / rightNumber)), nil

	case token.TokenTypeOperationMod:
		return controlflow.NewRegularResult(datavalue.Number(math.Mod(leftNumber, rightNumber))), nil

	case token.TokenTypeOperationPow:
		return controlflow.NewRegularResult(datavalue.Number(math.Pow(leftNumber, rightNumber))), nil

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.StartPosition(),
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
