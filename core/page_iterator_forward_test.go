package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sharedForwardPageIterator = NewForwardPageIterator(sharedPageIterator)
)

func TestForwardPageIterator_Iterate(t *testing.T) {
	var count int

	tests := []test[int]{
		{
			title:       "Missing Callback",
			input:       (func(*person) bool)(nil),
			expected:    0,
			shouldErr:   true,
			expectedErr: ErrNilCallback,
		},
		{
			title: "Callback false",
			input: func(person *person) bool {
				return false
			},
			expected:    0,
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Missing Response",
			setup: func() {
				sharedForwardPageIterator.Current().NextPageLink = fakeLinkNilResp
			},
			cleanup: func() {
				sharedForwardPageIterator.Current().NextPageLink = ""
			},
			input:       func(person *person) bool { return true },
			expected:    0,
			shouldErr:   true,
			expectedErr: ErrNilResponse,
		},
		{
			title: "Missing Next Link",
			input: func(person *person) bool {
				count += 1

				return true
			},
			expected:    len(sharedCurrentPage.Result),
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Single Page",
			setup: func() {
				sharedForwardPageIterator.Current().NextPageLink = fakeNextLink
			},
			input: func(person *person) bool {
				count += 1

				return true
			},
			expected:    len(sharedCurrentPage.Result) + len(sharedCurrentPage.Result),
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			count = 0
			sharedForwardPageIterator.pauseIndex = 0

			if tt.setup != nil {
				tt.setup()
			}

			err := sharedForwardPageIterator.Iterate(tt.input.(func(*person) bool))

			assert.Equal(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, count)

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestForwardPageIterator_enumerate(t *testing.T) {
	var count int

	tests := []test[[]interface{}]{
		{
			title: "All",
			input: func(item *person) bool {
				count += 1

				return true
			},
			expected:    []interface{}{len(sharedCurrentPage.Result), true},
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Once",
			input: func(item *person) bool {
				count += 1

				return false
			},
			expected:    []interface{}{1, false},
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Missing Page items",
			setup: func() {
				sharedForwardPageIterator.Current().Result = nil
			},
			input:       func(pageItem *person) bool { return true },
			expected:    []interface{}{0, false},
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			count = 0                                // Reset the count between runs
			sharedForwardPageIterator.pauseIndex = 0 // Reset the pause index between runs

			if tt.setup != nil {
				tt.setup()
			}

			keepIterating := sharedForwardPageIterator.enumerate(tt.input.(func(item *person) bool))

			assert.ErrorIs(t, nil, tt.expectedErr)
			assert.Equal(t, tt.expected, []interface{}{count, keepIterating})
		})
	}
}

func TestForwardPageIterator_Next(t *testing.T) {
	tests := []test[PageResult[person]]{
		// TODO: Needs "valid" test
		{
			title:       "Missing Next Link",
			expected:    PageResult[person]{},
			expectedErr: ErrEmptyURI,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			pageResult, err := sharedForwardPageIterator.Next()

			assert.ErrorIs(t, err, tt.expectedErr)

			assert.Equal(t, tt.expected, pageResult)
		})
	}
}
