package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileStatementList(s *ast.StatementList) error {
	for _, stmt := range s.Statements {
		err := c.compileNode(stmt)

		if err != nil {
			return err
		}
	}

	return nil
}
