package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateArrayConcatenation(
	leftArray []datavalue.Value,
	rightArray []datavalue.Value,
	node *ast.BinaryExpr,
) (*controlflow.EvaluationResult, error) {
	if len(leftArray) > 0 && len(rightArray) > 0 {
		leftType := leftArray[0].DataType
		rightType := rightArray[0].DataType

		if leftType != rightType {
			return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeMismatch,
				node.GetRange(),
				leftType.AsString(),
				rightType.AsString(),
			)
		}
	}

	return controlflow.NewRegularResult(
		datavalue.Array(append(leftArray, rightArray...)...),
	), nil
}
