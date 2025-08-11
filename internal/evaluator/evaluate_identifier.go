package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIdentifier(
	i *ast.Identifier,
) (datavalue.Value, error) {
	for idx := range e.blockScopesLen {
		scopedValue, hasScopedValue := e.blockScopes[e.blockScopesLen-idx-1][i.Value]

		if hasScopedValue {
			return scopedValue.GetValue(), nil
		}
	}

	scopedValue, hasScopedValue := e.outerScope[i.Value]

	if hasScopedValue {
		return scopedValue.GetValue(), nil
	}

	identifier, hasIdentifier := identifierRegistry[i.Value]

	if !hasIdentifier {
		return datavalue.Null(), errorutil.NewErrorAt(
			errorutil.ErrorMsgUndefinedIdentifier,
			i.Position(),
			i.Value,
		)
	}

	return identifier.handler()
}
