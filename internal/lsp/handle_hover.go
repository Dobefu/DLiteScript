package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

func (h *Handler) handleHover(
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	var hoverParams lsptypes.HoverParams
	err := json.Unmarshal(params, &hoverParams)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	response := lsptypes.Hover{
		Contents: "TODO: Implement hover",
		Range: &lsptypes.Range{
			Start: lsptypes.Position{
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
