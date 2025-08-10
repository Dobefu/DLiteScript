package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateConstantDeclaration(
	node *ast.ConstantDeclaration,
) (datavalue.Value, error) {
	value, err := e.Evaluate(node.Value)

	if err != nil {
		return datavalue.Null(), err
	}

	var constant ScopedValue = &Constant{
		Value: value,
		Type:  node.Type,
	}

	e.outerScope[node.Name] = constant

	return datavalue.Null(), nil
}
