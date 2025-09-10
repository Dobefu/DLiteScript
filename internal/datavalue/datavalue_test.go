package datavalue

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestDatavalueNull(t *testing.T) {
	t.Parallel()

	value := Null()

	if value.DataType() != datatype.DataTypeNull {
		t.Errorf("expected DataTypeNull, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeNumber {
		t.Errorf("expected DataTypeNumber, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeString {
		t.Errorf("expected DataTypeString, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeBool {
		t.Errorf("expected DataTypeBool, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeFunction {
		t.Errorf("expected DataTypeFunction, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeTuple {
		t.Errorf("expected DataTypeTuple, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeTuple {
		t.Errorf("expected DataTypeTuple, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeArray {
		t.Errorf("expected DataTypeArray, got '%v'", value.DataType())
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

	if value.DataType() != datatype.DataTypeArray {
		t.Errorf("expected DataTypeArray, got '%v'", value.DataType())
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

func TestDatavalueUnknown(t *testing.T) {
	t.Parallel()

	val := Value{
		dataType: datatype.DataType(-1),
		Num:      0,
		Str:      "",
		Bool:     false,
		Func:     nil,
		Values:   nil,
	}

	if val.DataType() != datatype.DataType(-1) {
		t.Errorf("expected DataType(-1), got '%v'", val.DataType())
	}

	if val.ToString() != errorutil.ErrorMsgTypeUnknownDataType {
		t.Errorf("expected '%s', got '%s'", errorutil.ErrorMsgTypeUnknownDataType, val.ToString())
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
			name:        "unknown type",
			value:       Value{dataType: datatype.DataType(-1)}, //nolint:exhaustruct
			other:       Value{dataType: datatype.DataType(-1)}, //nolint:exhaustruct
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
