// Package datavalue provides the Value type and related methods.
package datavalue

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// Value represents a value in the language.
type Value struct {
	dataType datatype.DataType

	Num    float64
	Str    string
	Bool   bool
	Func   *ast.FuncDeclarationStatement
	Values []Value
	Any    any
}

// DataType returns the data type of the value.
func (v Value) DataType() datatype.DataType {
	return v.dataType
}

// Null creates a new null value.
func Null() Value {
	return Value{
		dataType: datatype.DataTypeNull,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: nil,
		Any:    nil,
	}
}

// ToString returns the string representation of the value.
func (v Value) ToString() string {
	switch v.dataType {
	case
		datatype.DataTypeNull:
		return "null"

	case
		datatype.DataTypeNumber:
		return strconv.FormatFloat(v.Num, 'f', -1, 64)

	case
		datatype.DataTypeString:
		return v.Str

	case
		datatype.DataTypeBool:
		return strconv.FormatBool(v.Bool)

	case
		datatype.DataTypeFunction:
		return fmt.Sprintf("func %s", v.Func.Name)

	case
		datatype.DataTypeTuple:
		if len(v.Values) == 0 {
			return "()"
		}

		valueStrings := make([]string, len(v.Values))

		for i, val := range v.Values {
			valueStrings[i] = val.ToString()
		}

		return fmt.Sprintf("(%s)", strings.Join(valueStrings, ", "))

	case
		datatype.DataTypeArray:
		if len(v.Values) == 0 {
			return "[]"
		}

		valueStrings := make([]string, len(v.Values))

		for i, val := range v.Values {
			valueStrings[i] = val.ToString()
		}

		return fmt.Sprintf("[%s]", strings.Join(valueStrings, ", "))

	case
		datatype.DataTypeAny:
		return "any"

	default:
		return errorutil.ErrorMsgTypeUnknownDataType
	}
}

// Number creates a new number value.
func Number(n float64) Value {
	return Value{
		dataType: datatype.DataTypeNumber,

		Num:    n,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: nil,
		Any:    nil,
	}
}

// String creates a new string value.
func String(s string) Value {
	return Value{
		dataType: datatype.DataTypeString,

		Num:    0,
		Str:    s,
		Bool:   false,
		Func:   nil,
		Values: nil,
		Any:    nil,
	}
}

// Bool creates a new boolean value.
func Bool(b bool) Value {
	return Value{
		dataType: datatype.DataTypeBool,

		Num:    0,
		Str:    "",
		Bool:   b,
		Func:   nil,
		Values: nil,
		Any:    nil,
	}
}

// Function creates a new function value.
func Function(fn *ast.FuncDeclarationStatement) Value {
	return Value{
		dataType: datatype.DataTypeFunction,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   fn,
		Values: nil,
		Any:    nil,
	}
}

// Tuple creates a new tuple value.
func Tuple(values ...Value) Value {
	return Value{
		dataType: datatype.DataTypeTuple,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: values,
		Any:    nil,
	}
}

// Array creates a new array value.
func Array(values ...Value) Value {
	return Value{
		dataType: datatype.DataTypeArray,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: values,
		Any:    nil,
	}
}

// Any creates a new any value.
func Any(a any) Value {
	return Value{
		dataType: datatype.DataTypeAny,

		Num:    0,
		Str:    "",
		Bool:   false,
		Func:   nil,
		Values: nil,
		Any:    a,
	}
}

// AsNumber returns the value as a number.
func (v Value) AsNumber() (float64, error) {
	if v.dataType == datatype.DataTypeNumber {
		return v.Num, nil
	}

	if v.dataType == datatype.DataTypeAny {
		if v.Any == nil {
			return 0, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeNumber.AsString(),
				v.dataType.AsString(),
			)
		}

		switch a := v.Any.(type) {
		case float64:
			return a, nil

		case int:
			return float64(a), nil

		case int8:
			return float64(a), nil

		case int16:
			return float64(a), nil

		case int32:
			return float64(a), nil

		case int64:
			return float64(a), nil

		case uint:
			return float64(a), nil

		case uint8:
			return float64(a), nil

		case uint16:
			return float64(a), nil

		case uint32:
			return float64(a), nil

		case uint64:
			return float64(a), nil

		case float32:
			return float64(a), nil

		default:
			return 0, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeNumber.AsString(),
				v.dataType.AsString(),
			)
		}
	}

	return 0, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeNumber.AsString(),
		v.dataType.AsString(),
	)
}

