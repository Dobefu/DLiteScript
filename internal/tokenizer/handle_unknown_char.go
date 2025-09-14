package tokenizer

import (
	"unicode"

	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleUnknownChar(
	next rune,
	startPos int,
) (*token.Token, error) {
	if unicode.IsLetter(next) || next == '_' {
		return t.handleIdentifier(next, startPos)
	}

	return nil, errorutil.NewErrorAt(
		errorutil.StageTokenize,
		errorutil.ErrorMsgUnexpectedChar,
		t.expIdx,
		string(next),
	)
}
