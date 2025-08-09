package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateIdentifier(
	i *ast.Identifier,
) (float64, error) {
	identifier, ok := identifierRegistry[i.Value]

	if !ok {
		return 0, errorutil.NewErrorAt(errorutil.ErrorMsgUndefinedIdentifier, i.Position(), i.Value)
	}

	return identifier.handler()
}
