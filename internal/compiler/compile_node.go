package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileNode(node ast.ExprNode) error {
	switch n := node.(type) {
	case *ast.NumberLiteral:
		return c.compileNumberLiteral(n)

	case *ast.BoolLiteral:
		return c.compileBoolLiteral(n)

	case *ast.NullLiteral:
		return c.compileNullLiteral(n)

	case *ast.StringLiteral:
		return c.compileStringLiteral(n)

	case *ast.BinaryExpr:
		return c.compileBinaryExpr(n)

	case *ast.PrefixExpr:
		return c.compilePrefixExpr(n)

	case *ast.SpreadExpr:
		return c.compileSpreadExpr(n)

	case *ast.StatementList:
		return c.compileStatementList(n)

	case *ast.FunctionCall:
		return c.compileFunctionCall(n)

	case *ast.VariableDeclaration:
		return c.compileVariableDeclaration(n)

	case *ast.ConstantDeclaration:
		return c.compileConstantDeclaration(n)

	case *ast.Identifier:
		return c.compileIdentifier(n)

	case *ast.AssignmentStatement:
		return c.compileAssignmentStatement(n)

	case *ast.BlockStatement:
		return c.compileBlockStatement(n)

	case *ast.IfStatement:
		return c.compileIfStatement(n)

	case *ast.ForStatement:
		return c.compileForStatement(n)

	case *ast.BreakStatement:
		return c.compileBreakStatement(n)

	case *ast.ContinueStatement:
		return c.compileContinueStatement(n)

	case *ast.FuncDeclarationStatement:
		return c.compileFuncDeclarationStatement(n)

	case *ast.ReturnStatement:
		return c.compileReturnStatement(n)

	case *ast.CommentLiteral, *ast.NewlineLiteral:
		return nil

	default:
		return nil
	}
}
