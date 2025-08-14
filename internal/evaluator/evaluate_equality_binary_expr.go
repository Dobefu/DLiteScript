package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateEqualityBinaryExpr(
	leftValue datavalue.Value,
	rightValue datavalue.Value,
	node *ast.BinaryExpr,
) (datavalue.Value, error) {
	if leftValue.DataType() != rightValue.DataType() {
		result := node.Operator.TokenType == token.TokenTypeNotEqual

		return datavalue.Bool(result), nil
	}

	switch leftValue.DataType() {
	case datatype.DataTypeNumber:
		leftNumber, rightNumber, err := e.getBinaryExprValueAsNumber(leftValue, rightValue)

		if err != nil {
			return datavalue.Null(), err
		}

		isEqual := leftNumber == rightNumber
		result := isEqual == (node.Operator.TokenType == token.TokenTypeEqual)

		return datavalue.Bool(result), nil

	case datatype.DataTypeString:
		leftString, rightString, err := e.getBinaryExprValueAsString(leftValue, rightValue)

		if err != nil {
			return datavalue.Null(), err
		}

		isEqual := leftString == rightString
		result := isEqual == (node.Operator.TokenType == token.TokenTypeEqual)

		return datavalue.Bool(result), nil

	case datatype.DataTypeBool:
		leftBool, rightBool, err := e.getBinaryExprValueAsBool(leftValue, rightValue)

		if err != nil {
			return datavalue.Null(), err
		}

		isEqual := leftBool == rightBool
		result := isEqual == (node.Operator.TokenType == token.TokenTypeEqual)

		return datavalue.Bool(result), nil

	case datatype.DataTypeNull:
		result := node.Operator.TokenType == token.TokenTypeEqual

		return datavalue.Bool(result), nil

	default:
		return datavalue.Null(), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownOperator,
			node.Position(),
			node.Operator.Atom,
		)
	}
}
