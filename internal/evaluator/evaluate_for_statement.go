package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateForStatement(
	node *ast.ForStatement,
) (datavalue.Value, error) {
	result := datavalue.Null()

	for {
		if node.Condition != nil {
			evaluatedCondition, err := e.Evaluate(node.Condition)

			if err != nil {
				return datavalue.Null(), err
			}

			conditionResultBool, err := evaluatedCondition.AsBool()

			if err != nil {
				return datavalue.Null(), err
			}

			if !conditionResultBool {
				break
			}
		}

		evaluatedBody, err := e.Evaluate(node.Body)

		if err != nil {
			return datavalue.Null(), err
		}

		result = evaluatedBody
	}

	return result, nil
}
