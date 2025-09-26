package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func (e *Evaluator) evaluateAssignmentStatement(
	node *ast.AssignmentStatement,
) (*controlflow.EvaluationResult, error) {
	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	return e.assignVariable(
		node.Left.Value,
		rightValue.Value,
		node.Left.GetRange(),
	)
}
