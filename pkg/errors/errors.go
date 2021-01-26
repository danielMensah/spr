package errors

import (
	"fmt"
)

// ContextError defines a error struct for context purposes
type ContextError struct {
	message string
}

type contextInterface interface {
	Context() error
}

// Error implements the built-in error interface
func (c ContextError) Error() string {
	return c.message
}

type causeError struct {
	contextError ContextError
	cause        error
}

// Context returns the context from the cause
func (c causeError) Context() error {
	return c.contextError
}

// Error implements the built-in error interface
func (c causeError) Error() string {
	return fmt.Sprintf("%s: %s", c.contextError, c.cause)
}

// New creates a new contextError with the passed message
func New(msg string) ContextError {
	return ContextError{
		message: msg,
	}
}

// WithCause creates a new causeError with the current context as the parent context
func (c ContextError) WithCause(cause error) error {
	return &causeError{
		contextError: c,
		cause:        cause,
	}
}

// Context is a helper function which tries to see if the passed error implements contextInterface and if so returns the
// parent context. Otherwise it will return the passed error
func Context(err error) error {
	if context, ok := err.(contextInterface); ok {
		return context.Context()
	}

	return err
}
