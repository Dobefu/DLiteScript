package evaluator

import (
	"fmt"
	"io"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateBlockStatement(
	node *ast.BlockStatement,
) (*controlflow.EvaluationResult, error) {
	e.pushBlockScope()

	result := controlflow.NewRegularResult(datavalue.Null())

	for _, statement := range node.Statements {
		val, err := e.Evaluate(statement)

		if err != nil {
			e.popBlockScope()

			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		result = val

		if !result.IsNormalResult() {
			e.popBlockScope()

			return result, nil
		}

		if e.buf.Len() > 0 && e.outFile != io.Discard {
			_, err := fmt.Fprint(e.outFile, e.buf.String())

			if err != nil {
				return controlflow.NewRegularResult(datavalue.Null()), err
			}

			e.buf.Reset()
		}
	}

	e.popBlockScope()

	return result, nil
}
