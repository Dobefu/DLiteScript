package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) assignVariable(
	varName string,
	value datavalue.Value,
	startPos int,
) (*controlflow.EvaluationResult, error) {
	for idx := range e.blockScopesLen {
		_, hasScopedValue := e.blockScopes[e.blockScopesLen-idx-1][varName]

		if hasScopedValue {
			variable, isVariable := e.blockScopes[e.blockScopesLen-idx-1][varName].(*Variable)

			if !isVariable {
				return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
					errorutil.ErrorMsgReassignmentToConstant,
					startPos,
					varName,
				)
			}

			variable.Value = value

			return controlflow.NewRegularResult(value), nil
		}
	}

	_, hasScopedValue := e.outerScope[varName]

	if hasScopedValue {
		variable, isVariable := e.outerScope[varName].(*Variable)

		if !isVariable {
			return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
				errorutil.ErrorMsgReassignmentToConstant,
				startPos,
				varName,
			)
		}

		variable.Value = value

		return controlflow.NewRegularResult(value), nil
	}

	return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
		errorutil.ErrorMsgUndefinedIdentifier,
		startPos,
		varName,
	)
}
