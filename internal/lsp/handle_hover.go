package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

func (h *Handler) handleHover(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var hoverParams HoverParams
	err := json.Unmarshal(params, &hoverParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	response := Hover{
		Contents: "TODO: Implement hover",
		Range: &Range{
			Start: Position{
				Line:      hoverParams.Position.Line,
				Character: hoverParams.Position.Character,
			},
			End: hoverParams.Position,
		},
	}

	data, err := json.Marshal(response)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInternalError,
			err.Error(),
			nil,
		)
	}

	return data, nil
}
