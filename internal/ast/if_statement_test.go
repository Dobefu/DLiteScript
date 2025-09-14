package ast

import (
	"testing"
)

func TestIfStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            ExprNode
		expectedValue    string
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "if statement without else",
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			expectedValue:    "if true { (1) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if true { (1) }",
				"true",
				"true",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			continueOn: "",
		},
		{
			name: "if statement with else",
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   5,
				ElseBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "2", StartPos: 4, EndPos: 5},
					},
					StartPos: 4,
					EndPos:   5,
				},
			},
			expectedValue:    "if true { (1) } else { (2) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if true { (1) } else { (2) }",
				"true",
				"true",
				"(1)",
				"(1)",
				"1",
				"1",
				"(2)",
				"(2)",
				"2",
				"2",
			},
			continueOn: "",
		},
		{
			name: "if statement with nil condition",
			input: &IfStatement{
				Condition: nil,
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			continueOn: "",
		},
		{
			name: "walk early return after if statement",
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "false", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			expectedValue:    "if false { (1) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if false { (1) }",
			},
			continueOn: "if false { (1) }",
		},
		{
			name: "walk early return after condition",
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			expectedValue:    "if true { (1) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if true { (1) }",
				"true",
			},
			continueOn: "true",
		},
		{
			name: "walk early return after then block",
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos:  0,
				EndPos:    5,
				ElseBlock: nil,
			},
			expectedValue:    "if true { (1) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if true { (1) }",
				"true",
				"true",
				"(1)",
			},
			continueOn: "(1)",
		},
		{
			name: "walk early return after else block",
			input: &IfStatement{
				Condition: &BoolLiteral{Value: "true", StartPos: 0, EndPos: 1},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "1", StartPos: 2, EndPos: 3},
					},
					StartPos: 2,
					EndPos:   3,
				},
				StartPos: 0,
				EndPos:   5,
				ElseBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{Value: "2", StartPos: 4, EndPos: 5},
					},
					StartPos: 4,
					EndPos:   5,
				},
			},
			expectedValue:    "if true { (1) } else { (2) }",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"if true { (1) } else { (2) }",
				"true",
				"true",
				"(1)",
				"(1)",
				"1",
				"1",
				"(2)",
			},
			continueOn: "(2)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
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
