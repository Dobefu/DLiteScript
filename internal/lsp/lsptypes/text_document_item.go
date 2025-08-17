package lsptypes

// TextDocumentItem represents a text document item.
type TextDocumentItem struct {
	LanguageID string `json:"languageId"`
	Text       string `json:"text"`
	URI        string `json:"uri"`
	Version    int    `json:"version"`
}
