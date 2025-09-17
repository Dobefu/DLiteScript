// Package controlflow provides functionality for controlling the flow of execution.
package controlflow

import "github.com/Dobefu/DLiteScript/internal/datavalue"

// FlowType represents the type of a flow control statement.
type FlowType int

const (
	// FlowTypeBreak represents a break flow control.
	FlowTypeBreak FlowType = iota
	// FlowTypeContinue represents a continue flow control.
	FlowTypeContinue
	// FlowTypeReturn represents a return flow control.
	FlowTypeReturn
	// FlowTypeExit represents an exit flow control.
	FlowTypeExit
)

// Control represents a control flow from a statement.
type Control struct {
	Count int
	Type  FlowType
}

// EvaluationResult represents a result from a statement.
type EvaluationResult struct {
	Value   datavalue.Value
	Control *Control
}

// NewRegularResult creates a new regular result.
func NewRegularResult(value datavalue.Value) *EvaluationResult {
	return &EvaluationResult{
		Value:   value,
		Control: nil,
	}
}

// NewBreakResult creates a new break result.
func NewBreakResult(count int) *EvaluationResult {
	return &EvaluationResult{
		Value: datavalue.Null(),
		Control: &Control{
			Type:  FlowTypeBreak,
			Count: count,
		},
	}
}

// NewContinueResult creates a new continue result.
func NewContinueResult(count int) *EvaluationResult {
	return &EvaluationResult{
		Value: datavalue.Null(),
		Control: &Control{
			Type:  FlowTypeContinue,
			Count: count,
		},
	}
}

// NewReturnResult creates a new return result.
func NewReturnResult(value datavalue.Value) *EvaluationResult {
	return &EvaluationResult{
		Value: value,
		Control: &Control{
			Type:  FlowTypeReturn,
			Count: 0,
		},
	}
}

// NewExitResult creates a new exit result.
func NewExitResult(code byte) *EvaluationResult {
	return &EvaluationResult{
		Value: datavalue.Null(),
		Control: &Control{
			Type:  FlowTypeExit,
			Count: int(code),
		},
	}
}

// IsNormalResult returns true if this is a normal result (no control flow).
func (r *EvaluationResult) IsNormalResult() bool {
	return r.Control == nil
}

// IsBreakResult returns true if this is a break result.
func (r *EvaluationResult) IsBreakResult() bool {
	return r.Control != nil && r.Control.Type == FlowTypeBreak
}

// IsContinueResult returns true if this is a continue result.
func (r *EvaluationResult) IsContinueResult() bool {
	return r.Control != nil && r.Control.Type == FlowTypeContinue
}

// IsReturnResult returns true if this is a return result.
func (r *EvaluationResult) IsReturnResult() bool {
	return r.Control != nil && r.Control.Type == FlowTypeReturn
}

// IsExitResult returns true if this is an exit result.
func (r *EvaluationResult) IsExitResult() bool {
	return r.Control != nil && r.Control.Type == FlowTypeExit
}
