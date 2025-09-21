package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateVariableDeclaration(
	node *ast.VariableDeclaration,
) (*controlflow.EvaluationResult, error) {
	var value *controlflow.EvaluationResult

	if node.Value != nil {
		evaluatedValue, err := e.Evaluate(node.Value)

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		value = evaluatedValue
	} else {
		zeroValue := e.getZeroValueForType(node.Type)
		value = controlflow.NewRegularResult(zeroValue)
	}

	if node.Type[:2] != "[]" &&
		node.Type != datatype.DataTypeAny.AsString() &&
		value.Value.DataType.AsString() != node.Type {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeMismatch,
			node.StartPosition(),
			node.Type,
			value.Value.DataType.AsString(),
		)
	}

	var variable ScopedValue = &Variable{
		Value: value.Value,
		Type:  node.Type,
	}

	if e.blockScopesLen > 0 {
		e.blockScopes[e.blockScopesLen-1][node.Name] = variable

		return controlflow.NewRegularResult(datavalue.Null()), nil
	}

	e.outerScope[node.Name] = variable

	return controlflow.NewRegularResult(datavalue.Null()), nil
}

// getZeroValueForType returns the zero value for a given type string.
func (e *Evaluator) getZeroValueForType(typeStr string) datavalue.Value {
	switch typeStr {
	case
		"string":
		return datavalue.String("")

	case
		"number":
		return datavalue.Number(0)

	case
		"bool":
		return datavalue.Bool(false)

	case
		"any":
		return datavalue.Any(nil)

	case
		"error":
		return datavalue.Error(nil)

	default:
		if len(typeStr) > 2 && typeStr[:2] == "[]" {
			return datavalue.Array()
		}

		return datavalue.Null()
	}
}
