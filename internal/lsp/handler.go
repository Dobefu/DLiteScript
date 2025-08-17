package lsp

import (
	"encoding/json"
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
)

// Handler represents the LSP handler.
type Handler struct {
	isDebugMode  bool
	documents    map[string]string
	shutdownChan chan struct{}
}

// NewHandler creates a new LSP handler.
func NewHandler(isDebugMode bool) *Handler {
	return &Handler{
		isDebugMode:  isDebugMode,
		documents:    make(map[string]string),
		shutdownChan: make(chan struct{}),
	}
}

// GetShutdownChan returns the shutdown channel.
func (h *Handler) GetShutdownChan() chan struct{} {
	return h.shutdownChan
}

// Handle handles a JSON-RPC request.
func (h *Handler) Handle(
	method string,
	params json.RawMessage,
) (json.RawMessage, *jsonrpc2.Error) {
	if h.isDebugMode {
		h.printDebugMessage(method, params)
	}

	switch method {
	case "initialize":
		return h.handleInitialize()

	case "initialized":
		return nil, nil

	case "textDocument/didOpen":
		return h.handleDidOpen(params)

	case "$/setTrace":
		return nil, nil

	case "textDocument/didChange":
		return nil, nil

	case "textDocument/didClose":
		return nil, nil

	case "textDocument/hover":
		return h.handleHover(params)

	case "shutdown":
		return h.handleShutdown()

	case "exit":
		return h.handleExit()

	default:
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeMethodNotFound,
			fmt.Sprintf("method %s not found", method),
			nil,
		)
	}
}
