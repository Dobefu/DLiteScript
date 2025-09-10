package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/controlflow"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (e *Evaluator) evaluateShorthandAssignmentExpr(
	node *ast.ShorthandAssignmentExpr,
) (*controlflow.EvaluationResult, error) {
	rightValue, err := e.Evaluate(node.Right)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	leftValue, err := e.Evaluate(node.Left)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	binaryExpr := &ast.BinaryExpr{
		Left: node.Left,
		Operator: *token.NewToken(
			e.getBaseOperatorString(node.Operator.TokenType),
			e.getBaseOperator(node.Operator.TokenType),
			node.Operator.StartPos,
			node.Operator.EndPos,
		),
		Right:    node.Right,
		StartPos: node.StartPos,
		EndPos:   node.EndPos,
	}

	result, err := e.evaluateArithmeticBinaryExpr(
		leftValue.Value,
		rightValue.Value,
		binaryExpr,
	)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if identifier, ok := node.Left.(*ast.Identifier); ok {
		return e.assignVariable(
			identifier.Value,
			result.Value,
			identifier.StartPosition(),
		)
	}

	if indexExpr, ok := node.Left.(*ast.IndexExpr); ok {
		return e.assignArrayIndex(indexExpr, result.Value)
	}

	return controlflow.NewRegularResult(datavalue.Null()), err
}

func (e *Evaluator) getBaseOperator(shorthandType token.Type) token.Type {
	switch shorthandType {
	case token.TokenTypeOperationAddAssign:
		return token.TokenTypeOperationAdd

	case token.TokenTypeOperationSubAssign:
		return token.TokenTypeOperationSub

	case token.TokenTypeOperationMulAssign:
		return token.TokenTypeOperationMul

	case token.TokenTypeOperationDivAssign:
		return token.TokenTypeOperationDiv

	case token.TokenTypeOperationModAssign:
		return token.TokenTypeOperationMod

	case token.TokenTypeOperationPowAssign:
		return token.TokenTypeOperationPow

	default:
		return token.TokenTypeOperationAdd
	}
}

func (e *Evaluator) getBaseOperatorString(shorthandType token.Type) string {
	switch shorthandType {
	case token.TokenTypeOperationAddAssign:
		return "+"

	case token.TokenTypeOperationSubAssign:
		return "-"

	case token.TokenTypeOperationMulAssign:
		return "*"

	case token.TokenTypeOperationDivAssign:
		return "/"

	case token.TokenTypeOperationModAssign:
		return "%"

	case token.TokenTypeOperationPowAssign:
		return "**"

	default:
		return "+"
	}
}

func (e *Evaluator) assignArrayIndex(indexExpr *ast.IndexExpr, result datavalue.Value) (*controlflow.EvaluationResult, error) {
	arrayValue, err := e.Evaluate(indexExpr.Array)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	indexValue, err := e.Evaluate(indexExpr.Index)

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	array, err := arrayValue.Value.AsArray()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	index, err := indexValue.Value.AsNumber()

	if err != nil {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	if index < 0 || int(index) >= len(array) {
		return controlflow.NewRegularResult(datavalue.Null()), err
	}

	array[int(index)] = result
	identifier, hasIdentifier := indexExpr.Array.(*ast.Identifier)

	if hasIdentifier {
		return e.assignVariable(identifier.Value, datavalue.Array(array...), indexExpr.StartPosition())
	}

	return controlflow.NewRegularResult(result), nil
}
