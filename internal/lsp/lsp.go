// Package lsp provides an LSP server for DLiteScript.
package lsp

// InitializeResult represents the result for the initialize method.
type InitializeResult struct {
	ServerInfo   ServerInfo         `json:"serverInfo"`
	Capabilities ServerCapabilities `json:"capabilities"`
}

// ServerCapabilities represents the capabilities of the server.
type ServerCapabilities struct {
	TextDocumentSync TextDocumentSync `json:"textDocumentSync"`
}

// TextDocumentSync represents the text document sync capabilities.
type TextDocumentSync struct {
	OpenClose bool       `json:"openClose"`
	Change    ChangeType `json:"change"`
}

// ServerInfo represents the server information.
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
