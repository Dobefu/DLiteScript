package evaluator

import (
	"fmt"
	"io"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateStatementList(
	list *ast.StatementList,
) (*controlflow.EvaluationResult, error) {
	lastResult := controlflow.NewRegularResult(datavalue.Null())

	for _, statement := range list.Statements {
		result, err := e.Evaluate(statement)

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		lastResult = result

		if e.buf.Len() > 0 && e.outFile != io.Discard {
			_, err := fmt.Fprint(e.outFile, e.buf.String())

			if err != nil {
				return controlflow.NewRegularResult(datavalue.Null()), err
			}

			e.buf.Reset()
		}
	}

	return lastResult, nil
}
