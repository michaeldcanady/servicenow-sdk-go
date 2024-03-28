package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sharedBaseHandler = &BaseHandler[any]{}
)

func TestBaseHandler_SetNext(t *testing.T) {
	tests := []Test[*BaseHandler[any]]{
		{
			Title:       "Valid",
			Input:       &BaseHandler[any]{},
			Expected:    &BaseHandler[any]{},
			ExpectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			sharedBaseHandler.SetNext(test.Input.(Handler[any]))
			assert.Equal(t, test.Expected, sharedBaseHandler.next)
		})
	}
}

func TestBaseHandler_Handle(t *testing.T) {
	tests := []Test[*BaseHandler[any]]{
		{
			Title: "Valid",
			Input: &MockRequestInformation{},
			Setup: func() {
				sharedBaseHandler.SetNext(&BaseHandler[any]{})
			},
			Expected:    nil,
			ExpectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			err := sharedBaseHandler.Handle(test.Input.(RequestInformation))
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}
