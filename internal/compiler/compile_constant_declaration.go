package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileConstantDeclaration(cd *ast.ConstantDeclaration) error {
	return c.compileVariableDeclaration(&ast.VariableDeclaration{
		Name:  cd.Name,
		Type:  cd.Type,
		Value: cd.Value,
		Range: cd.Range,
	})
}
