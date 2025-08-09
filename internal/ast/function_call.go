package ast

import (
	"fmt"
	"strings"
)

// FunctionCall defines a struct for a function call.
type FunctionCall struct {
	FunctionName string
	Arguments    []ExprNode
	Pos          int
}

// Expr returns the expression of the function call.
func (fc *FunctionCall) Expr() string {
	var args strings.Builder

	for i, arg := range fc.Arguments {
		args.WriteString(arg.Expr())

		if i < len(fc.Arguments)-1 {
			args.WriteString(", ")
		}
	}

	return fmt.Sprintf("%s(%s)", fc.FunctionName, args.String())
}

// Position returns the position of the function call.
func (fc *FunctionCall) Position() int {
	return fc.Pos
}
