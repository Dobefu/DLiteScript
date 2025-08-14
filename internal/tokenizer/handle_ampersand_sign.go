package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleAmpersandSign() (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '&' {
		_, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		return t.tokenPool.GetToken("&&", token.TokenTypeLogicalAnd), nil
	}

	return nil, errorutil.NewErrorAt(
		errorutil.ErrorMsgUnexpectedChar,
		t.expIdx,
		string(next),
	)
}
