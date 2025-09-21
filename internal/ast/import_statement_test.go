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
				Path:      &StringLiteral{Value: "test", StartPos: 0, EndPos: 5},
				Namespace: "test",
				Alias:     "",
				StartPos:  0,
				EndPos:    5,
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
				Path:      &StringLiteral{Value: "test", StartPos: 0, EndPos: 5},
				Namespace: "test",
				Alias:     "test",
				StartPos:  0,
				EndPos:    5,
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
				t.Errorf("expected '%s', got '%s'", test.expectedValue, test.input.Expr())
			}

			if test.input.StartPosition() != test.expectedStartPos {
				t.Errorf("expected pos '%d', got '%d'", test.expectedStartPos, test.input.StartPosition())
			}

			if test.input.EndPosition() != test.expectedEndPos {
				t.Errorf("expected pos '%d', got '%d'", test.expectedEndPos, test.input.EndPosition())
			}

			WalkUntil(t, test.input, test.expectedNodes, test.continueOn)
		})
	}
}
