package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateFunctionCall(
	fc *ast.FunctionCall,
) (datavalue.Value, error) {
	function, ok := functionRegistry[fc.FunctionName]

	if !ok {
		return datavalue.Null(), errorutil.NewErrorAt(
			errorutil.ErrorMsgUndefinedFunction,
			fc.Position(),
			fc.FunctionName,
		)
	}

	argValues, err := e.evaluateArguments(
		fc.Arguments,
		function.argCount,
		fc.FunctionName,
		fc,
	)

	if err != nil {
		return datavalue.Null(), err
	}

	return function.handler(argValues)
}

func (e *Evaluator) evaluateArguments(
	args []ast.ExprNode,
	expectedCount int,
	functionName string,
	fc *ast.FunctionCall,
) ([]datavalue.Value, error) {
	argValues := make([]datavalue.Value, len(args))

	for i, arg := range args {
		val, err := e.Evaluate(arg)

		if err != nil {
			return nil, err
		}

		argValues[i] = val
	}

	function, hasFunction := functionRegistry[functionName]

	if !hasFunction {
		return argValues, nil
	}

	if expectedCount > 0 && len(argValues) != expectedCount {
		return nil, errorutil.NewErrorAt(
			errorutil.ErrorMsgFunctionNumArgs,
			fc.Position(),
			functionName,
			expectedCount,
			len(argValues),
		)
	}

	for i, expectedKind := range function.argKinds {
		if argValues[i].DataType() != expectedKind {
			return nil, errorutil.NewErrorAt(
				errorutil.ErrorMsgFunctionArgType,
				fc.Position(),
				functionName,
				i+1,
				expectedKind.AsString(),
				argValues[i].DataType().AsString(),
			)
		}
	}

	return argValues, nil
}
