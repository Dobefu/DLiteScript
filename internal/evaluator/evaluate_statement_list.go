package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateStatementList(
	list *ast.StatementList,
) (datavalue.Value, error) {
	lastResult := datavalue.Null()

	for _, statement := range list.Statements {
		result, err := e.Evaluate(statement)

		if err != nil {
			return datavalue.Null(), err
		}

		lastResult = result
	}

	return lastResult, nil
}
