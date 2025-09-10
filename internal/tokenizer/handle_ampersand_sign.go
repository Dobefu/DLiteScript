package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleAmpersandSign(startPos int) (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '&' {
		_, _ = t.GetNext()

		return token.NewToken(
			"&&",
			token.TokenTypeLogicalAnd,
			startPos,
			t.expIdx,
		), nil
	}

	return nil, errorutil.NewErrorAt(
		errorutil.StageTokenize,
		errorutil.ErrorMsgUnexpectedChar,
		t.expIdx,
		string(next),
	)
}
