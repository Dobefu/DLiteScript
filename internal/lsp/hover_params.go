package lsp

// HoverParams represents the parameters for a hover request.
type HoverParams struct {
	TextDocument TextDocument `json:"textDocument"`
	Position     Position     `json:"position"`
}

// TextDocument represents a text document.
type TextDocument struct {
	URI string `json:"uri"`
}

// Position represents a position in a text document.
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
