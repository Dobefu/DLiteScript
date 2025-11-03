package compiler

import (
	"reflect"
	"testing"

	vm "github.com/Dobefu/vee-em"
)

func TestEmitLoadConstPoolIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		src      byte
		index    int
		expected []byte
	}{
		{
			name:     "load const pool index",
			src:      1,
			index:    0,
			expected: []byte{byte(vm.OpcodeLoadImmediate), 1, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewCompiler()
			err := c.emitLoadConstPoolIndex(test.src, test.index)

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
