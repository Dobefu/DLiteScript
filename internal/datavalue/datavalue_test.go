package datavalue

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestDatavalueNull(t *testing.T) {
	t.Parallel()

	value := Null()

	if value.DataType != datatype.DataTypeNull {
		t.Errorf("expected DataTypeNull, got '%v'", value.DataType)
	}

	if value.ToString() != "null" {
		t.Errorf("expected 'null', got '%s'", value.ToString())
	}

	_, err := value.AsFunction()

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDatavalueNumber(t *testing.T) {
	t.Parallel()

	value := Number(1)

	if value.DataType != datatype.DataTypeNumber {
		t.Errorf("expected DataTypeNumber, got '%v'", value.DataType)
	}

	if value.ToString() != "1" {
		t.Errorf("expected '1', got '%s'", value.ToString())
	}

	num, err := value.AsNumber()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if num != 1 {
		t.Errorf("expected 1, got %f", num)
	}

	_, err = value.AsString()

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = value.AsBool()

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = value.AsTuple()

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	_, err = value.AsArray()

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDatavalueString(t *testing.T) {
	t.Parallel()

	value := String("test")

	if value.DataType != datatype.DataTypeString {
		t.Errorf("expected DataTypeString, got '%v'", value.DataType)
	}

	if value.ToString() != "test" {
		t.Errorf("expected 'test', got '%s'", value.ToString())
	}

	str, err := value.AsString()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if str != "test" {
		t.Errorf("expected 'test', got '%s'", str)
	}

	_, err = value.AsNumber()

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDatavalueBool(t *testing.T) {
	t.Parallel()

	value := Bool(true)

	if value.DataType != datatype.DataTypeBool {
		t.Errorf("expected DataTypeBool, got '%v'", value.DataType)
	}

	if value.ToString() != "true" {
		t.Errorf("expected 'true', got '%s'", value.ToString())
	}

	boolean, err := value.AsBool()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if boolean != true {
		t.Errorf("expected 'true', got '%t'", boolean)
	}

	_, err = value.AsNumber()

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDatavalueFunction(t *testing.T) {
	t.Parallel()

	value := Function(&ast.FuncDeclarationStatement{
		Name: "test",
		Args: []ast.FuncParameter{
			{Name: "a", Type: "number"},
		},
		Body: &ast.NumberLiteral{
			Value:    "1",
			StartPos: 0,
			EndPos:   3,
		},
		ReturnValues: []string{
			"number",
		},
		NumReturnValues: 1,
		StartPos:        0,
		EndPos:          3,
	})

	if value.DataType != datatype.DataTypeFunction {
		t.Errorf("expected DataTypeFunction, got '%v'", value.DataType)
	}

	if value.ToString() != "func test" {
		t.Errorf("expected 'func test', got '%s'", value.ToString())
	}

	function, err := value.AsFunction()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if function.Name != "test" {
		t.Errorf("expected 'test', got '%s'", function.Name)
	}
}

func TestDatavalueTuple(t *testing.T) {
	t.Parallel()

	value := Tuple(Number(1), String("test"))

	if value.DataType != datatype.DataTypeTuple {
		t.Errorf("expected DataTypeTuple, got '%v'", value.DataType)
	}

	if value.ToString() != "(1, test)" {
		t.Errorf("expected '(1, test)', got '%s'", value.ToString())
	}

	tuple, err := value.AsTuple()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if len(tuple) != 2 {
		t.Errorf("expected 2 values, got %d", len(tuple))
	}

	if tuple[0].ToString() != "1" {
		t.Errorf("expected '1', got '%s'", tuple[0].ToString())
	}

	if tuple[1].ToString() != "test" {
		t.Errorf("expected 'test', got '%s'", tuple[1].ToString())
	}

	_, err = value.AsNumber()

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDatavalueTupleEmpty(t *testing.T) {
	t.Parallel()

	value := Tuple()

	if value.DataType != datatype.DataTypeTuple {
		t.Errorf("expected DataTypeTuple, got '%v'", value.DataType)
	}

	if value.ToString() != "()" {
		t.Errorf("expected '()', got '%s'", value.ToString())
	}

	tuple, err := value.AsTuple()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if len(tuple) != 0 {
		t.Errorf("expected 0 values, got %d", len(tuple))
	}
}

func TestDatavalueArray(t *testing.T) {
	t.Parallel()

	value := Array(Number(1), String("test"))

	if value.DataType != datatype.DataTypeArray {
		t.Errorf("expected DataTypeArray, got '%v'", value.DataType)
	}

	if value.ToString() != "[1, test]" {
		t.Errorf("expected '[1, test]', got '%s'", value.ToString())
	}

	array, err := value.AsArray()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if len(array) != 2 {
		t.Errorf("expected 2 values, got %d", len(array))
	}

	if array[0].ToString() != "1" {
		t.Errorf("expected '1', got '%s'", array[0].ToString())
	}

	if array[1].ToString() != "test" {
		t.Errorf("expected 'test', got '%s'", array[1].ToString())
	}

	_, err = value.AsNumber()

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDatavalueArrayEmpty(t *testing.T) {
	t.Parallel()

	value := Array()

	if value.DataType != datatype.DataTypeArray {
		t.Errorf("expected DataTypeArray, got '%v'", value.DataType)
	}

	if value.ToString() != "[]" {
		t.Errorf("expected '[]', got '%s'", value.ToString())
	}

	array, err := value.AsArray()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if len(array) != 0 {
		t.Errorf("expected 0 values, got %d", len(array))
	}
}

func TestDatavalueAny(t *testing.T) {
	t.Parallel()

	value := Any(1)

	if value.DataType != datatype.DataTypeAny {
		t.Errorf(
			"expected DataType %s, got '%s'",
			datatype.DataTypeAny.AsString(),
			value.DataType.AsString(),
		)
	}

	if value.ToString() != "1" {
		t.Errorf(
			"expected '%s', got '%s'",
			"1",
			value.ToString(),
		)
	}
}

func TestDatavalueError(t *testing.T) {
	t.Parallel()

	value := Error(errors.New("test"))

	if value.DataType != datatype.DataTypeError {
		t.Errorf("expected DataTypeError, got '%v'", value.DataType)
	}

	if value.ToString() != "test" {
		t.Errorf("expected 'test', got '%s'", value.ToString())
	}

	errValue, err := value.AsError()

	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}

	if errValue.Error() != "test" {
		t.Errorf("expected 'test', got '%s'", errValue.Error())
	}

	str := Error(nil).ToString()

	if str != "null" {
		t.Errorf("expected 'null', got '%s'", str)
	}
}

func TestDatavalueAsNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected float64
	}{
		{
			name:     "float64",
			input:    Any(float64(1)),
			expected: 1,
		},
		{
			name:     "int",
			input:    Any(1),
			expected: 1,
		},
		{
			name:     "int8",
			input:    Any(int8(1)),
			expected: 1,
		},
		{
			name:     "int16",
			input:    Any(int16(1)),
			expected: 1,
		},
		{
			name:     "int32",
			input:    Any(int32(1)),
			expected: 1,
		},
		{
			name:     "int64",
			input:    Any(int64(1)),
			expected: 1,
		},
		{
			name:     "uint",
			input:    Any(uint(1)),
			expected: 1,
		},
		{
			name:     "uint8",
			input:    Any(uint8(1)),
			expected: 1,
		},
		{
			name:     "uint16",
			input:    Any(uint16(1)),
			expected: 1,
		},
		{
			name:     "uint32",
			input:    Any(uint32(1)),
			expected: 1,
		},
		{
			name:     "uint64",
			input:    Any(uint64(1)),
			expected: 1,
		},
		{
			name:     "float32",
			input:    Any(float32(1)),
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			num, err := test.input.AsNumber()

			if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			if num != test.expected {
				t.Errorf("expected %f, got %f", test.expected, num)
			}
		})
	}
}

func TestDatavalueAsNumberErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:  "nil",
			input: Any(nil),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeNumber.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
		{
			name:  "invalid type",
			input: Any("bogus"),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeNumber.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.input.AsNumber()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestDatavalueUnknown(t *testing.T) {
	t.Parallel()

	val := Value{
		DataType: datatype.DataType(-1),
		Num:      0,
		Str:      "",
		Bool:     false,
		Func:     nil,
		Values:   nil,
		Error:    nil,
		Any:      nil,
	}

	if val.DataType != datatype.DataType(-1) {
		t.Errorf("expected DataType(-1), got '%v'", val.DataType)
	}

	if val.ToString() != errorutil.ErrorMsgTypeUnknownDataType {
		t.Errorf("expected '%s', got '%s'", errorutil.ErrorMsgTypeUnknownDataType, val.ToString())
	}
}

func TestDatavalueAsString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:     "string",
			input:    String("test"),
			expected: "test",
		},
		{
			name:     "any",
			input:    Any("test"),
			expected: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			str, err := test.input.AsString()

			if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			if str != test.expected {
				t.Errorf("expected '%s', got '%s'", test.expected, str)
			}
		})
	}
}

func TestDatavalueAsStringErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:  "nil",
			input: Any(nil),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeString.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
		{
			name:  "invalid type",
			input: Any(1),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeString.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.input.AsString()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestDatavalueAsBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected bool
	}{
		{
			name:     "bool",
			input:    Bool(true),
			expected: true,
		},
		{
			name:     "any",
			input:    Any(true),
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			boolean, err := test.input.AsBool()

			if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			if boolean != test.expected {
				t.Errorf("expected \"%t\", got \"%t\"", test.expected, boolean)
			}
		})
	}
}

func TestDatavalueAsBoolErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:  "nil",
			input: Any(nil),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeBool.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
		{
			name:  "invalid type",
			input: Any(1),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeBool.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.input.AsBool()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestDatavalueAsFunction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected *ast.FuncDeclarationStatement
	}{
		{
			name: "function",
			input: Function(&ast.FuncDeclarationStatement{ //nolint:exhaustruct
				Name: "test",
			}),
			expected: &ast.FuncDeclarationStatement{ //nolint:exhaustruct
				Name: "test",
			},
		},
		{
			name: "any",
			input: Any(&ast.FuncDeclarationStatement{ //nolint:exhaustruct
				Name: "test",
			}),
			expected: &ast.FuncDeclarationStatement{ //nolint:exhaustruct
				Name: "test",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			function, err := test.input.AsFunction()

			if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			if function.Name != test.expected.Name {
				t.Errorf("expected '%s', got '%s'", test.expected.Name, function.Name)
			}
		})
	}
}

func TestDatavalueAsFunctionErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:  "nil",
			input: Any(nil),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeFunction.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
		{
			name:  "invalid type",
			input: Any(1),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeFunction.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.input.AsFunction()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestDatavalueAsArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected []Value
	}{
		{
			name:  "array",
			input: Array(Number(1), String("test")),
			expected: []Value{
				Number(1),
				String("test"),
			},
		},
		{
			name:  "any",
			input: Any([]Value{Number(1), String("test")}),
			expected: []Value{
				Number(1),
				String("test"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			array, err := test.input.AsArray()

			if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			if len(array) != len(test.expected) {
				t.Errorf("expected %d values, got %d", len(test.expected), len(array))
			}

			for i, val := range array {
				if val.ToString() != test.expected[i].ToString() {
					t.Errorf("expected '%s', got '%s'", test.expected[i].ToString(), val.ToString())
				}
			}
		})
	}
}

func TestDatavalueAsArrayErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:  "nil",
			input: Any(nil),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeArray.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
		{
			name:  "invalid type",
			input: Any(1),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeArray.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.input.AsArray()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestDatavalueAsTuple(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected []Value
	}{
		{
			name:  "tuple",
			input: Tuple(Number(1), String("test")),
			expected: []Value{
				Number(1),
				String("test"),
			},
		},
		{
			name:  "any",
			input: Any([]Value{Number(1), String("test")}),
			expected: []Value{
				Number(1),
				String("test"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			tuple, err := test.input.AsTuple()

			if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			if len(tuple) != len(test.expected) {
				t.Errorf("expected %d values, got %d", len(test.expected), len(tuple))
			}

			for i, val := range tuple {
				if val.ToString() != test.expected[i].ToString() {
					t.Errorf("expected '%s', got '%s'", test.expected[i].ToString(), val.ToString())
				}
			}
		})
	}
}

func TestDatavalueAsTupleErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:  "nil",
			input: Any(nil),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeTuple.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
		{
			name:  "invalid type",
			input: Any(1),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeTuple.AsString(),
					datatype.DataTypeAny.AsString(),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
		})

		_, err := test.input.AsTuple()

		if err == nil {
			t.Fatalf("expected error, got nil")
		}

	}
}

func TestDatavalueAsErrorErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Value
		expected string
	}{
		{
			name:  "null value",
			input: Null(),
			expected: fmt.Sprintf(
				"%s: %s",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(
					errorutil.ErrorMsgTypeExpected,
					datatype.DataTypeError.AsString(),
					datatype.DataTypeNull.AsString(),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.input.AsError()

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, err.Error())
			}
		})
	}
}

func TestDatavalueEquals(t *testing.T) {
	t.Parallel()

	function := &ast.FuncDeclarationStatement{
		Name: "test",
		Args: []ast.FuncParameter{
			{Name: "a", Type: "number"},
		},
		Body: &ast.NumberLiteral{
			Value:    "1",
			StartPos: 0,
			EndPos:   3,
		},
		ReturnValues: []string{
			"number",
		},
		NumReturnValues: 1,
		StartPos:        0,
		EndPos:          3,
	}

	tests := []struct {
		name        string
		value       Value
		other       Value
		shouldMatch bool
	}{
		{
			name:        "two numbers",
			value:       Number(1),
			other:       Number(1),
			shouldMatch: true,
		},
		{
			name:        "two strings",
			value:       String(""),
			other:       String(""),
			shouldMatch: true,
		},
		{
			name:        "two booleans",
			value:       Bool(false),
			other:       Bool(false),
			shouldMatch: true,
		},
		{
			name:        "two nulls",
			value:       Null(),
			other:       Null(),
			shouldMatch: true,
		},
		{
			name:        "two functions",
			value:       Function(function),
			other:       Function(function),
			shouldMatch: true,
		},
		{
			name:        "two tuples with same values",
			value:       Tuple(Number(1), String("test")),
			other:       Tuple(Number(1), String("test")),
			shouldMatch: true,
		},
		{
			name:        "two tuples with different number of values",
			value:       Tuple(Number(1), String("test")),
			other:       Tuple(Number(1)),
			shouldMatch: false,
		},
		{
			name:        "two tuples with different values",
			value:       Tuple(Number(1), String("test")),
			other:       Tuple(Number(1), String("tes")),
			shouldMatch: false,
		},

		{
			name:        "different types",
			value:       Number(1),
			other:       Null(),
			shouldMatch: false,
		},
		{
			name:        "two errors with same error",
			value:       Error(errors.New("test")),
			other:       Error(errors.New("test")),
			shouldMatch: true,
		},
		{
			name:        "two errors with left nil",
			value:       Error(nil),
			other:       Error(errors.New("test2")),
			shouldMatch: false,
		},
		{
			name:        "two errors with right nil",
			value:       Error(errors.New("test")),
			other:       Error(nil),
			shouldMatch: false,
		},
		{
			name:        "any type",
			value:       Any(1),
			other:       Any(1),
			shouldMatch: true,
		},
		{
			name:        "unknown type",
			value:       Value{DataType: datatype.DataType(-1)}, //nolint:exhaustruct
			other:       Value{DataType: datatype.DataType(-1)}, //nolint:exhaustruct
			shouldMatch: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.shouldMatch && !test.value.Equals(test.other) {
				t.Errorf("expected true, got false")
			}

			if !test.shouldMatch && test.value.Equals(test.other) {
				t.Errorf("expected false, got true")
			}
		})
	}
}
