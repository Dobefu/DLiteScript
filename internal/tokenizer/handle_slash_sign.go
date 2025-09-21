package tokenizer

import (
	"strings"

	"github.com/Dobefu/DLiteScript/internal/token"
)

func (t *Tokenizer) handleSlashSign(startPos int) (*token.Token, error) {
	next, err := t.Peek()

	if err != nil {
		return nil, err
	}

	if next == '=' {
		_, _ = t.GetNext()

		return token.NewToken(
			"/=",
			token.TokenTypeOperationDivAssign,
			startPos,
			t.expIdx,
		), nil
	}

	if next == '/' {
		var comment strings.Builder
		comment.WriteString("/")

		for !t.isEOF {
			next, _ = t.GetNext()

			if next == '\n' {
				break
			}

			comment.WriteRune(next)
		}

		return token.NewToken(
			comment.String(),
			token.TokenTypeComment,
			startPos,
			t.expIdx,
		), nil
	}

	return token.NewToken(
		"/",
		token.TokenTypeOperationDiv,
		startPos,
		t.expIdx,
	), nil
}
