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
	}
}

// AsNumber returns the value as a number.
func (v Value) AsNumber() (float64, error) {
	if v.dataType != datatype.DataTypeNumber {
		return 0, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeNumber.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Num, nil
}

// AsString returns the value as a string.
func (v Value) AsString() (string, error) {
	if v.dataType != datatype.DataTypeString {
		return "", errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeString.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Str, nil
}

// AsBool returns the value as a boolean.
func (v Value) AsBool() (bool, error) {
	if v.dataType != datatype.DataTypeBool {
		return false, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeBool.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Bool, nil
}

// AsFunction returns the value as a function.
func (v Value) AsFunction() (*ast.FuncDeclarationStatement, error) {
	if v.dataType != datatype.DataTypeFunction {
		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeFunction.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Func, nil
}

// AsArray returns the value as an array.
func (v Value) AsArray() ([]Value, error) {
	if v.dataType != datatype.DataTypeArray {
		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeArray.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Values, nil
}

// AsTuple returns the value as a tuple.
func (v Value) AsTuple() ([]Value, error) {
	if v.dataType != datatype.DataTypeTuple {
		return nil, errorutil.NewError(
			errorutil.StageEvaluate,
			errorutil.ErrorMsgTypeExpected,
			datatype.DataTypeTuple.AsString(),
			v.dataType.AsString(),
		)
	}

	return v.Values, nil
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

	default:
		return false
	}
}
