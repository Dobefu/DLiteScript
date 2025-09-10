package errorutil

import "testing"

func TestStage_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Stage
		expected string
	}{
		{
			name:     "tokenize",
			input:    StageTokenize,
			expected: "tokenize",
		},
		{
			name:     "parse",
			input:    StageParse,
			expected: "parse",
		},
		{
			name:     "evaluate",
			input:    StageEvaluate,
			expected: "evaluate",
		},
		{
			name:     "unknown",
			input:    Stage(-1),
			expected: "unknown stage",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if test.input.String() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, test.input.String())
			}
		})
	}
}
