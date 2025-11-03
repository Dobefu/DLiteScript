package compiler

import (
	"reflect"
	"testing"

	vm "github.com/Dobefu/vee-em"
)

func TestEmitDiv(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		dest     byte
		src1     byte
		src2     byte
		expected []byte
	}{
		{
			name:     "div (1 / 1)",
			dest:     0,
			src1:     0,
			src2:     0,
			expected: []byte{byte(vm.OpcodeDiv), 0, 0, 0},
		},
		{
			name:     "div (2 / 2)",
			dest:     0,
			src1:     1,
			src2:     1,
			expected: []byte{byte(vm.OpcodeDiv), 0, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewCompiler()
			err := c.emitDiv(test.dest, test.src1, test.src2)

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
