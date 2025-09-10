package errorutil

// Stage represents a stage in the compilation process.
type Stage int

const (
	// StageTokenize represents the tokenization stage.
	StageTokenize Stage = iota
	// StageParse represents the parsing stage.
	StageParse
	// StageEvaluate represents the evaluation stage.
	StageEvaluate
)

func (s Stage) String() string {
	switch s {
	case StageTokenize:
		return "tokenize"

	case StageParse:
		return "parse"

	case StageEvaluate:
		return "evaluate"

	default:
		return "unknown stage"
	}
}
