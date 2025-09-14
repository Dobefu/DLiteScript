package function

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestMakeFunction(t *testing.T) {
	t.Parallel()

	function := MakeFunction(
		Documentation{
			Name:        "test",
			Description: "description",
			Since:       "v1.0.0",
			DeprecationInfo: DeprecationInfo{
				IsDeprecated: false,
				Description:  "",
				Version:      "",
			},
			Examples: []string{"example"},
		},
		"package",
		FunctionTypeFixed,
		[]ArgInfo{
			{
				Name:        "param",
				Type:        datatype.DataTypeNumber,
				Description: "test",
			},
		},
		[]ArgInfo{
			{
				Name:        "return",
				Type:        datatype.DataTypeNumber,
				Description: "test",
			},
		},
		false,
		func(_ EvaluatorInterface, args []datavalue.Value) datavalue.Value {
			num, err := args[0].AsNumber()

			if err != nil {
				return datavalue.Number(0)
			}

			return datavalue.Number(num)
		},
	)

	result, err := function.Handler(nil, []datavalue.Value{datavalue.Number(1)})

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result.DataType() != datatype.DataTypeNumber {
		t.Errorf("expected DataTypeNumber, got %v", result.DataType())
	}

	if function.Documentation.Name != "test" {
		t.Errorf("expected 'test', got '%s'", function.Documentation.Name)
	}

	if function.Documentation.Description != "description" {
		t.Errorf(
			"expected 'description', got '%s'",
			function.Documentation.Description,
		)
	}

	if function.PackageName != "package" {
		t.Errorf("expected 'package', got '%s'", function.PackageName)
	}

	if function.Documentation.Since != "v1.0.0" {
		t.Errorf("expected 'v1.0.0', got '%s'", function.Documentation.Since)
	}

	if function.Documentation.DeprecationInfo.IsDeprecated != false {
		t.Errorf(
			"expected 'false', got '%t'",
			function.Documentation.DeprecationInfo.IsDeprecated,
		)
	}

	if function.Documentation.Examples[0] != "example" {
		t.Errorf("expected 'example', got '%s'", function.Documentation.Examples[0])
	}
}

func TestExpr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		args     []ArgInfo
		returns  []ArgInfo
		expected string
	}{
		{
			name: "no return values",
			args: []ArgInfo{
				{
					Name:        "param",
					Type:        datatype.DataTypeNumber,
					Description: "test",
				},
			},
			returns:  []ArgInfo{},
			expected: "func test(param number)",
		},
		{
			name: "single return value",
			args: []ArgInfo{
				{
					Name:        "param",
					Type:        datatype.DataTypeNumber,
					Description: "test",
				},
			},
			returns: []ArgInfo{
				{
					Name:        "return",
					Type:        datatype.DataTypeNumber,
					Description: "test",
				},
			},
			expected: "func test(param number) number",
		},
		{
			name: "multiple return values",
			args: []ArgInfo{
				{
					Name:        "param",
					Type:        datatype.DataTypeNumber,
					Description: "test",
				},
			},
			returns: []ArgInfo{
				{
					Name:        "return",
					Type:        datatype.DataTypeNumber,
					Description: "test",
				},
				{
					Name:        "return",
					Type:        datatype.DataTypeString,
					Description: "test",
				},
			},
			expected: "func test(param number) (number, string)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			function := MakeFunction(
				Documentation{
					Name:        "test",
					Description: "description",
					Since:       "v1.0.0",
					DeprecationInfo: DeprecationInfo{
						IsDeprecated: false,
						Description:  "",
						Version:      "",
					},
					Examples: []string{"example"},
				},
				"package",
				FunctionTypeFixed,
				test.args,
				test.returns,
				false,
				func(_ EvaluatorInterface, args []datavalue.Value) datavalue.Value {
					num, err := args[0].AsNumber()

					if err != nil {
						return datavalue.Number(0)
					}

					return datavalue.Number(num)
				},
			)

			if function.Expr() != test.expected {
				t.Errorf("expected \"%s\", got \"%s\"", test.expected, function.Expr())
			}
		})
	}
}
