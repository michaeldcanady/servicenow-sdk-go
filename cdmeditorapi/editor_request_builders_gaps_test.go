// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmeditorapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const editorTestBaseURL = "https://example.service-now.com"

// errSend is a stand-in for a transport-level failure from the request adapter.
var errSend = errors.New("send failed")

func editorTestPathParameters() map[string]string {
	return map[string]string{"baseurl": editorTestBaseURL}
}

// editorTestOptions returns a single request option ready to be stored on a RequestInformation.
func editorTestOptions() []abstractions.RequestOption {
	option := mocking.NewMockRequestOption()
	option.On("GetKey").Return(abstractions.RequestOptionKey{Key: "test-option"})

	return []abstractions.RequestOption{option}
}

// newEditorAdapter returns an adapter that can serialize a body and whose Send returns the
// supplied response.
func newEditorAdapter(response any, sendErr error) *mocking.MockRequestAdapter {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(response, sendErr)
	adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(sendErr)

	return adapter
}

func TestNodesRequestBuilder_ByID(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	parent := NewNodesRequestBuilderInternal(editorTestPathParameters(), adapter)

	child := parent.ByID("d71f7935c0a80167")

	require.NotNil(t, child)
	assert.Equal(t, "d71f7935c0a80167", child.GetPathParameters()["node_sys_id"])
	assert.Equal(t, editorTestBaseURL, child.GetPathParameters()["baseurl"])
	assert.Equal(t, adapter, child.GetRequestAdapter())

	// The parent must not gain the node ID the child added.
	assert.NotContains(t, parent.GetPathParameters(), "node_sys_id")
}

// TestEditorRequestBuilders_ConfigurationIsApplied covers the inline request-configuration
// handling in each verb. These builders apply Headers, Options and QueryParameters by hand
// rather than going through abstractions.ConfigureRequestInformation, so each branch needs
// exercising separately.
func TestEditorRequestBuilders_ConfigurationIsApplied(t *testing.T) {
	ctx := context.Background()

	t.Run("Nodes Get applies headers, options and query parameters", func(t *testing.T) {
		adapter := newEditorAdapter(core.NewBaseServiceNowCollectionResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil)
		builder := NewNodesRequestBuilderInternal(editorTestPathParameters(), adapter)

		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &NodesRequestBuilderGetRequestConfiguration{
			Headers:         headers,
			Options:         editorTestOptions(),
			QueryParameters: &NodesRequestBuilderGetQueryParameters{SysID: internal.ToPointer("123")},
		}

		response, err := builder.Get(ctx, config)

		require.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("Nodes Post applies headers and options", func(t *testing.T) {
		adapter := newEditorAdapter(core.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil)
		builder := NewNodesRequestBuilderInternal(editorTestPathParameters(), adapter)

		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &NodesRequestBuilderPostRequestConfiguration{
			Headers: headers,
			Options: editorTestOptions(),
		}

		response, err := builder.Post(ctx, NewNodeCreateRequest(), config)

		require.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("NodeItem Put applies headers and options", func(t *testing.T) {
		adapter := newEditorAdapter(core.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue), nil)
		builder := NewNodeItemRequestBuilderInternal(editorTestPathParameters(), adapter)

		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &NodeItemRequestBuilderPutRequestConfiguration{
			Headers: headers,
			Options: editorTestOptions(),
		}

		response, err := builder.Put(ctx, NewNodeUpdateRequest(), config)

		require.NoError(t, err)
		assert.NotNil(t, response)
	})

	t.Run("NodeItem Delete applies headers and options", func(t *testing.T) {
		adapter := newEditorAdapter(nil, nil)
		builder := NewNodeItemRequestBuilderInternal(editorTestPathParameters(), adapter)

		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &NodeItemRequestBuilderDeleteRequestConfiguration{
			Headers: headers,
			Options: editorTestOptions(),
		}

		require.NoError(t, builder.Delete(ctx, config))
	})

	t.Run("Validation Get applies headers, options and query parameters", func(t *testing.T) {
		adapter := newEditorAdapter(core.NewBaseServiceNowItemResponse[*ValidationResultModel](CreateValidationResultFromDiscriminatorValue), nil)
		builder := NewValidationRequestBuilderInternal(editorTestPathParameters(), adapter)

		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &ValidationRequestBuilderGetRequestConfiguration{
			Headers:         headers,
			Options:         editorTestOptions(),
			QueryParameters: &ValidationRequestBuilderGetQueryParameters{},
		}

		response, err := builder.Get(ctx, config)

		require.NoError(t, err)
		assert.NotNil(t, response)
	})
}

// TestEditorRequestBuilders_AdapterErrorPropagates covers the Send-error branch of every verb.
func TestEditorRequestBuilders_AdapterErrorPropagates(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		call func(adapter *mocking.MockRequestAdapter) (any, error)
	}{
		{
			name: "Nodes Get",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				return NewNodesRequestBuilderInternal(editorTestPathParameters(), adapter).Get(ctx, nil)
			},
		},
		{
			name: "Nodes Post",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				return NewNodesRequestBuilderInternal(editorTestPathParameters(), adapter).Post(ctx, NewNodeCreateRequest(), nil)
			},
		},
		{
			name: "NodeItem Put",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				return NewNodeItemRequestBuilderInternal(editorTestPathParameters(), adapter).Put(ctx, NewNodeUpdateRequest(), nil)
			},
		},
		{
			name: "NodeItem Delete",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				// Delete reports only an error, so there is no response to hand back.
				return nil, NewNodeItemRequestBuilderInternal(editorTestPathParameters(), adapter).Delete(ctx, nil)
			},
		},
		{
			name: "Validation Get",
			call: func(adapter *mocking.MockRequestAdapter) (any, error) {
				return NewValidationRequestBuilderInternal(editorTestPathParameters(), adapter).Get(ctx, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := test.call(newEditorAdapter(nil, errSend))

			require.ErrorIs(t, err, errSend)
			assert.Nil(t, response)
		})
	}
}

// TestEditorRequestBuilders_SerializationFailure covers the body-write error branch of the two
// verbs that carry a request body.
func TestEditorRequestBuilders_SerializationFailure(t *testing.T) {
	ctx := context.Background()

	newFailingAdapter := func() *mocking.MockRequestAdapter {
		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(errSend)
		writer.On("Close").Return(nil)

		factory := mocking.NewMockSerializationWriterFactory()
		factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

		adapter := mocking.NewMockRequestAdapter()
		adapter.On("GetSerializationWriterFactory").Return(factory)

		return adapter
	}

	t.Run("Nodes Post", func(t *testing.T) {
		adapter := newFailingAdapter()

		response, err := NewNodesRequestBuilderInternal(editorTestPathParameters(), adapter).Post(ctx, NewNodeCreateRequest(), nil)

		require.ErrorIs(t, err, errSend)
		assert.Nil(t, response)
		adapter.AssertNotCalled(t, "Send")
	})

	t.Run("NodeItem Put", func(t *testing.T) {
		adapter := newFailingAdapter()

		response, err := NewNodeItemRequestBuilderInternal(editorTestPathParameters(), adapter).Put(ctx, NewNodeUpdateRequest(), nil)

		require.ErrorIs(t, err, errSend)
		assert.Nil(t, response)
		adapter.AssertNotCalled(t, "Send")
	})
}

// TestValidationRequestBuilder_NilInnerBuilder covers the second half of the nil-guard
// condition, where the embedded RequestBuilder is nil rather than the outer builder.
func TestValidationRequestBuilder_NilInnerBuilder(t *testing.T) {
	response, err := (&ValidationRequestBuilder{}).Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, response)
}
