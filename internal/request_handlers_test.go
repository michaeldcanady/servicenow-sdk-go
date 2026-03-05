package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseHandler_SetNext(t *testing.T) {
	tests := []struct {
		name     string
		input    RequestHandler
		expected RequestHandler
	}{
		{
			name:     "Valid",
			input:    &BaseHandler{},
			expected: &BaseHandler{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := &BaseHandler{}
			handler.SetNext(test.input)
			assert.Equal(t, test.expected, handler.next)
		})
	}
}

func TestBaseHandler_Handle(t *testing.T) {
	tests := []struct {
		name        string
		input       RequestInformation
		next        RequestHandler
		expectedErr error
	}{
		{
			name:        "Valid",
			input:       &MockRequestInformation{},
			next:        &BaseHandler{},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := &BaseHandler{}
			if test.next != nil {
				handler.SetNext(test.next)
			}
			err := handler.Handle(test.input)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}
