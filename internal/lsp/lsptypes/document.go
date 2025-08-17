package lsptypes

import (
	"fmt"
)

// Document represents a document.
type Document struct {
	Text        string `json:"text"`
	Version     int    `json:"version"`
	NumLines    int    `json:"numLines"`
	LineLengths []int  `json:"lineLengths"`
}

func (d *Document) isValidLine(line int) error {
	if line < 0 || line >= d.NumLines {
		return fmt.Errorf("line %d is out of range", line)
	}

	return nil
}

// GetLineLength returns the length of the line at the given index.
func (d *Document) GetLineLength(line int) (int, error) {
	err := d.isValidLine(line)

	if err != nil {
		return 0, err
	}

	return d.LineLengths[line], err
}

// GetLine returns the line at the given index.
func (d *Document) GetLine(line int) (string, error) {
	err := d.isValidLine(line)

	if err != nil {
		return "", err
	}

	startIndex := 0

	for i := range line {
		startIndex += d.LineLengths[i]
	}

	endIndex := startIndex + d.LineLengths[line]

	return d.Text[startIndex:endIndex], err
}

// PositionToIndex converts a position(line, character) to an index.
func (d *Document) PositionToIndex(position Position) (int, error) {
	err := d.isValidLine(position.Line)

	if err != nil {
		return 0, err
	}

	index := 0

	for i := 0; i < position.Line; i++ {
		index += d.LineLengths[i]
	}

	index += position.Character

	return index, nil
}
