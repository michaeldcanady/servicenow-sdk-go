package http

import (
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

var (
	sharedBaseHandler = &BaseHandler{}
)

func TestBaseHandler_SetNext(t *testing.T) {
	tests := []mocking.Test[*BaseHandler]{
		{
			Title:       "Valid",
			Input:       &BaseHandler{},
			Expected:    &BaseHandler{},
			ExpectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			sharedBaseHandler.SetNext(test.Input.(RequestHandler))
			assert.Equal(t, test.Expected, sharedBaseHandler.next)
		})
	}
}

func TestBaseHandler_Handle(t *testing.T) {
	tests := []mocking.Test[*BaseHandler]{
		{
			Title: "Valid",
			Input: &mocking.MockRequestInformation{},
			Setup: func() {
				sharedBaseHandler.SetNext(&BaseHandler{})
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
