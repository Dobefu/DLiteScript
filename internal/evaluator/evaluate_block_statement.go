package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateBlockStatement(node *ast.BlockStatement) (datavalue.Value, error) {
	e.pushBlockScope()

	result := datavalue.Null()

	for _, statement := range node.Statements {
		val, err := e.Evaluate(statement)

		if err != nil {
			e.popBlockScope()

			return datavalue.Null(), err
		}

		result = val
	}

	e.popBlockScope()

	return result, nil
}
