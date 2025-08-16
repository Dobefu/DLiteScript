package lsp

import (
	"encoding/json"
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

// Handler represents the LSP handler.
type Handler struct {
	documents map[string]string
}

// NewHandler creates a new LSP handler.
func NewHandler() *Handler {
	return &Handler{
		documents: make(map[string]string),
	}
}

// Handle handles a JSON-RPC request.
func (h *Handler) Handle(
	method string,
	_ json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	switch method {
	default:
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeMethodNotFound,
			fmt.Sprintf("method %s not found", method),
			nil,
		)
	}
}
