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
				Condition: &BoolLiteral{
					Value: "true",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "1",
							Range: Range{
								Start: Position{Offset: 2, Line: 0, Column: 0},
								End:   Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
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
				Condition: &BoolLiteral{
					Value: "true",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "1",
							Range: Range{
								Start: Position{Offset: 2, Line: 0, Column: 0},
								End:   Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
				ElseBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "2",
							Range: Range{
								Start: Position{Offset: 4, Line: 0, Column: 0},
								End:   Position{Offset: 5, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 4, Line: 0, Column: 0},
						End:   Position{Offset: 5, Line: 0, Column: 0},
					},
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
						&NumberLiteral{
							Value: "1",
							Range: Range{
								Start: Position{Offset: 2, Line: 0, Column: 0},
								End:   Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
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
				Condition: &BoolLiteral{
					Value: "false",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "1",
							Range: Range{
								Start: Position{Offset: 2, Line: 0, Column: 0},
								End:   Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
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
				Condition: &BoolLiteral{
					Value: "true",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "1",
							Range: Range{
								Start: Position{Offset: 2, Line: 0, Column: 0},
								End:   Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
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
				Condition: &BoolLiteral{
					Value: "true",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "1",
							Range: Range{
								Start: Position{Offset: 2, Line: 0, Column: 0},
								End:   Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
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
				Condition: &BoolLiteral{
					Value: "true",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 1, Line: 0, Column: 0},
					},
				},
				ThenBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "1",
							Range: Range{
								Start: Position{Offset: 2, Line: 0, Column: 0},
								End:   Position{Offset: 3, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 2, Line: 0, Column: 0},
						End:   Position{Offset: 3, Line: 0, Column: 0},
					},
				},
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
				ElseBlock: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value: "2",
							Range: Range{
								Start: Position{Offset: 4, Line: 0, Column: 0},
								End:   Position{Offset: 5, Line: 0, Column: 0},
							},
						},
					},
					Range: Range{
						Start: Position{Offset: 4, Line: 0, Column: 0},
						End:   Position{Offset: 5, Line: 0, Column: 0},
					},
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
				t.Fatalf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Fatalf(
					"expected pos '%d', got '%d'",
					test.expectedStartPos,
					test.input.GetRange().Start.Offset,
				)
			}

			if test.input.GetRange().End.Offset != test.expectedEndPos {
				t.Errorf(
					"expected pos '%d', got '%d'",
					test.expectedEndPos,
					test.input.GetRange().End.Offset,
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
