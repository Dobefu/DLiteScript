package formatter

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (f *Formatter) formatNode(
	node ast.ExprNode,
	result *strings.Builder,
	depth int,
) {
	switch n := node.(type) {
	case *ast.CommentLiteral:
		f.formatComment(n, result, depth)

	case *ast.StatementList:
		f.formatStatementList(n, result, depth)

	case *ast.BinaryExpr:
		f.formatBinaryExpr(n, result, depth)

	case *ast.NumberLiteral:
		f.formatNumberLiteral(n, result, depth)

	case *ast.StringLiteral:
		f.formatStringLiteral(n, result, depth)

	case *ast.BoolLiteral:
		f.formatBoolLiteral(n, result, depth)

	case *ast.AnyLiteral:
		f.formatAnyLiteral(n, result, depth)

	case *ast.NullLiteral:
		f.formatNullLiteral(n, result, depth)

	case *ast.PrefixExpr:
		f.formatPrefixExpr(n, result, depth)

	case *ast.FunctionCall:
		f.formatFunctionCall(n, result, depth)

	case *ast.Identifier:
		f.formatIdentifier(n, result, depth)

	case *ast.VariableDeclaration:
		f.formatVariableDeclaration(n, result, depth)

	case *ast.ConstantDeclaration:
		f.formatConstantDeclaration(n, result, depth)

	case *ast.IfStatement:
		f.formatIfStatement(n, result, depth)

	case *ast.ForStatement:
		f.formatForStatement(n, result, depth)

	case *ast.BreakStatement:
		f.formatBreakStatement(n, result, depth)

	case *ast.ContinueStatement:
		f.formatContinueStatement(n, result, depth)

	case *ast.BlockStatement:
		f.formatBlockStatement(n, result, depth, true)

	case *ast.AssignmentStatement:
		f.formatAssignmentStatement(n, result, depth)

	case *ast.IndexAssignmentStatement:
		f.formatIndexAssignmentStatement(n, result, depth)

	case *ast.ShorthandAssignmentExpr:
		f.formatShorthandAssignmentExpr(n, result, depth)

	case *ast.FuncDeclarationStatement:
		f.formatFuncDeclarationStatement(n, result, depth)

	case *ast.ReturnStatement:
		f.formatReturnStatement(n, result, depth)

	case *ast.SpreadExpr:
		f.formatSpreadExpr(n, result, depth)

	case *ast.ArrayLiteral:
		f.formatArrayLiteral(n, result, depth)

	case *ast.IndexExpr:
		f.formatIndexExpr(n, result, depth)

	case *ast.ImportStatement:
		f.formatImportStatement(n, result, depth)

	default:
		f.addWhitespace(result, depth)
		result.WriteString(n.Expr())
		result.WriteString("\n")
	}
}
