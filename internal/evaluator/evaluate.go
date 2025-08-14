package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// Evaluate runs the evaluation logic.
func (e *Evaluator) Evaluate(currentAst ast.ExprNode) (datavalue.Value, error) {
	if currentAst == nil {
		return datavalue.Null(), nil
	}

	switch node := currentAst.(type) {
	case *ast.StatementList:
		return e.evaluateStatementList(node)

	case *ast.BinaryExpr:
		return e.evaluateBinaryExpr(node)

	case *ast.NumberLiteral:
		return e.evaluateNumberLiteral(node)

	case *ast.StringLiteral:
		return e.evaluateStringLiteral(node)

	case *ast.BoolLiteral:
		return e.evaluateBoolLiteral(node)

	case *ast.NullLiteral:
		return e.evaluateNullLiteral()

	case *ast.PrefixExpr:
		return e.evaluatePrefixExpr(node)

	case *ast.FunctionCall:
		return e.evaluateFunctionCall(node)

	case *ast.Identifier:
		return e.evaluateIdentifier(node)

	case *ast.VariableDeclaration:
		return e.evaluateVariableDeclaration(node)

	case *ast.ConstantDeclaration:
		return e.evaluateConstantDeclaration(node)

	case *ast.IfStatement:
		return e.evaluateIfStatement(node)

	case *ast.BlockStatement:
		return e.evaluateBlockStatement(node)

	case *ast.AssignmentStatement:
		return e.evaluateAssignmentStatement(node)

	default:
		return datavalue.Null(), errorutil.NewErrorAt(
			errorutil.ErrorMsgUnknownNodeType,
			node.Position(),
			node,
		)
	}
}

// Output returns the current output buffer contents.
func (e *Evaluator) Output() string {
	return e.buf.String()
}
