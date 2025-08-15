package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateConstantDeclaration(
	node *ast.ConstantDeclaration,
) (*controlflow.EvaluationResult, error) {
	value, err := e.Evaluate(node.Value)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	var constant ScopedValue = &Constant{
		Value: value.Value,
		Type:  node.Type,
	}

	e.outerScope[node.Name] = constant

	return controlflow.NewRegularResult(datavalue.Null()), nil
}
