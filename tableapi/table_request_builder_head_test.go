// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package tableapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const headTestRawURL = "https://example.com/api/now/v1/table/test"

// errTransport is a stand-in for a transport-level failure from the adapter.
var errTransport = errors.New("transport error")

// TestTableRequestBuilder_Head covers the request lifecycle past the nil guards, which are
// covered separately in TestTableRequestBuilder_Head_NilGuards. Head is unusual in this
// package: it goes through SendNoContent and returns the inspected response headers rather
// than a deserialized body.
func TestTableRequestBuilder_Head(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path returns inspected response headers",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errTransport)
			},
			wantErr: errTransport,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			adapter := new(mocking.MockRequestAdapter)
			test.setupMock(adapter)

			builder := NewDefaultTableRequestBuilder(headTestRawURL, adapter)

			headers, err := builder.Head(context.Background(), nil)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, headers)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, headers)
			}

			adapter.AssertExpectations(t)
		})
	}
}

// TestTableRequestBuilder_HeadRegistersHeaderInspection guards the behaviour Head depends on:
// it must append a HeadersInspectionOptions with response inspection enabled, otherwise the
// headers it returns would always be empty.
func TestTableRequestBuilder_HeadRegistersHeaderInspection(t *testing.T) {
	adapter := new(mocking.MockRequestAdapter)
	adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	builder := NewDefaultTableRequestBuilder(headTestRawURL, adapter)

	config := &TableRequestBuilderGetRequestConfiguration{}
	headers, err := builder.Head(context.Background(), config)

	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.NotEmpty(t, config.Options, "Head must register a header-inspection option")
}

// TestNewDefaultTableRequestBuilder checks the default-parsable convenience constructor wires
// the raw URL through as a path parameter.
func TestNewDefaultTableRequestBuilder(t *testing.T) {
	adapter := new(mocking.MockRequestAdapter)

	builder := NewDefaultTableRequestBuilder(headTestRawURL, adapter)

	require.NotNil(t, builder)
	assert.Equal(t, headTestRawURL, builder.GetPathParameters()["request-raw-url"])
	assert.Equal(t, adapter, builder.GetRequestAdapter())
}
