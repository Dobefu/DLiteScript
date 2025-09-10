package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handlePercentSign(startPos int) (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '=' {
		_, _ = t.GetNext()

		return token.NewToken(
			"%=",
			token.TokenTypeOperationModAssign,
			startPos,
			t.expIdx,
		), nil
	}

	return token.NewToken(
		"%",
		token.TokenTypeOperationMod,
		startPos,
		t.expIdx,
	), nil
}