// AsString returns the value as a string.
func (v Value) AsString() (string, error) {
	if v.dataType == datatype.DataTypeString {
		return v.Str, nil
	}

	if v.dataType == datatype.DataTypeAny {
		if v.Any == nil {
			return "", errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeString.AsString(),
				v.dataType.AsString(),
			)
		}

		str, isString := v.Any.(string)

		if isString {
			return str, nil
		}

		return "", errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeString.AsString(),
			v.dataType.AsString(),
		)
	}

	return "", errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeString.AsString(),
		v.dataType.AsString(),
	)
}

// AsBool returns the value as a boolean.
func (v Value) AsBool() (bool, error) {
	if v.dataType == datatype.DataTypeBool {
		return v.Bool, nil
	}

	if v.dataType == datatype.DataTypeAny {
		if v.Any == nil {
			return false, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeBool.AsString(),
				v.dataType.AsString(),
			)
		}

		boolean, isBoolean := v.Any.(bool)

		if isBoolean {
			return boolean, nil
		}

		return false, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeBool.AsString(),
			v.dataType.AsString(),
		)
	}

	return false, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeBool.AsString(),
		v.dataType.AsString(),
	)
}

// AsFunction returns the value as a function.
func (v Value) AsFunction() (*ast.FuncDeclarationStatement, error) {
	if v.dataType == datatype.DataTypeFunction {
		return v.Func, nil
	}

	if v.dataType == datatype.DataTypeAny {
		if v.Any == nil {
			return nil, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeFunction.AsString(),
				v.dataType.AsString(),
			)
		}

		function, isFunction := v.Any.(*ast.FuncDeclarationStatement)

		if isFunction {
			return function, nil
		}

		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeFunction.AsString(),
			v.dataType.AsString(),
		)
	}

	return nil, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeFunction.AsString(),
		v.dataType.AsString(),
	)
}

// AsArray returns the value as an array.
func (v Value) AsArray() ([]Value, error) {
	if v.dataType == datatype.DataTypeArray {
		return v.Values, nil
	}

	if v.dataType == datatype.DataTypeAny {
		if v.Any == nil {
			return nil, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeArray.AsString(),
				v.dataType.AsString(),
			)
		}

		array, isArray := v.Any.([]Value)

		if isArray {
			return array, nil
		}

		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeArray.AsString(),
			v.dataType.AsString(),
		)
	}

	return nil, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeArray.AsString(),
		v.dataType.AsString(),
	)
}

// AsTuple returns the value as a tuple.
func (v Value) AsTuple() ([]Value, error) {
	if v.dataType == datatype.DataTypeTuple {
		return v.Values, nil
	}

	if v.dataType == datatype.DataTypeAny {
		if v.Any == nil {
			return nil, errorutil.NewError(
				errorutil.StageEvaluate,
				errorutil.ErrorMsgTypeExpected,
				datatype.DataTypeTuple.AsString(),
				v.dataType.AsString(),
			)
		}

		tuple, isTuple := v.Any.([]Value)

		if isTuple {
			return tuple, nil
		}

		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeTuple.AsString(),
			v.dataType.AsString(),
		)
	}

	return nil, errorutil.NewError(
		errorutil.StageEvaluate,
		errorutil.ErrorMsgTypeExpected,
		datatype.DataTypeTuple.AsString(),
		v.dataType.AsString(),
	)
}

// Equals returns if two values are equal.
func (v Value) Equals(other Value) bool {
	if v.dataType != other.dataType {
		return false
	}

	switch v.dataType {
	case
		datatype.DataTypeNull:
		return true

	case
		datatype.DataTypeNumber:
		return v.Num == other.Num

	case
		datatype.DataTypeString:
		return v.Str == other.Str

	case
		datatype.DataTypeBool:
		return v.Bool == other.Bool

	case
		datatype.DataTypeFunction:
		return v.Func == other.Func

	case
		datatype.DataTypeTuple,
		datatype.DataTypeArray:
		if len(v.Values) != len(other.Values) {
			return false
		}

		for i, val := range v.Values {
			if !val.Equals(other.Values[i]) {
				return false
			}
		}

		return true

	case
		datatype.DataTypeAny:
		return true

	default:
		return false
	}
}
