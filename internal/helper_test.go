package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockParseResponse struct {
	ParseHeadersCalled bool
}

func (m *MockParseResponse) ParseHeaders(headers http.Header) {
	m.ParseHeadersCalled = true
}

type Test[T any] struct {
	Title string
	// Setup to make needed modifications for a specific test
	Setup func()
	// Cleanup to undo changes do to reusable items
	Cleanup     func()
	Input       interface{}
	Expected    T
	ExpectedErr error
}

type TestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestIsNil(t *testing.T) {
	testCases := []Test[bool]{
		{
			Title:       "Map",
			Input:       nil,
			ExpectedErr: nil,
			Expected:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			output := IsNil(tc.Input)
			assert.Equal(t, tc.Expected, output)
		})
	}
}
