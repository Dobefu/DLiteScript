package evaluator

import (
	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func (e *Evaluator) evaluateFunctionCall(
	fc *ast.FunctionCall,
) (float64, error) {
	function, ok := functionRegistry[fc.FunctionName]

	if !ok {
		return 0, errorutil.NewErrorAt(
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
		return 0, err
	}

	return function.handler(argValues)
}

func (e *Evaluator) evaluateArguments(
	args []ast.ExprNode,
	expectedCount int,
	functionName string,
	fc *ast.FunctionCall,
) ([]float64, error) {
	argValues := make([]float64, len(args))

	for i, arg := range args {
		val, err := e.Evaluate(arg)

		if err != nil {
			return nil, err
		}

		argValues[i] = val
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

	return argValues, nil
}
