package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handlePlusSign(startPos int) (*token.Token, error) {
	nextToken, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if nextToken == '=' {
		_, _ = t.GetNext()

		return token.NewToken(
			"+=",
			token.TokenTypeOperationAddAssign,
			startPos,
			t.expIdx,
		), nil
	}

	return token.NewToken(
		"+",
		token.TokenTypeOperationAdd,
		startPos,
		t.expIdx,
	), nil
}
