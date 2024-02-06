package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sharedReversePageIterator = NewReversePageIterator(sharedForwardPageIterator)
)

func TestReversePageIterator_Iterate(t *testing.T) {
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
				sharedReversePageIterator.Current().PreviousPageLink = fakeLinkNilResp
			},
			cleanup: func() {
				sharedReversePageIterator.Current().PreviousPageLink = ""
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
				sharedReversePageIterator.Current().PreviousPageLink = fakeNextLink
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
			sharedReversePageIterator.pauseIndex = len(sharedReversePageIterator.Current().Result)

			if tt.setup != nil {
				tt.setup()
			}

			err := sharedReversePageIterator.Reverse(tt.input.(func(*person) bool))

			assert.Equal(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, count)

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

func TestReversePageIterator_enumerate(t *testing.T) {
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
				sharedReversePageIterator.Current().Result = nil
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
			sharedReversePageIterator.pauseIndex = 0 // Reset the pause index between runs

			if tt.setup != nil {
				tt.setup()
			}

			keepIterating := sharedReversePageIterator.enumerate(tt.input.(func(item *person) bool))

			assert.ErrorIs(t, nil, tt.expectedErr)
			assert.Equal(t, tt.expected, []interface{}{count, keepIterating})
		})
	}
}

func TestReversePageIterator_Previous(t *testing.T) {
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
			pageResult, err := sharedReversePageIterator.Previous()

			assert.ErrorIs(t, err, tt.expectedErr)

			assert.Equal(t, tt.expected, pageResult)
		})
	}
}

func TestPageIterator_Last(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title: "Valid",
			setup: func() {
				sharedReversePageIterator.Current().LastPageLink = fakeLastLink
			},
			input:       nil,
			expected:    sharedPageResult,
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			pageResult, err := sharedReversePageIterator.Last()

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, pageResult)
		})
	}
}
