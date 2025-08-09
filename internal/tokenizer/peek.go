package tokenizer

import (
	"unicode/utf8"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// Peek gets the char of the expression at the current index without advancing it.
func (t *Tokenizer) Peek() (rune, error) {
	if t.isEOF {
		return 0, errorutil.NewErrorAt(errorutil.ErrorMsgUnexpectedEOF, t.expIdx)
	}

	r, _ := utf8.DecodeRuneInString(t.exp[t.byteIdx:])

	if r == utf8.RuneError {
		return 0, errorutil.NewErrorAt(errorutil.ErrorMsgInvalidUTF8Char, t.expIdx)
	}

	return r, nil
}
