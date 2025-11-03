package compiler

import (
	"reflect"
	"testing"

	vm "github.com/Dobefu/vee-em"
)

func TestEmitFunctionCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		functionIndex int
		firstArgReg   byte
		argCount      byte
		expected      []byte
	}{
		{
			name:          "emit function call",
			functionIndex: 0,
			firstArgReg:   1,
			argCount:      1,
			expected:      []byte{byte(vm.OpcodeHostCall), 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewCompiler()
			err := c.emitFunctionCall(
				test.functionIndex,
				test.firstArgReg,
				test.argCount,
			)

			if err != nil {
				t.Fatalf("Expected no error, got: \"%s\"", err.Error())
			}

			if !reflect.DeepEqual(c.bytecode, test.expected) {
				t.Fatalf(
					"expected bytecode to be %v, got %v",
					test.expected,
					c.bytecode,
				)
			}
		})
	}
}
