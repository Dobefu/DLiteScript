package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateAssignmentStatement(
	node *ast.AssignmentStatement,
) (*controlflow.EvaluationResult, error) {
	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	for idx := range e.blockScopesLen {
		_, hasScopedValue := e.blockScopes[e.blockScopesLen-idx-1][node.Left.Value]

		if hasScopedValue {
			variable, isVariable := e.blockScopes[e.blockScopesLen-idx-1][node.Left.Value].(*Variable)

			if !isVariable {
				return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
					errorutil.ErrorMsgReassignmentToConstant,
					node.Left.StartPosition(),
					node.Left.Value,
				)
			}

			variable.Value = rightValue.Value

			return controlflow.NewRegularResult(rightValue.Value), nil
		}
	}

	_, hasScopedValue := e.outerScope[node.Left.Value]

	if hasScopedValue {
		variable, isVariable := e.outerScope[node.Left.Value].(*Variable)

		if !isVariable {
			return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
				errorutil.ErrorMsgReassignmentToConstant,
				node.Left.StartPosition(),
				node.Left.Value,
			)
		}

		variable.Value = rightValue.Value

		return controlflow.NewRegularResult(rightValue.Value), nil
	}

	return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
		errorutil.ErrorMsgUndefinedIdentifier,
		node.Left.StartPosition(),
		node.Left.Value,
	)
}
