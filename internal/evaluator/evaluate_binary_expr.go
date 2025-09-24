package evaluator

import (
	"fmt"
	"math"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
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
		return e.evaluateArithmeticBinaryExpr(
			leftValue.Value,
			rightValue.Value,
			node,
		)

	case
		token.TokenTypeEqual,
		token.TokenTypeNotEqual:
		return e.evaluateEqualityBinaryExpr(
			leftValue.Value,
			rightValue.Value,
			node,
		)

	case
		token.TokenTypeGreaterThan,
		token.TokenTypeGreaterThanOrEqual,
		token.TokenTypeLessThan,
		token.TokenTypeLessThanOrEqual:
		return e.evaluateComparisonBinaryExpr(
			leftValue.Value,
			rightValue.Value,
			node,
		)

	case
		token.TokenTypeLogicalAnd,
		token.TokenTypeLogicalOr:
		return e.evaluateLogicalBinaryExpr(
			leftValue.Value,
			rightValue.Value,
			node,
		)
	}

	return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgUnknownOperator,
		node.GetRange().Start.Offset,
		node.Operator.Atom,
	)
}

func (e *Evaluator) evaluateArithmeticBinaryExpr(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	if leftValue.DataType != rightValue.DataType {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			node.GetRange().Start.Offset,
			rightValue.DataType.AsString(),
			leftValue.DataType.AsString(),
		)
	}

	switch leftValue.DataType {
	case
		datatype.DataTypeNumber:
		return e.evaluateArithmeticBinaryExprNumber(
			leftValue,
			rightValue,
			node,
		)

	case
		datatype.DataTypeArray:
		return e.evaluateArithmeticBinaryExprArray(
			leftValue,
			rightValue,
			node,
		)

	case
		datatype.DataTypeString:
		return e.evaluateArithmeticBinaryExprString(
			leftValue,
			rightValue,
			node,
		)

	case
		datatype.DataTypeBool,
		datatype.DataTypeFunction,
		datatype.DataTypeTuple,
		datatype.DataTypeError,
		datatype.DataTypeAny,
		datatype.DataTypeNull:
		fallthrough

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgCannotConcat,
			node.GetRange().Start.Offset,
			leftValue.DataType.AsString(),
			rightValue.DataType.AsString(),
		)
	}
}

func (e *Evaluator) evaluateArithmeticBinaryExprNumber(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	leftNumber, rightNumber, err := e.getBinaryExprValueAsNumber(
		leftValue,
		rightValue,
	)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if node.Operator.TokenType == token.TokenTypeOperationDiv &&
		rightNumber == 0 {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgDivByZero,
			node.GetRange().Start.Offset,
		)
	}

	if node.Operator.TokenType == token.TokenTypeOperationMod &&
		rightNumber == 0 {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgModByZero,
			node.GetRange().Start.Offset,
		)
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return controlflow.NewRegularResult(
			datavalue.Number(leftNumber + rightNumber),
		), nil

	case token.TokenTypeOperationSub:
		return controlflow.NewRegularResult(
			datavalue.Number(leftNumber - rightNumber),
		), nil

	case token.TokenTypeOperationMul:
		return controlflow.NewRegularResult(
			datavalue.Number(leftNumber * rightNumber),
		), nil

	case token.TokenTypeOperationDiv:
		return controlflow.NewRegularResult(
			datavalue.Number(leftNumber / rightNumber),
		), nil

	case token.TokenTypeOperationMod:
		return controlflow.NewRegularResult(
			datavalue.Number(math.Mod(leftNumber, rightNumber)),
		), nil

	case token.TokenTypeOperationPow:
		return controlflow.NewRegularResult(
			datavalue.Number(math.Pow(leftNumber, rightNumber)),
		), nil

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgUnknownOperator,
			node.GetRange().Start.Offset,
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
		return false, false, fmt.Errorf(
			"could not get binary expr value as bool: %s",
			err.Error(),
		)
	}

	rightBool, err := rightValue.AsBool()

	if err != nil {
		return false, false, fmt.Errorf(
			"could not get binary expr value as bool: %s",
			err.Error(),
		)
	}

	return leftBool, rightBool, nil
}

func (e *Evaluator) getBinaryExprValueAsNumber(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
) (float64, float64, error) {
	leftNumber, err := leftValue.AsNumber()

	if err != nil {
		return 0, 0, fmt.Errorf(
			"could not get binary expr value as number: %s",
			err.Error(),
		)
	}

	rightNumber, err := rightValue.AsNumber()

	if err != nil {
		return 0, 0, fmt.Errorf(
			"could not get binary expr value as number: %s",
			err.Error(),
		)
	}

	return leftNumber, rightNumber, nil
}

func (e *Evaluator) getBinaryExprValueAsString(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
) (string, string, error) {
	leftString, err := leftValue.AsString()

	if err != nil {
		return "", "", fmt.Errorf(
			"could not get binary expr value as string: %s",
			err.Error(),
		)
	}

	rightString, err := rightValue.AsString()

	if err != nil {
		return "", "", fmt.Errorf(
			"could not get binary expr value as string: %s",
			err.Error(),
		)
	}

	return leftString, rightString, nil
}

func (e *Evaluator) evaluateArithmeticBinaryExprArray(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	leftArray, err := leftValue.AsArray()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	rightArray, err := rightValue.AsArray()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return e.evaluateArrayConcatenation(leftArray, rightArray, node)

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgUnknownOperator,
			node.GetRange().Start.Offset,
			node.Operator.Atom,
		)
	}
}

func (e *Evaluator) evaluateArithmeticBinaryExprString(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	leftString, rightString, err := e.getBinaryExprValueAsString(
		leftValue,
		rightValue,
	)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationAdd:
		return controlflow.NewRegularResult(
			datavalue.String(fmt.Sprintf("%s%s", leftString, rightString)),
		), nil

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgUnknownOperator,
			node.GetRange().Start.Offset,
			node.Operator.Atom,
		)
	}
}
