package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIfStatement(node *ast.IfStatement) (datavalue.Value, error) {
	expr, err := e.Evaluate(node.Condition)

	if err != nil {
		return datavalue.Null(), err
	}

	if expr.DataType() != datatype.DataTypeBool {
		return datavalue.Null(), errorutil.NewError(
			errorutil.ErrorMsgTypeExpected,
			"bool",
			expr.DataType().AsString(),
		)
	}

	exprResult, err := expr.AsBool()

	if err != nil {
		return datavalue.Null(), err
	}

	if exprResult {
		return e.Evaluate(node.ThenBlock)
	}

	if node.ElseBlock != nil {
		return e.Evaluate(node.ElseBlock)
	}

	return datavalue.Null(), nil
}
