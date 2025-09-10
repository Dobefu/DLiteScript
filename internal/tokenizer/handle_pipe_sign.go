package tokenizer

import (
	"github.com/Dobefu/DLiteScript/internal/errorutil"
	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handlePipeSign(startPos int) (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '|' {
		_, err = t.GetNext()

		if err != nil {
			return nil, err
		}

		return token.NewToken(
			"||",
			token.TokenTypeLogicalOr,
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
