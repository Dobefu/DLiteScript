package lsp

import (
	"encoding/json"
	"fmt"

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

	document, hasDocument := h.documents[hoverParams.TextDocument.URI]

	if !hasDocument {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			"Document not found",
			nil,
		)
	}

	ast, err := parseDocumentToAst(document.Text)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	charIndex, err := document.PositionToIndex(hoverParams.Position)

	if err != nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			err.Error(),
			nil,
		)
	}

	node := getAstNodeAtPosition(ast, charIndex)

	if node == nil {
		return nil, jsonrpc2.NewError(
			jsonrpc2.ErrorCodeInvalidParams,
			"Could not find AST node at position",
			nil,
		)
	}

	content := fmt.Sprintf("%T\n\n%s", node, node.Expr())

	response := lsptypes.Hover{
		Contents: content,
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
