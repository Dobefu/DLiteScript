package ast

import (
	"fmt"
)

// VariableDeclaration represents a variable declaration.
type VariableDeclaration struct {
	Name     string
	Type     string
	Value    ExprNode
	StartPos int
	EndPos   int
}

// Expr returns the expression of the variable declaration.
func (v *VariableDeclaration) Expr() string {
	if v.Value == nil {
		return fmt.Sprintf("var %s %s", v.Name, v.Type)
	}

	return fmt.Sprintf("var %s %s = %s", v.Name, v.Type, v.Value.Expr())
}

// StartPosition returns the start position of the variable declaration.
func (v *VariableDeclaration) StartPosition() int {
	return v.StartPos
}

// EndPosition returns the end position of the variable declaration.
func (v *VariableDeclaration) EndPosition() int {
	return v.EndPos
}

// Walk walks the variable declaration and its value.
func (v *VariableDeclaration) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(v)

	if !shouldContinue {
		return
	}

	if v.Value != nil {
		shouldContinue = fn(v.Value)

		if !shouldContinue {
			return
		}

		v.Value.Walk(fn)
	}
}
