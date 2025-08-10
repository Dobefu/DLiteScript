package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIdentifier(
	i *ast.Identifier,
) (datavalue.Value, error) {
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
