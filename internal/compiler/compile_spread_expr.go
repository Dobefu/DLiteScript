package compiler

import "github.com/Dobefu/DLiteScript/internal/ast"

func (c *Compiler) compileSpreadExpr(node *ast.SpreadExpr) error {
	return c.compileNode(node.Expression)
}
