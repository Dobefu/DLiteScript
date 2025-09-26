package tokenizer

import (
	"unicode/utf8"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// Peek gets the char of the expression at the current index without advancing it.
func (t *Tokenizer) Peek() (rune, error) {
	if t.isEOF {
		pos := t.GetCurrentPosition()

		return 0, errorutil.NewErrorAt(
			errorutil.StageTokenize,
			errorutil.ErrorMsgUnexpectedEOF,
			ast.Range{Start: pos, End: pos},
		)
	}

	r, _ := utf8.DecodeRuneInString(t.exp[t.byteIdx:])

	if r == utf8.RuneError {
		pos := t.GetCurrentPosition()

		return 0, errorutil.NewErrorAt(
			errorutil.StageTokenize,
			errorutil.ErrorMsgInvalidUTF8Char,
			ast.Range{Start: pos, End: pos},
		)
	}

	return r, nil
}
