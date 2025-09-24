package ast

import (
	"testing"
)

func TestImportStatement(t *testing.T) {
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
			name: "import statement",
			input: &ImportStatement{
				Path: &StringLiteral{
					Value: "test",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    `import "test"`,
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				`import "test"`,
			},
			continueOn: "",
		},
		{
			name: "import statement with alias",
			input: &ImportStatement{
				Path: &StringLiteral{
					Value: "test",
					Range: Range{
						Start: Position{Offset: 0, Line: 0, Column: 0},
						End:   Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "test",
				Range: Range{
					Start: Position{Offset: 0, Line: 0, Column: 0},
					End:   Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expectedValue:    `import "test" as test`,
			expectedStartPos: 0,
			expectedEndPos:   5,
			expectedNodes: []string{
				`import "test" as test`,
			},
			continueOn: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.Expr() != test.expectedValue {
				t.Errorf(
					"expected '%s', got '%s'",
					test.expectedValue,
					test.input.Expr(),
				)
			}

			if test.input.GetRange().Start.Offset != test.expectedStartPos {
				t.Errorf(
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
