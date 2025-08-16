package lsp

// ChangeType represents the change type.
type ChangeType int

// ChangeTypeFull represents the full change type.
// For more information, see the [specification]:
//
// [specification]: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentSyncKind
const (
	ChangeTypeFull ChangeType = iota + 1
	ChangeTypeIncremental
)
