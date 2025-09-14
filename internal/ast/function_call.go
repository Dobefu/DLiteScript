package ast

import (
	"fmt"
	"strings"
)

// FunctionCall defines a struct for a function call.
type FunctionCall struct {
	Namespace    string
	FunctionName string
	Arguments    []ExprNode
	StartPos     int
	EndPos       int
}

// Expr returns the expression of the function call.
func (fc *FunctionCall) Expr() string {
	if len(fc.Arguments) == 0 {
		return fmt.Sprintf("%s()", fc.FunctionName)
	}

	var args strings.Builder

	for i, arg := range fc.Arguments {
		if arg == nil {
			continue
		}

		args.WriteString(arg.Expr())

		if i < len(fc.Arguments)-1 {
			args.WriteString(", ")
		}
	}

	functionName := fc.FunctionName

	if fc.Namespace != "" {
		functionName = fmt.Sprintf("%s.%s", fc.Namespace, fc.FunctionName)
	}

	return fmt.Sprintf("%s(%s)", functionName, args.String())
}

// StartPosition returns the start position of the function call.
func (fc *FunctionCall) StartPosition() int {
	return fc.StartPos
}

// EndPosition returns the end position of the function call.
func (fc *FunctionCall) EndPosition() int {
	return fc.EndPos
}

// Walk walks the function call and its arguments.
func (fc *FunctionCall) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(fc)

	if !shouldContinue {
		return
	}

	for _, arg := range fc.Arguments {
		shouldContinue = fn(arg)

		if !shouldContinue {
			return
		}

		arg.Walk(fn)
	}
}
