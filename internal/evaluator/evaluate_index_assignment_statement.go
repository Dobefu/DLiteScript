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

	if arrayValue.Value.DataType() != datatype.DataTypeArray {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgTypeExpected,
			node.StartPosition(),
			node.Array.Expr(),
		)
	}

	indexValue, err := e.Evaluate(node.Index)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if indexValue.Value.DataType() != datatype.DataTypeNumber {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgTypeExpected,
			node.StartPosition(),
			node.Index.Expr(),
		)
	}

	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	array, err := arrayValue.Value.AsArray()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	index, err := indexValue.Value.AsNumber()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if index < 0 || int(index) >= len(array) {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgArrayIndexOutOfBounds,
			node.StartPosition(),
			node.Index.Expr(),
		)
	}

	array[int(index)] = rightValue.Value
	identifier, ok := node.Array.(*ast.Identifier)

	if ok {
		return e.assignVariable(
			identifier.Value,
			datavalue.Array(array...),
			node.StartPosition(),
		)
	}

	return controlflow.NewRegularResult(rightValue.Value), nil
}
