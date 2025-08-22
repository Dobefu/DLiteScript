package datavalue

import (
	"testing"

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

func TestDatavalueUnknown(t *testing.T) {
	t.Parallel()

	val := Value{
		dataType: datatype.DataType(-1),
		Num:      0,
		Str:      "",
		Bool:     false,
	}

	if val.DataType() != datatype.DataType(-1) {
		t.Errorf("expected DataType(-1), got '%v'", val.DataType())
	}

	if val.ToString() != errorutil.ErrorMsgTypeUnknownDataType {
		t.Errorf("expected '%s', got '%s'", errorutil.ErrorMsgTypeUnknownDataType, val.ToString())
	}
}
