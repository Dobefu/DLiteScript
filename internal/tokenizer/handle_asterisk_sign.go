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
		_, _ = t.GetNext()

		nextAfterSecond, err := t.Peek()
		if err == nil && nextAfterSecond == '=' {
			_, _ = t.GetNext()

			return token.NewToken(
				"**=",
				token.TokenTypeOperationPowAssign,
				startPos,
				t.expIdx,
			), nil
		}

		return token.NewToken(
			"**",
			token.TokenTypeOperationPow,
			startPos,
			t.expIdx,
		), nil
	}

	if next == '=' {
		_, _ = t.GetNext()

		return token.NewToken(
			"*=",
			token.TokenTypeOperationMulAssign,
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
