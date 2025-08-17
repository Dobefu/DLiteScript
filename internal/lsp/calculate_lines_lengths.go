package lsp

func calculateLineCountAndLengths(text string) (int, []int) {
	numNewLines := 0
	lineLengths := []int{}
	currentLineLength := 0

	for _, char := range text {
		if char == '\n' {
			numNewLines++
			lineLengths = append(lineLengths, currentLineLength)
			currentLineLength = 0
		} else {
			currentLineLength++
		}
	}

	if currentLineLength > 0 || len(lineLengths) == 0 {
		lineLengths = append(lineLengths, currentLineLength)
	}

	return numNewLines + 1, lineLengths
}
