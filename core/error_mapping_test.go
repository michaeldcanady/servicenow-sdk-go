package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMappingSet(t *testing.T) {
	tests := []struct {
		name           string
		key            string
		value          string
		expectedLength int
	}{
		{
			name:           "Set 404",
			key:            "404",
			value:          "Not Found",
			expectedLength: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			eM := NewErrorMapping()
			eM.Set(test.key, test.value)

			assert.Equal(t, test.expectedLength, eM.Len())

			msg, found := eM.Get(404)
			assert.True(t, found)
			assert.Equal(t, test.value, msg)
		})
	}
}

func TestErrorMappingGet(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(eM ErrorMapping)
		statusCode    int
		expectedFound bool
		expectedMsg   string
	}{
		{
			name: "Exact match 404",
			setup: func(eM ErrorMapping) {
				eM.Set("404", "Not Found")
			},
			statusCode:    404,
			expectedFound: true,
			expectedMsg:   "Not Found",
		},
		{
			name: "Relative match 400",
			setup: func(eM ErrorMapping) {
				eM.Set("4XX", "Client Error")
			},
			statusCode:    400,
			expectedFound: true,
			expectedMsg:   "Client Error",
		},
		{
			name: "Relative match 502",
			setup: func(eM ErrorMapping) {
				eM.Set("5XX", "Server Error")
			},
			statusCode:    502,
			expectedFound: true,
			expectedMsg:   "Server Error",
		},
		{
			name:          "Not found 500",
			setup:         func(eM ErrorMapping) {},
			statusCode:    500,
			expectedFound: false,
			expectedMsg:   "",
		},
		{
			name:          "Non-error status 200",
			setup:         func(eM ErrorMapping) {},
			statusCode:    200,
			expectedFound: false,
			expectedMsg:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			eM := NewErrorMapping()
			test.setup(eM)
			msg, found := eM.Get(test.statusCode)
			assert.Equal(t, test.expectedFound, found)
			assert.Equal(t, test.expectedMsg, msg)
		})
	}
}
