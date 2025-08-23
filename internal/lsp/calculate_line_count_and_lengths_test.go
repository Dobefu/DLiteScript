package lsp

import (
	"testing"
)

func TestCalculateLineCountAndLengths(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		input               string
		expectedLineCount   int
		expectedLineLengths []int
	}{
		{
			name:                "empty string",
			input:               "",
			expectedLineCount:   1,
			expectedLineLengths: []int{0},
		},
		{
			name:                "single line",
			input:               "test",
			expectedLineCount:   1,
			expectedLineLengths: []int{4},
		},
		{
			name:                "multiple lines",
			input:               "test\ntest",
			expectedLineCount:   2,
			expectedLineLengths: []int{4, 4},
		},
		{
			name:                "multiple lines with trailing newline",
			input:               "test\ntest\n",
			expectedLineCount:   3,
			expectedLineLengths: []int{4, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			lineCount, lineLengths := calculateLineCountAndLengths(test.input)

			if lineCount != test.expectedLineCount {
				t.Errorf("expected %d, got %d", test.expectedLineCount, lineCount)
			}

			if len(lineLengths) != len(test.expectedLineLengths) {
				t.Errorf("expected %d, got %d", len(test.expectedLineLengths), len(lineLengths))
			}
		})
	}
}
