package ast

import (
	"testing"
)

func TestIndexAssignmentStatement(t *testing.T) {
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
			name: "index assignment statement",
			input: &IndexAssignmentStatement{
				Array:    &Identifier{Value: "array", StartPos: 0, EndPos: 5},
				Index:    &NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "array[0] = 1",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				"array[0] = 1",
				"array",
				"array",
				"0",
				"0",
				"1",
			},
			continueOn: "",
		},
		{
			name: "index assignment with nil array",
			input: &IndexAssignmentStatement{
				Array:    nil,
				Index:    &NumberLiteral{Value: "0", StartPos: 0, EndPos: 1},
				Right:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"", "0", "0", "1"},
			continueOn:       "",
		},
		{
			name: "index assignment with nil index",
			input: &IndexAssignmentStatement{
				Array:    &Identifier{Value: "array", StartPos: 0, EndPos: 5},
				Index:    nil,
				Right:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"", "array", "array", "1"},
			continueOn:       "",
		},
		{
			name: "walk early return after index assignment",
			input: &IndexAssignmentStatement{
				Array:    &Identifier{Value: "arr", StartPos: 0, EndPos: 3},
				Index:    &NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				Right:    &NumberLiteral{Value: "3", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "arr[2] = 3",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"arr[2] = 3"},
			continueOn:       "arr[2] = 3",
		},
		{
			name: "walk early return after array",
			input: &IndexAssignmentStatement{
				Array:    &Identifier{Value: "list", StartPos: 0, EndPos: 4},
				Index:    &NumberLiteral{Value: "1", StartPos: 0, EndPos: 1},
				Right:    &NumberLiteral{Value: "2", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "list[1] = 2",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"list[1] = 2", "list"},
			continueOn:       "list",
		},
		{
			name: "walk early return after index",
			input: &IndexAssignmentStatement{
				Array:    &Identifier{Value: "data", StartPos: 0, EndPos: 4},
				Index:    &NumberLiteral{Value: "5", StartPos: 0, EndPos: 1},
				Right:    &NumberLiteral{Value: "6", StartPos: 0, EndPos: 1},
				StartPos: 0,
				EndPos:   5,
			},
			expectedValue:    "data[5] = 6",
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes:    []string{"data[5] = 6", "data", "data", "5"},
			continueOn:       "5",
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
