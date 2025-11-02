package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileNode(node ast.ExprNode) error {
	switch n := node.(type) {
	case *ast.NumberLiteral:
		return c.compileNumberLiteral(n)

	case *ast.StringLiteral:
		return c.compileStringLiteral(n)

	case *ast.BinaryExpr:
		return c.compileBinaryExpr(n)

	case *ast.StatementList:
		return c.compileStatementList(n)

	case *ast.FunctionCall:
		return c.compileFunctionCall(n)

	case *ast.CommentLiteral, *ast.NewlineLiteral:
		return nil

	default:
		return nil
	}
}
