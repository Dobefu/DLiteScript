package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluatePrefixExpr(
	node *ast.PrefixExpr,
) (*controlflow.EvaluationResult, error) {
	rawResult, err := e.Evaluate(node.Operand)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	switch node.Operator.TokenType {
	case token.TokenTypeOperationSub:
		number, err := rawResult.Value.AsNumber()

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		return controlflow.NewRegularResult(datavalue.Number(-number)), nil

	case token.TokenTypeOperationAdd:
		number, err := rawResult.Value.AsNumber()

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		return controlflow.NewRegularResult(datavalue.Number(number)), nil

	case token.TokenTypeNot:
		boolean, err := rawResult.Value.AsBool()

		if err != nil {
			return controlflow.NewRegularResult(datavalue.Null()), err
		}

		return controlflow.NewRegularResult(datavalue.Bool(!boolean)), nil

	default:
		return controlflow.NewRegularResult(datavalue.Null()), errorutil.NewErrorAt(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgUnknownOperator,
			node.GetRange(),
			node.Operator.Atom,
		)
	}
}
