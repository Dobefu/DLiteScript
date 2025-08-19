package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func (h *Handler) handleDidChange(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var didChangeParams lsptypes.DidChangeParams
	err := json.Unmarshal(params, &didChangeParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	document, hasDocument := h.documents[didChangeParams.TextDocument.URI]

	if !hasDocument {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			"Document not found",
			nil,
		)
	}

	for _, change := range didChangeParams.ContentChanges {
		startIndex, err := document.PositionToIndex(change.Range.Start)

		if err != nil {
			return nil, jsonrpc2.NewError(
				jsonrpc2.ErrorCodeInvalidParams,
				err.Error(),
				nil,
			)
		}

		endIndex, err := document.PositionToIndex(change.Range.End)

		if err != nil {
			return nil, jsonrpc2.NewError(
				jsonrpc2.ErrorCodeInvalidParams,
				err.Error(),
				nil,
			)
		}

		if startIndex > endIndex {
			return nil, jsonrpc2.NewError(
				jsonrpc2.ErrorCodeInvalidParams,
				"Start index is greater than end index",
				nil,
			)
		}

		if endIndex > len(document.Text) {
			return nil, jsonrpc2.NewError(
				jsonrpc2.ErrorCodeInvalidParams,
				"End index is out of bounds",
				nil,
			)
		}

		newText := document.Text[:startIndex] + change.Text + document.Text[endIndex:]
		numLines, lineLengths := calculateLineCountAndLengths(newText)

		h.documents[didChangeParams.TextDocument.URI] = lsptypes.Document{
			Text:        newText,
			Version:     didChangeParams.TextDocument.Version,
			NumLines:    numLines,
			LineLengths: lineLengths,
		}
	}

	return nil, nil
}
