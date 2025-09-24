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
	Range        Range
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

// GetRange returns the range of the function call.
func (fc *FunctionCall) GetRange() Range {
	return fc.Range
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
