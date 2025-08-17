package lsptypes

// DidChangeParams represents the parameters for a didChange request.
type DidChangeParams struct {
	ContentChanges []ContentChange  `json:"contentChanges"`
	TextDocument   TextDocumentItem `json:"textDocument"`
}

// ContentChange represents a content change.
type ContentChange struct {
	Range       *Range `json:"range"`
	RangeLength int    `json:"rangeLength"`
	Text        string `json:"text"`
}
