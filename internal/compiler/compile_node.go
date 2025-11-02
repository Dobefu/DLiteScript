package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileNode(node ast.ExprNode) error {
	switch n := node.(type) {
	case *ast.NumberLiteral:
		return c.compileNumberLiteral(n)

	case *ast.BinaryExpr:
		return c.compileBinaryExpr(n)

	default:
		return nil
	}
}
