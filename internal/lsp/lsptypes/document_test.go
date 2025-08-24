package lsptypes

import (
	"testing"
)

func TestDocumentGetLine(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		doc                 Document
		targetLine          int
		expectedLineLength  int
		expectedLineContent string
		expectedIndex       int
	}{
		{
			name: "basic document",
			doc: Document{
				Text:        "test\n",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{4},
			},
			targetLine:          0,
			expectedLineLength:  4,
			expectedLineContent: "test",
			expectedIndex:       0,
		},
		{
			name: "multi-line document",
			doc: Document{
				Text:        "test\nline2",
				Version:     1,
				NumLines:    2,
				LineLengths: []int{4, 5},
			},
			targetLine:          1,
			expectedLineLength:  5,
			expectedLineContent: "line2",
			expectedIndex:       5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			lineLength, err := test.doc.GetLineLength(test.targetLine)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if lineLength != test.expectedLineLength {
				t.Fatalf("expected line length to be %d, got %d", test.expectedLineLength, lineLength)
			}

			lineContent, err := test.doc.GetLine(test.targetLine)

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if lineContent != test.expectedLineContent {
				t.Fatalf("expected line content to be \"%s\", got \"%s\"", test.expectedLineContent, lineContent)
			}
		})
	}
}

func TestDocumentPositionToIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		doc                 Document
		targetLine          int
		expectedLineLength  int
		expectedLineContent string
		expectedIndex       int
	}{
		{
			name: "basic document",
			doc: Document{
				Text:        "test\n",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{4},
			},
			targetLine:          0,
			expectedLineLength:  4,
			expectedLineContent: "test",
			expectedIndex:       0,
		},
		{
			name: "multi-line document",
			doc: Document{
				Text:        "test\nline2",
				Version:     1,
				NumLines:    2,
				LineLengths: []int{4, 5},
			},
			targetLine:          1,
			expectedLineLength:  5,
			expectedLineContent: "line2",
			expectedIndex:       5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			index, err := test.doc.PositionToIndex(Position{
				Line:      test.targetLine,
				Character: 0,
			})

			if err != nil {
				t.Fatalf("expected no error, got \"%s\"", err.Error())
			}

			if index != test.expectedIndex {
				t.Fatalf("expected index to be %d, got %d", test.expectedIndex, index)
			}
		})
	}
}

func TestDocumentErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		doc        Document
		targetLine int
		expected   string
	}{
		{
			name: "line out of range",
			doc: Document{
				Text:        "",
				Version:     1,
				NumLines:    1,
				LineLengths: []int{4},
			},
			targetLine: 1,
			expected:   "line 1 is out of range",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			_, err := test.doc.GetLineLength(test.targetLine)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}

			if err.Error() != test.expected {
				t.Fatalf("expected error to be \"%s\", got \"%s\"", test.expected, err)
			}

			_, err = test.doc.GetLine(test.targetLine)

			if err.Error() != test.expected {
				t.Fatalf("expected error to be \"%s\", got \"%s\"", test.expected, err)
			}

			_, err = test.doc.PositionToIndex(Position{
				Line:      test.targetLine,
				Character: 0,
			})

			if err.Error() != test.expected {
				t.Fatalf("expected error to be \"%s\", got \"%s\"", test.expected, err)
			}
		})
	}
}
