package compiler

import (
	"reflect"
	"testing"

	vm "github.com/Dobefu/vee-em"
)

func TestEmitLoadImmediate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		src      byte
		value    int64
		expected []byte
	}{
		{
			name:     "load immediate",
			src:      1,
			value:    0,
			expected: []byte{byte(vm.OpcodeLoadImmediate), 1, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewCompiler()
			err := c.emitLoadImmediate(test.src, test.value)

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
