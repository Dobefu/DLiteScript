package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIfStatement(
	node *ast.IfStatement,
) (*controlflow.EvaluationResult, error) {
	expr, err := e.Evaluate(node.Condition)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if expr.Value.DataType() != datatype.DataTypeBool {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			node.StartPosition(),
			"bool",
			expr.Value.DataType().AsString(),
		)
	}

	exprResult, err := expr.Value.AsBool()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if exprResult {
		return e.Evaluate(node.ThenBlock)
	}

	if node.ElseBlock != nil {
		return e.Evaluate(node.ElseBlock)
	}

	return controlflow.NewRegularResult(datavalue.Null()), nil
}
