package evaluator

import (
	"fmt"
	"io"

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

		if e.buf.Len() > 0 && e.outFile != io.Discard {
			_, err := fmt.Fprint(e.outFile, e.buf.String())

			if err != nil {
				return datavalue.Null(), err
			}

			e.buf.Reset()
		}
	}

	e.popBlockScope()

	return result, nil
}
