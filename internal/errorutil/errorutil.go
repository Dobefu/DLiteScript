// Package errorutil provides utility functions for handling errors.
package errorutil

import (
	"errors"
	"fmt"
)

// ErrorMsg represents a predefined error message.
type ErrorMsg string

const (
	// ErrorMsgUnexpectedEOF occurs when EOF is reached while still parsing.
	ErrorMsgUnexpectedEOF = "unexpected end of expression"
	// ErrorMsgInvalidUTF8Char occurs when an invalid UTF-8 sequence is encountered.
	ErrorMsgInvalidUTF8Char = "invalid character in expression"
	// ErrorMsgParenNotClosedAtEOF occurs when a closing parenthesis is expected but EOF is reached.
	ErrorMsgParenNotClosedAtEOF = "expected ')' at end of expression"
	// ErrorMsgDivByZero occurs when attempting to divide by zero.
	ErrorMsgDivByZero = "division by zero"
	// ErrorMsgModByZero occurs when attempting to perform modulo by zero.
	ErrorMsgModByZero = "modulo by zero"
	// ErrorMsgUndefinedIdentifier occurs when an undefined identifier is encountered.
	ErrorMsgUndefinedIdentifier = "undefined identifier: '%s'"
	// ErrorMsgUndefinedFunction occurs when an undefined function is encountered.
	ErrorMsgUndefinedFunction = "undefined function: '%s'"
	// ErrorMsgUnexpectedToken occurs when an unexpected token is encountered.
	ErrorMsgUnexpectedToken = "unexpected token: '%s'"
	// ErrorMsgExpectedOpenParen occurs when an opening parenthesis is expected but not provided.
	ErrorMsgExpectedOpenParen = "expected '(', but got: '%s'"
	// ErrorMsgExpectedCloseParen occurs when a closing parenthesis is expected but not provided.
	ErrorMsgExpectedCloseParen = "expected ')', but got: '%s'"
	// ErrorMsgExpectedCloseBracket occurs when a closing bracket is expected but not provided.
	ErrorMsgExpectedCloseBracket = "expected ']', but got: '%s'"
	// ErrorMsgUnknownOperator occurs when an unknown operator is encountered.
	ErrorMsgUnknownOperator = "unknown operator: '%s'"
	// ErrorMsgUnknownNodeType occurs when an unknown node type is encountered.
	ErrorMsgUnknownNodeType = "unknown node type: '%T'"
	// ErrorMsgUnexpectedChar occurs when an unexpected character is encountered.
	ErrorMsgUnexpectedChar = "unexpected character: '%s'"
	// ErrorMsgFunctionNumArgs occurs when a function receives the wrong number of arguments.
	ErrorMsgFunctionNumArgs = "'%s()' expects exactly %d argument(s), but got %d"
	// ErrorMsgFunctionArgType occurs when a function receives an argument of the wrong type.
	ErrorMsgFunctionArgType = "'%s()' expects argument %d to be '%s', but got '%s'"
	// ErrorMsgFunctionReturnCount occurs when a function returns the wrong number of values.
	ErrorMsgFunctionReturnCount = "'%s()' expects to return %d value(s), but returned %d"
	// ErrorMsgNumberTrailingChar occurs when a number has non-numeric trailing characters.
	ErrorMsgNumberTrailingChar = "trailing character in number: '%s'"
	// ErrorMsgNumberMultipleUnderscores occurs when a number has multiple consecutive underscores.
	ErrorMsgNumberMultipleUnderscores = "multiple consecutive underscores in number: '%s'"
	// ErrorMsgNumberMultipleDecimalPoints occurs when a number has multiple decimal points.
	ErrorMsgNumberMultipleDecimalPoints = "multiple decimal points in number: '%s'"
	// ErrorMsgNumberMultipleExponentSigns occurs when a number has multiple exponent signs.
	ErrorMsgNumberMultipleExponentSigns = "multiple exponent signs in number: '%s'"
	// ErrorMsgNumberMultipleConsecutiveExponentSigns occurs when an exponent has multiple consecutive signs.
	ErrorMsgNumberMultipleConsecutiveExponentSigns = "multiple consecutive addition or subtraction signs in exponent: '%s'"
	// ErrorMsgTypeUnknownDataType occurs when an unknown data type is encountered.
	ErrorMsgTypeUnknownDataType = "type error: unknown data type: '%T'"
	// ErrorMsgTypeExpected occurs when a type is expected but a different type is encountered.
	ErrorMsgTypeExpected = "type error: expected %s, but got %s"
	// ErrorMsgTypeMismatch occurs when there's a type mismatch.
	ErrorMsgTypeMismatch = "expected %s, got %s"
	// ErrorMsgUnexpectedIdentifier occurs when an unexpected identifier is encountered.
	ErrorMsgUnexpectedIdentifier = "unexpected identifier: '%s'"
	// ErrorMsgInvalidDataType occurs when an invalid data type is encountered.
	ErrorMsgInvalidDataType = "invalid data type: '%s'"
	// ErrorMsgConstantDeclarationWithNoValue occurs when a constant declaration has no value.
	ErrorMsgConstantDeclarationWithNoValue = "constant declaration '%s' must have a value"
	// ErrorMsgReassignmentToConstant occurs when trying to re-assign a value to a constant.
	ErrorMsgReassignmentToConstant = "cannot re-assign value to constant: '%s'"
	// ErrorMsgBlockStatementExpected occurs when a block statement is expected but a different type is encountered.
	ErrorMsgBlockStatementExpected = "block statement expected, but got '%T'"
	// ErrorMsgBreakCountLessThanOne occurs when a break count is less than 1.
	ErrorMsgBreakCountLessThanOne = "break count must be greater than 0"
	// ErrorMsgContinueCountLessThanOne occurs when a continue count is less than 1.
	ErrorMsgContinueCountLessThanOne = "continue count must be greater than 0"
	// ErrorMsgVariableNotFound occurs when a variable is not found.
	ErrorMsgVariableNotFound = "variable not found: '%s'"
	// ErrorMsgInvalidForStatement occurs when a for statement is invalid.
	ErrorMsgInvalidForStatement = "invalid for statement: '%s'"
	// ErrorMsgInvalidNumber occurs when an invalid number is encountered.
	ErrorMsgInvalidNumber = "invalid number: '%s'"
	// ErrorMsgUndefinedNamespace occurs when an undefined namespace is encountered.
	ErrorMsgUndefinedNamespace = "undefined namespace: '%s'"
	// ErrorMsgArrayIndexOutOfBounds occurs when an array index is out of bounds.
	ErrorMsgArrayIndexOutOfBounds = "array index out of bounds: '%s'"
	// ErrorMsgCannotConcat occurs when two values of the same type cannot be concatenated.
	ErrorMsgCannotConcat = "cannot concatenate %s and %s"
)

// Error represents an error with a message.
type Error struct {
	msg   ErrorMsg
	pos   int
	stage Stage
}

// NewError creates a new error with the given message.
func NewError(phase Stage, msg ErrorMsg, args ...any) *Error {
	return &Error{
		msg:   ErrorMsg(fmt.Sprintf(string(msg), args...)),
		pos:   -1,
		stage: phase,
	}
}

// NewErrorAt creates a new error with the given message at a specific position.
func NewErrorAt(phase Stage, msg ErrorMsg, pos int, args ...any) *Error {
	return &Error{
		msg:   ErrorMsg(fmt.Sprintf(string(msg), args...)),
		pos:   pos,
		stage: phase,
	}
}

// Error returns the error message with the position information.
func (e *Error) Error() string {
	msg := fmt.Sprintf("%s: %s", e.stage.String(), string(e.msg))

	// If the position is less than 0, there's no position information to return.
	if e.pos < 0 {
		return msg
	}

	return fmt.Sprintf("%s at position %d", msg, e.pos)
}

// Unwrap returns the error message without any additional information.
func (e *Error) Unwrap() error {
	return errors.New(string(e.msg))
}

// Position gets the position of the error.
func (e *Error) Position() int {
	return e.pos
}
