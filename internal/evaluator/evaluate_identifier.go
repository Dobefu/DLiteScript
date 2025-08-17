package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIdentifier(
	i *ast.Identifier,
) (*controlflow.EvaluationResult, error) {
	for idx := range e.blockScopesLen {
		scopedValue, hasScopedValue := e.blockScopes[e.blockScopesLen-idx-1][i.Value]

		if hasScopedValue {
			return controlflow.NewRegularResult(scopedValue.GetValue()), nil
		}
	}

	scopedValue, hasScopedValue := e.outerScope[i.Value]

	if hasScopedValue {
		return controlflow.NewRegularResult(scopedValue.GetValue()), nil
	}

	identifier, hasIdentifier := identifierRegistry[i.Value]

	if !hasIdentifier {
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.ErrorMsgUndefinedIdentifier,
			i.StartPosition(),
			i.Value,
		)
	}

	handlerResult, err := identifier.handler()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	return controlflow.NewRegularResult(handlerResult), nil
}
