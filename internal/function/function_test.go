package function

import (
	"testing"

	"github.com/Dobefu/DLiteScript/internal/datatype"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
)

func TestMakeFunction(t *testing.T) {
	t.Parallel()

	function := MakeFunction(
		"test",
		"description",
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
		"v1.0.0",
		DeprecationInfo{
			IsDeprecated: false,
			Description:  "",
			Version:      "",
		},
		[]string{
			"example",
		},
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

	if function.Name != "test" {
		t.Errorf("expected 'test', got '%s'", function.Name)
	}

	if function.Description != "description" {
		t.Errorf("expected 'description', got '%s'", function.Description)
	}

	if function.PackageName != "package" {
		t.Errorf("expected 'package', got '%s'", function.PackageName)
	}

	if function.Since != "v1.0.0" {
		t.Errorf("expected 'v1.0.0', got '%s'", function.Since)
	}

	if function.DeprecationInfo.IsDeprecated != false {
		t.Errorf("expected 'false', got '%t'", function.DeprecationInfo.IsDeprecated)
	}

	if function.Examples[0] != "example" {
		t.Errorf("expected 'example', got '%s'", function.Examples[0])
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
				"test",
				"description",
				"package",
				FunctionTypeFixed,
				test.args,
				test.returns,
				false,
				"v1.0.0",
				DeprecationInfo{
					IsDeprecated: false,
					Description:  "",
					Version:      "",
				},
				[]string{
					"example",
				},
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
