// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package tableapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestTableRequestBuilder_GetSuccess covers Get's success path, including the ParseHeaders call
// that turns the inspected response headers into the collection's pagination links. The
// existing Get tests only cover its failure branches.
func TestTableRequestBuilder_GetSuccess(t *testing.T) {
	adapter := new(mocking.MockRequestAdapter)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(core.NewBaseServiceNowCollectionResponse[*TableRecord](CreateTableRecordFromDiscriminatorValue), nil)

	builder := NewDefaultTableRequestBuilder(headTestRawURL, adapter)

	response, err := builder.Get(context.Background(), nil)

	require.NoError(t, err)
	require.NotNil(t, response)
	adapter.AssertExpectations(t)
}

// TestTableRequestBuilder_GetRegistersHeaderInspection guards the option Get depends on to
// populate pagination links: without response-header inspection, ParseHeaders has nothing to
// read and paging silently stops working.
func TestTableRequestBuilder_GetRegistersHeaderInspection(t *testing.T) {
	adapter := new(mocking.MockRequestAdapter)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(core.NewBaseServiceNowCollectionResponse[*TableRecord](CreateTableRecordFromDiscriminatorValue), nil)

	builder := NewDefaultTableRequestBuilder(headTestRawURL, adapter)

	config := &TableRequestBuilderGetRequestConfiguration{}
	_, err := builder.Get(context.Background(), config)

	require.NoError(t, err)
	assert.NotEmpty(t, config.Options, "Get must register a header-inspection option")
}

// TestTableRequestBuilder_Post covers Post's outcomes past its nil-body guard.
func TestTableRequestBuilder_PostOutcomes(t *testing.T) {
	tests := []struct {
		name       string
		adapter    func() *mocking.MockRequestAdapter
		wantErr    error
		wantErrMsg string
	}{
		{
			name: "happy path returns the item response",
			adapter: func() *mocking.MockRequestAdapter {
				adapter := newWritingAdapter()
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*TableRecord](CreateTableRecordFromDiscriminatorValue), nil)

				return adapter
			},
		},
		{
			name:    "body serialization failure",
			adapter: newFailingBodyAdapter,
			wantErr: errTransport,
		},
		{
			name: "adapter error propagates",
			adapter: func() *mocking.MockRequestAdapter {
				adapter := newWritingAdapter()
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errTransport)

				return adapter
			},
			wantErr: errTransport,
		},
		{
			name: "nil response returns ErrNilResponse",
			adapter: func() *mocking.MockRequestAdapter {
				adapter := newWritingAdapter()
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

				return adapter
			},
			wantErr: snerrors.ErrNilResponse,
		},
		{
			name: "wrong response type",
			adapter: func() *mocking.MockRequestAdapter {
				adapter := newWritingAdapter()
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(mocking.NewMockParsable(), nil)

				return adapter
			},
			wantErrMsg: "resp is not",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			builder := NewDefaultTableRequestBuilder(headTestRawURL, test.adapter())

			response, err := builder.Post(context.Background(), NewTableRecord(), nil)

			switch {
			case test.wantErr != nil:
				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, response)
			case test.wantErrMsg != "":
				require.ErrorContains(t, err, test.wantErrMsg)
				assert.Nil(t, response)
			default:
				require.NoError(t, err)
				assert.NotNil(t, response)
			}
		})
	}
}

// TestTableRequestBuilder_ToRequestInformationNilGuards covers the nil-receiver guard on each
// of the three request-information builders.
func TestTableRequestBuilder_ToRequestInformationNilGuards(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		call func(builder *TableRequestBuilder[*TableRecord]) (any, error)
	}{
		{
			name: "ToGetRequestInformation",
			call: func(b *TableRequestBuilder[*TableRecord]) (any, error) {
				return b.ToGetRequestInformation(ctx, nil)
			},
		},
		{
			name: "ToPostRequestInformation",
			call: func(b *TableRequestBuilder[*TableRecord]) (any, error) {
				return b.ToPostRequestInformation(ctx, NewTableRecord(), nil)
			},
		},
		{
			name: "ToHeadRequestInformation",
			call: func(b *TableRequestBuilder[*TableRecord]) (any, error) {
				return b.ToHeadRequestInformation(ctx, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name+"/nil builder", func(t *testing.T) {
			var builder *TableRequestBuilder[*TableRecord]

			requestInfo, err := test.call(builder)

			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, requestInfo)
		})

		t.Run(test.name+"/nil inner request builder", func(t *testing.T) {
			builder := &TableRequestBuilder[*TableRecord]{}

			requestInfo, err := test.call(builder)

			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, requestInfo)
		})
	}
}

// TestTableRequestBuilder_ToPostRequestInformationSerializationFailure covers the body-write
// error branch inside ToPostRequestInformation.
func TestTableRequestBuilder_ToPostRequestInformationSerializationFailure(t *testing.T) {
	builder := NewDefaultTableRequestBuilder(headTestRawURL, newFailingBodyAdapter())

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), NewTableRecord(), nil)

	require.ErrorIs(t, err, errTransport)
	assert.Nil(t, requestInfo)
}
