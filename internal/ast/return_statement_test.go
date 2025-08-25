package ast

import "testing"

func TestReturnStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *ReturnStatement
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
	}{
		{
			name: "simple",
			input: &ReturnStatement{
				Values:    []ExprNode{},
				NumValues: 0,
				StartPos:  0,
				EndPos:    1,
			},
			expectedValue:    "return",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"return"},
		},
		{
			name: "single value",
			input: &ReturnStatement{
				Values: []ExprNode{
					&NumberLiteral{
						Value:    "1",
						StartPos: 0,
						EndPos:   3,
					},
				},
				NumValues: 1,
				StartPos:  0,
				EndPos:    3,
			},
			expectedValue:    "return 1",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"return 1", "1"},
		},
	}

	for _, test := range tests {
		visitedNodes := []string{}

		test.input.Walk(func(node ExprNode) bool {
			visitedNodes = append(visitedNodes, node.Expr())

			return true
		})

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
