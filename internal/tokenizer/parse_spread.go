package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) parseSpread(startPos int) (*token.Token, error) {
	_, err := t.GetNext()

	if err != nil {
		return nil, err
	}

	_, err = t.GetNext()

	if err != nil {
		return nil, err
	}

	return token.NewToken(
		"...",
		token.TokenTypeOperationSpread,
		startPos,
		t.expIdx,
	), nil
}
