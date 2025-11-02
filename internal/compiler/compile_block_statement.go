package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileBlockStatement(bs *ast.BlockStatement) error {
	c.variableScopes = append(c.variableScopes, make(map[string]uint64))

	for _, stmt := range bs.Statements {
		err := c.compileNode(stmt)

		if err != nil {
			return err
		}
	}

	c.variableScopes = c.variableScopes[:len(c.variableScopes)-1]

	return nil
}
