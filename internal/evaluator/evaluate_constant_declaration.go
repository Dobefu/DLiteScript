package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateConstantDeclaration(
	node *ast.ConstantDeclaration,
) (*controlflow.EvaluationResult, error) {
	value, err := e.Evaluate(node.Value)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if value.Value.DataType().AsString() != node.Type {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeMismatch,
			node.StartPosition(),
			node.Type,
			value.Value.DataType().AsString(),
		)
	}

	var constant ScopedValue = &Constant{
		Value: value.Value,
		Type:  node.Type,
	}

	if e.blockScopesLen > 0 {
		e.blockScopes[e.blockScopesLen-1][node.Name] = constant

		return controlflow.NewRegularResult(datavalue.Null()), nil
	}

	e.outerScope[node.Name] = constant

	return controlflow.NewRegularResult(datavalue.Null()), nil
}
