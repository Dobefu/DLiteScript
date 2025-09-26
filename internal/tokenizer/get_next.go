package tokenizer

import (
	"unicode/utf8"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// GetNext gets the next character in the expression.
func (t *Tokenizer) GetNext() (rune, error) {
	if t.isEOF {
		pos := t.GetCurrentPosition()

		return 0, errorutil.NewErrorAt(
			errorutil.StageTokenize,
			errorutil.ErrorMsgUnexpectedEOF,
			ast.Range{Start: pos, End: pos},
		)
	}

	r, size := utf8.DecodeRuneInString(t.exp[t.byteIdx:])

	if r == utf8.RuneError {
		pos := t.GetCurrentPosition()

		return 0, errorutil.NewErrorAt(
			errorutil.StageTokenize,
			errorutil.ErrorMsgInvalidUTF8Char,
			ast.Range{Start: pos, End: pos},
		)
	}

	t.byteIdx += size
	t.expIdx++

	if r == '\n' {
		t.line++
		t.col = 0
	} else {
		t.col++
	}

	if t.expIdx >= t.expLen {
		t.isEOF = true
	}

	return r, nil
}
