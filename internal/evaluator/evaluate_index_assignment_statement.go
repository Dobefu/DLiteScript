package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIndexAssignmentStatement(
	node *ast.IndexAssignmentStatement,
) (*controlflow.EvaluationResult, error) {
	arrayValue, err := e.Evaluate(node.Array)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if arrayValue.Value.DataType != datatype.DataTypeArray {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			node.StartPosition(),
			datatype.DataTypeArray.AsString(),
			node.Array.Expr(),
		)
	}

	indexValue, err := e.Evaluate(node.Index)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if indexValue.Value.DataType != datatype.DataTypeNumber {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			node.StartPosition(),
			datatype.DataTypeNumber.AsString(),
			node.Index.Expr(),
		)
	}

	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	array, _ := arrayValue.Value.AsArray()
	index, _ := indexValue.Value.AsNumber()

	if index < 0 || int(index) >= len(array) {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgArrayIndexOutOfBounds,
			node.StartPosition(),
			node.Index.Expr(),
		)
	}

	array[int(index)] = rightValue.Value
	identifier, hasIdentifier := node.Array.(*ast.Identifier)

	if hasIdentifier {
		return e.assignVariable(
			identifier.Value,
			datavalue.Array(array...),
			node.StartPosition(),
		)
	}

	return controlflow.NewRegularResult(rightValue.Value), nil
}
