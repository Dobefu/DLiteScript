package compiler

import (
	"reflect"
	"testing"

	vm "github.com/Dobefu/vee-em"
)

func TestEmitLoadRegister(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		src      byte
		addrReg  byte
		expected []byte
	}{
		{
			name:     "load register",
			src:      1,
			addrReg:  0,
			expected: []byte{byte(vm.OpcodeLoadRegister), 1, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewCompiler()
			err := c.emitLoadRegister(test.src, test.addrReg)

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
