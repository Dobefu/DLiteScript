package ast

import (
	"testing"
)

func validatePosition(
	t *testing.T,
	statement *ForStatement,
	expectedStart int,
	expectedEnd int,
) {
	t.Helper()

	if statement.StartPosition() != expectedStart {
		t.Fatalf(
			"Expected start position to be %d, got %d",
			expectedStart,
			statement.StartPosition(),
		)
	}

	if statement.EndPosition() != expectedEnd {
		t.Fatalf(
			"Expected end position to be %d, got %d",
			expectedEnd,
			statement.EndPosition(),
		)
	}
}

func TestForStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		statement     *ForStatement
		expectedNodes []string
		expectedStart int
		expectedEnd   int
	}{
		{
			name: "basic for loop",
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
			expectedStart: 0,
			expectedEnd:   1,
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
			expectedStart: 0,
			expectedEnd:   1,
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
			expectedStart: 0,
			expectedEnd:   1,
		},
		{
			name: "for loop with range",
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
			expectedStart: 0,
			expectedEnd:   1,
		},
		{
			name: "for loop with range to only",
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
			expectedStart: 0,
			expectedEnd:   1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			visitedNodes := []string{}

			test.statement.Walk(func(node ExprNode) bool {
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

			validatePosition(t, test.statement, test.expectedStart, test.expectedEnd)
		})
	}
}
