package lsp

import (
	"encoding/json"
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/jsonrpc2"
	"github.com/Dobefu/DLiteScript/internal/lsp/lsptypes"
)

// Handler represents the LSP handler.
type Handler struct {
	isDebugMode  bool
	documents    map[string]lsptypes.Document
	shutdownChan chan struct{}
	exitChan     chan struct{}
}

// NewHandler creates a new LSP handler.
func NewHandler(isDebugMode bool) *Handler {
	return &Handler{
		isDebugMode:  isDebugMode,
		documents:    make(map[string]lsptypes.Document),
		shutdownChan: make(chan struct{}),
		exitChan:     make(chan struct{}),
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
		return h.handleDidChange(params)

	case "textDocument/didClose":
		return h.handleDidClose(params)

	case "textDocument/hover":
		return h.handleHover(params)

	case "textDocument/signatureHelp":
		return h.handleSignatureHelp(params)

	case "textDocument/completion":
		return h.handleCompletion(params)

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
