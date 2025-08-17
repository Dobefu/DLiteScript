package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleAsteriskSign(startPos int) (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '*' {
		_, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		return token.NewToken(
			"**",
			token.TokenTypeOperationPow,
			startPos,
			t.expIdx,
		), nil
	}

	return token.NewToken(
		"*",
		token.TokenTypeOperationMul,
		startPos,
		t.expIdx,
	), nil
}
