package ast

import (
	"testing"
)

func TestVariableDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *VariableDeclaration
		expectedStartPos int
		expectedEndPos   int
		expectedNodes    []string
		continueOn       string
	}{
		{
			name: "variable declaration with value",
			input: &VariableDeclaration{
				Name: "x",
				Type: "int",
				Value: &NumberLiteral{
					Value:    "1",
					StartPos: 0,
					EndPos:   1,
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"var x int = 1", "1", "1"},
			continueOn:       "",
		},
		{
			name: "variable declaration without value",
			input: &VariableDeclaration{
				Name:     "x",
				Type:     "int",
				Value:    nil,
				StartPos: 0,
				EndPos:   1,
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"var x int"},
			continueOn:       "",
		},
		{
			name: "walk early return after declaration node",
			input: &VariableDeclaration{
				Name: "y",
				Type: "string",
				Value: &StringLiteral{
					Value:    "hello",
					StartPos: 0,
					EndPos:   2,
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"var y string = \"hello\""},
			continueOn:       "var y string = \"hello\"",
		},
		{
			name: "walk early return after value",
			input: &VariableDeclaration{
				Name: "y",
				Type: "string",
				Value: &StringLiteral{
					Value:    "hello",
					StartPos: 0,
					EndPos:   2,
				},
				StartPos: 0,
				EndPos:   2,
			},
			expectedStartPos: 0,
			expectedEndPos:   2,
			expectedNodes:    []string{"var y string = \"hello\"", "\"hello\""},
			continueOn:       "\"hello\"",
		},
		{
			name: "variable declaration with different type",
			input: &VariableDeclaration{
				Name: "z",
				Type: "bool",
				Value: &BoolLiteral{
					Value:    "true",
					StartPos: 0,
					EndPos:   1,
				},
				StartPos: 0,
				EndPos:   1,
			},
			expectedStartPos: 0,
			expectedEndPos:   1,
			expectedNodes:    []string{"var z bool = true", "true", "true"},
			continueOn:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.StartPosition() != test.expectedStartPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedStartPos,
					test.input.StartPosition(),
				)
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Fatalf(
					"expected %d, got %d",
					test.expectedEndPos,
					test.input.EndPosition(),
				)
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
