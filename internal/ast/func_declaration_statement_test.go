package ast

import "testing"

func TestFuncDeclarationStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *FuncDeclarationStatement
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "simple",
			input: &FuncDeclarationStatement{
				Name:            "test",
				Args:            []FuncParameter{},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				Body: &NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expectedValue:    "func test()",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"func test()", "1", "1"},
			continueOn:       "",
		},
		{
			name: "single return value",
			input: &FuncDeclarationStatement{
				Name: "test",
				Args: []FuncParameter{
					{
						Name: "a",
						Type: "number",
					},
				},
				ReturnValues:    []string{"number"},
				NumReturnValues: 1,
				Body: &NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expectedValue:    "func test(a number) number",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"func test(a number) number", "1", "1"},
			continueOn:       "",
		},
		{
			name: "multiple return values",
			input: &FuncDeclarationStatement{
				Name: "test",
				Args: []FuncParameter{
					{
						Name: "a",
						Type: "number",
					},
				},
				ReturnValues:    []string{"number", "string"},
				NumReturnValues: 2,
				Body: &NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   3,
			},
			expectedValue:    "func test(a number) number, string",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes: []string{
				"func test(a number) number, string",
				"1",
				"1",
			},
			continueOn: "",
		},
		{
			name: "walk early return after function declaration",
			input: &FuncDeclarationStatement{
				Name:            "test",
				Args:            []FuncParameter{},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				Body: &NumberLiteral{
					Value:    "42",
					StartPos: 0,
					EndPos:   2,
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedValue:    "func test()",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"func test()"},
			continueOn:       "func test()",
		},
		{
			name: "walk early return after body",
			input: &FuncDeclarationStatement{
				Name:            "test",
				Args:            []FuncParameter{},
				ReturnValues:    []string{},
				NumReturnValues: 0,
				Body: &NumberLiteral{
					Value:    "42",
					StartPos: 0,
					EndPos:   2,
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedValue:    "func test()",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"func test()", "42"},
			continueOn:       "42",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Fatalf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.StartPosition() != test.expectedStartPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.StartPosition(),
				)
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.EndPosition(),
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
