// Package function provides the Function type and related methods.
package function

import (
	"fmt"
	"strings"

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

// ArgInfo defines the information for a function argument or return value.
type ArgInfo struct {
	Name        string
	Type        datatype.DataType
	Description string
}

// DeprecationInfo defines the information for a deprecation.
type DeprecationInfo struct {
	IsDeprecated bool
	Description  string
	Version      string
}

// Info defines the information for a function.
type Info struct {
	Name            string
	Description     string
	PackageName     string
	Handler         Handler
	FunctionType    Type
	Parameters      []ArgInfo
	ReturnValues    []ArgInfo
	IsBuiltin       bool
	Since           string
	DeprecationInfo DeprecationInfo
	Examples        []string
}

// MakeFunction creates a new function definition.
func MakeFunction(
	name string,
	description string,
	packageName string,
	functionType Type,
	parameters []ArgInfo,
	returnValues []ArgInfo,
	isBuiltin bool,
	since string,
	deprecationInfo DeprecationInfo,
	examples []string,
	impl func(e EvaluatorInterface, args []datavalue.Value) datavalue.Value,
) Info {
	handler := func(
		e EvaluatorInterface,
		args []datavalue.Value,
	) (datavalue.Value, error) {
		return impl(e, args), nil
	}

	return Info{
		Name:            name,
		Description:     description,
		PackageName:     packageName,
		Handler:         handler,
		FunctionType:    functionType,
		Parameters:      parameters,
		ReturnValues:    returnValues,
		IsBuiltin:       isBuiltin,
		Since:           since,
		DeprecationInfo: deprecationInfo,
		Examples:        examples,
	}
}

// Expr returns the function signature as a string.
func (f *Info) Expr() string {
	params := make([]string, len(f.Parameters))

	for i, param := range f.Parameters {
		paramStr := param.Name

		if param.Type != 0 {
			paramStr = fmt.Sprintf("%s %s", param.Name, param.Type.AsString())
		}

		params[i] = paramStr
	}

	returns := make([]string, len(f.ReturnValues))

	for i, ret := range f.ReturnValues {
		returns[i] = ret.Type.AsString()
	}

	signature := fmt.Sprintf("func %s(%s)", f.Name, strings.Join(params, ", "))

	if len(returns) == 0 {
		return signature
	}

	if len(returns) == 1 {
		return fmt.Sprintf("%s %s", signature, returns[0])
	}

	return fmt.Sprintf("%s (%s)", signature, strings.Join(returns, ", "))
}
