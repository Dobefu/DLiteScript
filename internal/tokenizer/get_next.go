package tokenizer

import (
	"unicode/utf8"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

// GetNext gets the next character in the expression.
func (t *Tokenizer) GetNext() (rune, error) {
	if t.isEOF {
		return 0, errorutil.NewErrorAt(errorutil.ErrorMsgUnexpectedEOF, t.expIdx)
	}

	r, size := utf8.DecodeRuneInString(t.exp[t.byteIdx:])

	if r == utf8.RuneError {
		return 0, errorutil.NewErrorAt(errorutil.ErrorMsgInvalidUTF8Char, t.expIdx)
	}

	t.byteIdx += size
	t.expIdx++

	if t.expIdx >= t.expLen {
		t.isEOF = true
	}

	return r, nil
}
