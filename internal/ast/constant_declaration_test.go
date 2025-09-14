package ast

import (
	"testing"
)

func TestConstantDeclaration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		input            *ConstantDeclaration
		expectedNodes    []string
		expectedStartPos int
		expectedEndPos   int
		continueOn       string
	}{
		{
			name: "constant declaration with int type",
			input: &ConstantDeclaration{
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
			expectedNodes:    []string{"const x int = 1", "1", "1"},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "constant declaration with string type",
			input: &ConstantDeclaration{
				Name: "y",
				Type: "string",
				Value: &StringLiteral{
					Value:    "hello",
					StartPos: 0,
					EndPos:   5,
				},
				StartPos: 0,
				EndPos:   5,
			},
			expectedNodes:    []string{`const y string = "hello"`, `"hello"`, `"hello"`},
			expectedStartPos: 0,
			expectedEndPos:   5,
			continueOn:       "",
		},
		{
			name: "constant declaration without value",
			input: &ConstantDeclaration{
				Name:     "z",
				Type:     "bool",
				Value:    nil,
				StartPos: 0,
				EndPos:   1,
			},
			expectedNodes:    []string{""},
			expectedStartPos: 0,
			expectedEndPos:   1,
			continueOn:       "",
		},
		{
			name: "walk early return after constant declaration",
			input: &ConstantDeclaration{
				Name: "w",
				Type: "float",
				Value: &NumberLiteral{
					Value:    "3.14",
					StartPos: 0,
					EndPos:   4,
				},
				StartPos: 0,
				EndPos:   4,
			},
			expectedNodes:    []string{"const w float = 3.14"},
			expectedStartPos: 0,
			expectedEndPos:   4,
			continueOn:       "const w float = 3.14",
		},
		{
			name: "walk early return after value",
			input: &ConstantDeclaration{
				Name: "w",
				Type: "float",
				Value: &NumberLiteral{
					Value:    "3.14",
					StartPos: 0,
					EndPos:   4,
				},
				StartPos: 0,
				EndPos:   4,
			},
			expectedNodes:    []string{"const w float = 3.14", "3.14"},
			expectedStartPos: 0,
			expectedEndPos:   4,
			continueOn:       "3.14",
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
