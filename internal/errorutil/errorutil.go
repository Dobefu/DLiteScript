// Package errorutil provides utility functions for handling errors.
package errorutil

import (
	"errors"
	"fmt"
)

// ErrorMsg represents a predefined error message.
type ErrorMsg string

const (
	// ErrorMsgEmptyExpression occurs when an expression is empty.
	ErrorMsgEmptyExpression = "empty expression"
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
	// ErrorMsgUnknownOperator occurs when an unknown operator is encountered.
	ErrorMsgUnknownOperator = "unknown operator: '%s'"
	// ErrorMsgUnknownNodeType occurs when an unknown node type is encountered.
	ErrorMsgUnknownNodeType = "unknown node type: '%T'"
	// ErrorMsgUnexpectedChar occurs when an unexpected character is encountered.
	ErrorMsgUnexpectedChar = "unexpected character: '%s'"
	// ErrorMsgFunctionNumArgs occurs when a function receives the wrong number of arguments.
	ErrorMsgFunctionNumArgs = "'%s()' expects exactly %d argument(s), but got %d"
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
)

// Error represents an error with a message.
type Error struct {
	msg ErrorMsg
	pos int
}

// NewError creates a new error with the given message.
func NewError(msg ErrorMsg, args ...any) *Error {
	return &Error{
		msg: ErrorMsg(fmt.Sprintf(string(msg), args...)),
		pos: -1,
	}
}

// NewErrorAt creates a new error with the given message at a specific position.
func NewErrorAt(msg ErrorMsg, pos int, args ...any) *Error {
	return &Error{
		msg: ErrorMsg(fmt.Sprintf(string(msg), args...)),
		pos: pos,
	}
}

// Error returns the error message with the position information.
func (e *Error) Error() string {
	// If the position is less than 0, there's no position information to return.
	if e.pos < 0 {
		return string(e.msg)
	}

	return fmt.Sprintf("%s at position %d", e.msg, e.pos)
}

// Unwrap returns the error message without any additional information.
func (e *Error) Unwrap() error {
	return errors.New(string(e.msg))
}

// Position gets the position of the error.
func (e *Error) Position() int {
	return e.pos
}
