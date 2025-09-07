// Package function provides the Function type and related methods.
package function

import (
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

// EvaluatorInterface defines the interface that functions need from the evaluator.
type EvaluatorInterface interface {
	AddToBuffer(format string, args ...any)
}

// Type defines the type of function.
type Type int

const (
	// FunctionTypeFixed defines function with a fixed number of arguments.
	FunctionTypeFixed Type = iota
	// FunctionTypeVariadic defines function with a variadic number of arguments.
	FunctionTypeVariadic
	// FunctionTypeMixedVariadic defines function with a mixed number of arguments.
	FunctionTypeMixedVariadic
)

// Handler defines the handler for a function.
type Handler func(
	e EvaluatorInterface,
	args []datavalue.Value,
) (datavalue.Value, error)

// PackageInfo defines the information for a package.
type PackageInfo struct {
	Functions map[string]Info
}

// Info defines the information for a function.
type Info struct {
	Handler      Handler
	FunctionType Type
	ArgKinds     []datatype.DataType
}

// MakeFunction creates a new function definition.
func MakeFunction(
	functionType Type,
	argKinds []datatype.DataType,
	impl func(e EvaluatorInterface, args []datavalue.Value) datavalue.Value,
) Info {
	handler := func(
		e EvaluatorInterface,
		args []datavalue.Value,
	) (datavalue.Value, error) {
		return impl(e, args), nil
	}

	return Info{
		Handler:      handler,
		FunctionType: functionType,
		ArgKinds:     argKinds,
	}
}
