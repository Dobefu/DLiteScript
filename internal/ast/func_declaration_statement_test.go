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
		},
	}

	for _, test := range tests {
		visitedNodes := []string{}

		test.input.Walk(func(node ExprNode) bool {
			visitedNodes = append(visitedNodes, node.Expr())

			return true
		})

		if len(visitedNodes) != len(test.expectedNodes) {
			t.Fatalf(
				"Expected %d visited nodes, got %d",
				len(test.expectedNodes),
				len(visitedNodes),
			)
		}

		for idx, node := range visitedNodes {
			if node != test.expectedNodes[idx] {
				t.Fatalf("Expected \"%s\", got \"%s\"", test.expectedNodes[idx], node)
			}
		}

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
	}
}
