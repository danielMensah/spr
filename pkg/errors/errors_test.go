package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCauseError_Context(t *testing.T) {
	tests := []struct {
		name        string
		causeError  causeError
		expectedErr error
	}{
		{
			name: "returns context",
			causeError: causeError{
				contextError: ContextError{
					message: "context for error",
				},
				cause: errors.New("an error"),
			},
			expectedErr: ContextError{
				message: "context for error",
			},
		},
		{
			name: "empty context",
			causeError: causeError{
				contextError: ContextError{},
				cause:        errors.New("an error"),
			},
			expectedErr: ContextError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.causeError.Context()

			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestCauseError_Error(t *testing.T) {
	tests := []struct {
		name       string
		causeError causeError
		expected   string
	}{
		{
			name: "constructs error string",
			causeError: causeError{
				contextError: ContextError{
					message: "context for error",
				},
				cause: errors.New("an error"),
			},
			expected: "context for error: an error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.causeError.Error()

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestContext(t *testing.T) {
	tests := []struct {
		name               string
		err                error
		expectedContextErr error
	}{
		{
			name: "cause error",
			err: &causeError{
				contextError: ContextError{
					message: "context for error",
				},
				cause: errors.New("cause error"),
			},
			expectedContextErr: ContextError{
				message: "context for error",
			},
		},
		{
			name: "context error",
			err: ContextError{
				message: "context for error",
			},
			expectedContextErr: ContextError{
				message: "context for error",
			},
		},
		{
			name:               "normal error",
			err:                errors.New("context for error"),
			expectedContextErr: errors.New("context for error"),
		},
		{
			name:               "nil error",
			err:                nil,
			expectedContextErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Context(tt.err)

			assert.Equal(t, tt.expectedContextErr, err)
		})
	}
}

func TestContextError_Error(t *testing.T) {
	tests := []struct {
		name         string
		contextError ContextError
		expected     string
	}{
		{
			name: "non-empty",
			contextError: ContextError{
				"context for error",
			},
			expected: "context for error",
		},
		{
			name: "empty",
			contextError: ContextError{
				"",
			},
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.contextError.Error()

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestContextError_WithCause(t *testing.T) {
	tests := []struct {
		name         string
		contextError ContextError
		cause        error
		expectedErr  error
	}{
		{
			name: "contains an error",
			contextError: ContextError{
				message: "context for error",
			},
			cause: errors.New("an error"),
			expectedErr: &causeError{
				contextError: ContextError{
					message: "context for error",
				},
				cause: errors.New("an error"),
			},
		},
		{
			name: "missing an error",
			contextError: ContextError{
				message: "context for error",
			},
			cause: nil,
			expectedErr: &causeError{
				contextError: ContextError{
					message: "context for error",
				},
				cause: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.contextError.WithCause(tt.cause)

			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected ContextError
	}{
		{
			name: "creates new context error",
			msg:  "context for error",
			expected: ContextError{
				message: "context for error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.msg)

			assert.Equal(t, tt.expected, got)
		})
	}
}
