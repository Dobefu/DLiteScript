package compiler

import (
	"reflect"
	"testing"
)

func TestEmitAdd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		dest     byte
		src1     byte
		src2     byte
		expected []byte
	}{
		{
			name:     "add (0 + 0)",
			dest:     0,
			src1:     0,
			src2:     0,
			expected: []byte{7, 0, 0, 0},
		},
		{
			name:     "add (1 + 1)",
			dest:     0,
			src1:     1,
			src2:     1,
			expected: []byte{7, 0, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewCompiler()
			err := c.emitAdd(test.dest, test.src1, test.src2)

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
