package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateVariableDeclaration(
	node *ast.VariableDeclaration,
) (datavalue.Value, error) {
	value := datavalue.Null()

	if node.Value != nil {
		evaluatedValue, err := e.Evaluate(node.Value)

		if err != nil {
			return datavalue.Null(), err
		}

		value = evaluatedValue
	}

	var variable ScopedValue = &Variable{
		Value: value,
		Type:  node.Type,
	}

	if e.blockScopesLen > 0 {
		e.blockScopes[e.blockScopesLen-1][node.Name] = variable

		return datavalue.Null(), nil
	}

	e.outerScope[node.Name] = variable

	return datavalue.Null(), nil
}
