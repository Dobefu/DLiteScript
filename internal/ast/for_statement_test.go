package ast

import (
	"testing"
)

func TestForStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		statement        *ForStatement
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "empty for loop",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body:             nil,
				StartPos:         0,
				EndPos:           0,
				RangeVariable:    "",
				RangeFrom:        nil,
				RangeTo:          nil,
				IsRange:          false,
			},
			expectedNodes:    []string{"for { }"},
			expectedStartPos: 0,
			expectedEndPos:   0,
			continueOn:       "",
		},
		{
			name: "infinite for loop",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				StartPos:      0,
				EndPos:        1,
				RangeVariable: "",
				RangeFrom:     nil,
				RangeTo:       nil,
				IsRange:       false,
			},
			expectedNodes: []string{
				"for { (1) }",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "for loop with condition and declared variable",
			statement: &ForStatement{
				DeclaredVariable: "i",
				Condition: &BoolLiteral{
					Value:    "true",
					StartPos: 0,
					EndPos:   1,
				},
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				StartPos:      0,
				EndPos:        1,
				RangeVariable: "i",
				RangeFrom:     nil,
				RangeTo:       nil,
				IsRange:       false,
			},
			expectedNodes: []string{
				"for var i true { (1) }",
				"true",
				"true",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "for loop with literal condition",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition: &BoolLiteral{
					Value:    "true",
					StartPos: 0,
					EndPos:   1,
				},
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				StartPos:      0,
				EndPos:        1,
				RangeVariable: "",
				RangeFrom:     nil,
				RangeTo:       nil,
				IsRange:       false,
			},
			expectedNodes: []string{
				"for true { (1) }",
				"true",
				"true",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "for loop with from and to",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				StartPos:      0,
				EndPos:        1,
				RangeVariable: "",
				RangeFrom: &NumberLiteral{
					Value:    "0",
					StartPos: 0,
					EndPos:   1,
				},
				RangeTo: &NumberLiteral{
					Value:    "10",
					StartPos: 0,
					EndPos:   1,
				},
				IsRange: true,
			},
			expectedNodes: []string{
				"for from 0 to 10 { (1) }",
				"0",
				"0",
				"10",
				"10",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "for loop with variable and from and to",
			statement: &ForStatement{
				DeclaredVariable: "i",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				StartPos:      0,
				EndPos:        1,
				RangeVariable: "i",
				RangeFrom: &NumberLiteral{
					Value:    "0",
					StartPos: 0,
					EndPos:   1,
				},
				RangeTo: &NumberLiteral{
					Value:    "10",
					StartPos: 0,
					EndPos:   1,
				},
				IsRange: true,
			},
			expectedNodes: []string{
				"for var i from 0 to 10 { (1) }",
				"0",
				"0",
				"10",
				"10",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "for loop with to",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				StartPos:      0,
				EndPos:        1,
				RangeVariable: "",
				RangeFrom:     nil,
				RangeTo: &NumberLiteral{
					Value:    "10",
					StartPos: 0,
					EndPos:   1,
				},
				IsRange: true,
			},
			expectedNodes: []string{
				"for from 0 to 10 { (1) }",
				"10",
				"10",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "for loop with variable and to",
			statement: &ForStatement{
				DeclaredVariable: "i",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "1",
							StartPos: 0,
							EndPos:   1,
						},
					},
					StartPos: 0,
					EndPos:   1,
				},
				StartPos:      0,
				EndPos:        1,
				RangeVariable: "i",
				RangeFrom:     nil,
				RangeTo: &NumberLiteral{
					Value:    "10",
					StartPos: 0,
					EndPos:   1,
				},
				IsRange: true,
			},
			expectedNodes: []string{
				"for var i to 10 { (1) }",
				"10",
				"10",
				"(1)",
				"(1)",
				"1",
				"1",
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after for statement",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "42",
							StartPos: 0,
							EndPos:   2,
						},
					},
					StartPos: 0,
					EndPos:   2,
				},
				StartPos:      0,
				EndPos:        2,
				RangeVariable: "",
				RangeFrom:     nil,
				RangeTo:       nil,
				IsRange:       false,
			},
			expectedNodes:    []string{"for { (42) }"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "for { (42) }",
		},
		{
			name: "walk early return after condition",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition: &BoolLiteral{
					Value:    "true",
					StartPos: 0,
					EndPos:   4,
				},
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "42",
							StartPos: 0,
							EndPos:   2,
						},
					},
					StartPos: 0,
					EndPos:   2,
				},
				StartPos:      0,
				EndPos:        2,
				RangeVariable: "",
				RangeFrom:     nil,
				RangeTo:       nil,
				IsRange:       false,
			},
			expectedNodes:    []string{"for true { (42) }", "true"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "true",
		},
		{
			name: "walk early return after range from",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "42",
							StartPos: 0,
							EndPos:   2,
						},
					},
					StartPos: 0,
					EndPos:   2,
				},
				StartPos:      0,
				EndPos:        2,
				RangeVariable: "",
				RangeFrom: &NumberLiteral{
					Value:    "0",
					StartPos: 0,
					EndPos:   1,
				},
				RangeTo: &NumberLiteral{
					Value:    "10",
					StartPos: 0,
					EndPos:   2,
				},
				IsRange: true,
			},
			expectedNodes:    []string{"for from 0 to 10 { (42) }", "0"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "0",
		},
		{
			name: "walk early return after range to",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "42",
							StartPos: 0,
							EndPos:   2,
						},
					},
					StartPos: 0,
					EndPos:   2,
				},
				StartPos:      0,
				EndPos:        2,
				RangeVariable: "",
				RangeFrom: &NumberLiteral{
					Value:    "0",
					StartPos: 0,
					EndPos:   1,
				},
				RangeTo: &NumberLiteral{
					Value:    "10",
					StartPos: 0,
					EndPos:   2,
				},
				IsRange: true,
			},
			expectedNodes:    []string{"for from 0 to 10 { (42) }", "0", "0", "10"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "10",
		},
		{
			name: "walk early return after body",
			statement: &ForStatement{
				DeclaredVariable: "",
				Condition:        nil,
				Body: &BlockStatement{
					Statements: []ExprNode{
						&NumberLiteral{
							Value:    "42",
							StartPos: 0,
							EndPos:   2,
						},
					},
					StartPos: 0,
					EndPos:   2,
				},
				StartPos:      0,
				EndPos:        2,
				RangeVariable: "",
				RangeFrom:     nil,
				RangeTo:       nil,
				IsRange:       false,
			},
			expectedNodes:    []string{"for { (42) }", "(42)"},
			expectedStartPos: 0,
			expectedEndPos:   2,
			continueOn:       "(42)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.statement.StartPosition() != test.expectedStartPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedStartPos,
					test.statement.StartPosition(),
				)
			}

			if test.statement.EndPosition() != test.expectedEndPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedEndPos,
					test.statement.EndPosition(),
				)
			}

			WalkUntil(t, test.statement, test.expectedNodes, test.continueOn)
		})
	}
}
