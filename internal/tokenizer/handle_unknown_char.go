package tokenizer

import (
	"unicode"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleUnknownChar(next rune) (*token.Token, error) {
	if unicode.IsLetter(rune(next)) || next == '_' {
		return t.handleIdentifier(rune(next))
	}

	return nil, errorutil.NewErrorAt(
		errorutil.ErrorMsgUnexpectedChar,
		t.expIdx,
		string(next),
	)
}
