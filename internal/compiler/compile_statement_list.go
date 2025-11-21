package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileStatementList(s *ast.StatementList) error {
	for _, stmt := range s.Statements {
		funcDecl, hasFuncDecl := stmt.(*ast.FuncDeclarationStatement)

		if !hasFuncDecl {
			continue
		}

		count := funcDecl.NumReturnValues

		if count == 0 && len(funcDecl.ReturnValues) > 0 {
			count = len(funcDecl.ReturnValues)
		}

		c.functionReturnCounts[funcDecl.Name] = count
	}

	for _, stmt := range s.Statements {
		err := c.compileNode(stmt)

		if err != nil {
			return err
		}
	}

	return nil
}
