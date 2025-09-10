package errorutil

// Stage represents a stage in the compilation process.
type Stage int

const (
	// StageTokenization represents the tokenization stage.
	StageTokenization Stage = iota
	// StageParsing represents the parsing stage.
	StageParsing
	// StageEvaluation represents the evaluation stage.
	StageEvaluation
)
