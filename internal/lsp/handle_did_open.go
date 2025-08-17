package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func (h *Handler) handleDidOpen(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var didOpenParams lsptypes.DidOpenParams
	err := json.Unmarshal(params, &didOpenParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	numLines, lineLengths := calculateLineCountAndLengths(didOpenParams.TextDocument.Text)

	h.documents[didOpenParams.TextDocument.URI] = lsptypes.Document{
		Text:        didOpenParams.TextDocument.Text,
		Version:     didOpenParams.TextDocument.Version,
		NumLines:    numLines,
		LineLengths: lineLengths,
	}

	return json.RawMessage("null"), nil
}
