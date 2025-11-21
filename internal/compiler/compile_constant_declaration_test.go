package compiler

import (
	"reflect"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	vm "github.com/Dobefu/vee-em"
)

func TestCompileConstantDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.ConstantDeclaration
		expected []byte
	}{
		{
			name: "constant declaration with value",
			input: &ast.ConstantDeclaration{
				Name: "test",
				Type: "string",
				Value: &ast.StringLiteral{
					Value: "test",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 1, Column: 1},
						End:   ast.Position{Offset: 0, Line: 1, Column: 1},
					},
				},
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 1, Column: 1},
					End:   ast.Position{Offset: 0, Line: 1, Column: 1},
				},
			},
			expected: []byte{
				byte(vm.OpcodeLoadImmediate), 0, 0, 0, 0, 0, 0, 0, 0, 3,
				byte(vm.OpcodeLoadImmediate), 1, 0, 0, 0, 0, 0, 0, 0, 0,
				byte(vm.OpcodeStoreMemory), 0, 1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			c := NewCompiler()
			err := c.compileConstantDeclaration(test.input)

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
