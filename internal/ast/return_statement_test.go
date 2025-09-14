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
		continueOn       string
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
			continueOn:       "",
		},
		{
			name: "nil return value",
			input: &ReturnStatement{
				Values:    []ExprNode{nil},
				NumValues: 1,
				StartPos:  0,
				EndPos:    1,
			},
			expectedValue:    "return ",
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"return "},
			continueOn:       "",
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
			expectedNodes:    []string{"return 1", "1", "1"},
			continueOn:       "",
		},
		{
			name: "multiple values",
			input: &ReturnStatement{
				Values: []ExprNode{
					&NumberLiteral{
						Value:    "1",
						StartPos: 0,
						EndPos:   1,
					},
					&NumberLiteral{
						Value:    "2",
						StartPos: 2,
						EndPos:   3,
					},
				},
				NumValues: 2,
				StartPos:  0,
				EndPos:    3,
			},
			expectedValue:    "return 1, 2",
			expectedStartPos: 0,
			expectedEndPos:   3,
			expectedNodes:    []string{"return 1, 2", "1", "1", "2", "2"},
			continueOn:       "",
		},
		{
			name: "return with string value",
			input: &ReturnStatement{
				Values: []ExprNode{
					&StringLiteral{
						Value:    "hello",
						StartPos: 0,
						EndPos:   5,
					},
				},
				NumValues: 1,
				StartPos:  0,
				EndPos:    5,
			},
			expectedValue:    "return \"hello\"",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"return \"hello\"", "\"hello\"", "\"hello\""},
			continueOn:       "",
		},
		{
			name: "walk early return after return statement",
			input: &ReturnStatement{
				Values: []ExprNode{
					&NumberLiteral{
						Value:    "42",
						StartPos: 0,
						EndPos:   2,
					},
				},
				NumValues: 1,
				StartPos:  0,
				EndPos:    2,
			},
			expectedValue:    "return 42",
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"return 42"},
			continueOn:       "return 42",
		},
		{
			name: "walk early return after first value",
			input: &ReturnStatement{
				Values: []ExprNode{
					&NumberLiteral{
						Value:    "10",
						StartPos: 0,
						EndPos:   2,
					},
					&NumberLiteral{
						Value:    "20",
						StartPos: 3,
						EndPos:   5,
					},
				},
				NumValues: 2,
				StartPos:  0,
				EndPos:    5,
			},
			expectedValue:    "return 10, 20",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"return 10, 20", "10"},
			continueOn:       "10",
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
				t.Fatalf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.StartPosition(),
				)
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Fatalf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.EndPosition(),
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
