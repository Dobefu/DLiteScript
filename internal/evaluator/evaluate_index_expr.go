package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIndexExpr(
	node *ast.IndexExpr,
) (*controlflow.EvaluationResult, error) {
	value, err := e.Evaluate(node.Array)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if value.Value.DataType() != datatype.DataTypeArray {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluation,
			errorutil.ErrorMsgTypeExpected,
			node.StartPosition(),
			datatype.DataTypeArray.AsString(),
			value.Value.DataType().AsString(),
		)
	}

	idxValue, err := e.Evaluate(node.Index)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if idxValue.Value.DataType() != datatype.DataTypeNumber {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluation,
			errorutil.ErrorMsgTypeExpected,
			node.StartPosition(),
			datatype.DataTypeNumber.AsString(),
			idxValue.Value.DataType().AsString(),
		)
	}

	array, err := value.Value.AsArray()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	idx, err := idxValue.Value.AsNumber()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if idx < 0 || int(idx) >= len(array) {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluation,
			errorutil.ErrorMsgArrayIndexOutOfBounds,
			node.StartPosition(),
			node.Index.Expr(),
		)
	}

	return controlflow.NewRegularResult(array[int(idx)]), nil
}
