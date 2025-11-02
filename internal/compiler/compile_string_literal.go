package compiler

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
)

func (c *Compiler) compileStringLiteral(s *ast.StringLiteral) error {
	index := c.addToConstPool(s.Value)

	return c.emitLoadConstPoolIndex(c.incrementRegCounter(), index)
}

func (c *Compiler) addToConstPool(str string) int {
	index, hasIndex := c.constPoolMap[str]

	if hasIndex {
		return index
	}

	index = len(c.constPool)
	c.constPool = append(c.constPool, str)
	c.constPoolMap[str] = index

	return index
}
