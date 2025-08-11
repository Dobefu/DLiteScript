package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateAssignmentStatement(node *ast.AssignmentStatement) (datavalue.Value, error) {
	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return datavalue.Null(), err
	}

	for idx := range e.blockScopesLen {
		if _, hasScopedValue := e.blockScopes[e.blockScopesLen-idx-1][node.Left.Value]; hasScopedValue {
			variable, isVariable := e.blockScopes[e.blockScopesLen-idx-1][node.Left.Value].(*Variable)

			if !isVariable {
				return datavalue.Null(), errorutil.NewErrorAt(
					errorutil.ErrorMsgReassignmentToConstant,
					node.Left.Position(),
					node.Left.Value,
				)
			}

			variable.Value = rightValue

			return rightValue, nil
		}
	}

	if _, hasScopedValue := e.outerScope[node.Left.Value]; hasScopedValue {
		variable, isVariable := e.outerScope[node.Left.Value].(*Variable)

		if !isVariable {
			return datavalue.Null(), errorutil.NewErrorAt(
				errorutil.ErrorMsgReassignmentToConstant,
				node.Left.Position(),
				node.Left.Value,
			)
		}

		variable.Value = rightValue

		return rightValue, nil
	}

	return datavalue.Null(), errorutil.NewErrorAt(
		errorutil.ErrorMsgUndefinedIdentifier,
		node.Left.Position(),
		node.Left.Value,
	)
}
