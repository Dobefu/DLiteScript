package lsp

import (
	"encoding/json"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

func (h *Handler) handleExit() (json.RawMessage, *jsonrpc2.Error) {
	go close(h.shutdownChan)

	return json.RawMessage("null"), nil
}
