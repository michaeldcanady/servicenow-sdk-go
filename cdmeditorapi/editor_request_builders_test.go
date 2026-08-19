package cdmeditorapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCdmEditorRequestBuilder_Hierarchy(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewCdmEditorRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	assert.NotNil(t, builder.Nodes())
	assert.NotNil(t, builder.Validation())
}

func TestNodesRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	t.Run("URI Construction", func(t *testing.T) {
		sysID := "123"
		config := &NodesRequestBuilderGetRequestConfiguration{
			QueryParameters: &NodesRequestBuilderGetQueryParameters{
				SysID: &sysID,
			},
		}
		requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
		requestInfo.AddQueryParameters(*config.QueryParameters)

		uri, _ := requestInfo.GetUri()
		assert.Equal(t, "https://example.service-now.com/api/sn_cdm/editor/v1/nodes?sys_id=123", uri.String())
	})

	t.Run("Execution", func(t *testing.T) {
		mockRes := core.NewBaseServiceNowCollectionResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue)
		adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

		resp, err := builder.Get(context.Background(), nil)

		require.NoError(t, err)
		assert.Equal(t, mockRes, resp)
	})

	t.Run("Nil response", func(t *testing.T) {
		nilAdapter := mocking.NewMockRequestAdapter()
		nilBuilder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nilAdapter)
		nilAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		resp, err := nilBuilder.Get(context.Background(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilResponse)
		assert.Nil(t, resp)
	})
}

func TestNodesRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewNodeCreateRequest(), nil)

	require.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestNodesRequestBuilder_Post_NilResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

	resp, err := builder.Post(context.Background(), NewNodeCreateRequest(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilResponse)
	assert.Nil(t, resp)
}

func TestNodeItemRequestBuilder_Put(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewNodeItemRequestBuilderInternal(map[string]string{
		"baseurl":     "https://example.service-now.com",
		"node_sys_id": "node123",
	}, adapter)

	t.Run("URI Construction", func(t *testing.T) {
		requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, builder.GetURLTemplate(), builder.GetPathParameters())
		uri, _ := requestInfo.GetUri()
		assert.Equal(t, "https://example.service-now.com/api/sn_cdm/editor/v1/nodes/node123", uri.String())
	})

	t.Run("Execution", func(t *testing.T) {
		mockRes := core.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue)
		adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

		resp, err := builder.Put(context.Background(), NewNodeUpdateRequest(), nil)

		require.NoError(t, err)
		assert.Equal(t, mockRes, resp)
	})

	t.Run("Nil response", func(t *testing.T) {
		nilAdapter := mocking.NewMockRequestAdapter()
		nilAdapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
		nilBuilder := NewNodeItemRequestBuilderInternal(map[string]string{
			"baseurl":     "https://example.service-now.com",
			"node_sys_id": "node123",
		}, nilAdapter)
		nilAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		resp, err := nilBuilder.Put(context.Background(), NewNodeUpdateRequest(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilResponse)
		assert.Nil(t, resp)
	})
}

func TestNodeItemRequestBuilder_Delete(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewNodeItemRequestBuilderInternal(map[string]string{
		"baseurl":     "https://example.service-now.com",
		"node_sys_id": "node123",
	}, adapter)

	// Delete goes through SendNoContent: the endpoint answers with no body, so the method
	// reports only an error.
	adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	require.NoError(t, builder.Delete(context.Background(), nil))

	adapter.AssertExpectations(t)
}

func TestValidationRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewValidationRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	t.Run("URI Construction", func(t *testing.T) {
		cdmID := "cdm456"
		config := &ValidationRequestBuilderGetRequestConfiguration{
			QueryParameters: &ValidationRequestBuilderGetQueryParameters{
				CdmID: &cdmID,
			},
		}
		requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
		requestInfo.AddQueryParameters(*config.QueryParameters)

		uri, _ := requestInfo.GetUri()
		assert.Equal(t, "https://example.service-now.com/api/sn_cdm/editor/v1/validation?cdm_id=cdm456", uri.String())
	})

	t.Run("Execution", func(t *testing.T) {
		mockRes := core.NewBaseServiceNowItemResponse[*ValidationResultModel](CreateValidationResultFromDiscriminatorValue)
		adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

		resp, err := builder.Get(context.Background(), nil)

		require.NoError(t, err)
		assert.Equal(t, mockRes, resp)
	})

	t.Run("Nil response", func(t *testing.T) {
		nilAdapter := mocking.NewMockRequestAdapter()
		nilBuilder := NewValidationRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, nilAdapter)
		nilAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		resp, err := nilBuilder.Get(context.Background(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilResponse)
		assert.Nil(t, resp)
	})
}

func TestNodesRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*NodesRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)

			postResp, err := builder.Post(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, postResp)
		})
	}
}

func TestNodeItemRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*NodeItemRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			putResp, err := builder.Put(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, putResp)

			require.ErrorIs(t, builder.Delete(context.Background(), nil), snerrors.ErrNilRequestBuilder)
		})
	}
}

func TestValidationRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*ValidationRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestNodesRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewNodesRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)

	postResp, err := builder.Post(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, postResp)
}

func TestNodeItemRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewNodeItemRequestBuilderInternal(map[string]string{}, nil)

	putResp, err := builder.Put(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, putResp)

	require.ErrorIs(t, builder.Delete(context.Background(), nil), snerrors.ErrNilRequestAdapter)
}

func TestValidationRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewValidationRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

// isDefaultErrorMapping reports whether the given argument is a non-nil
// abstractions.ErrorMappings with core.DefaultErrorMapping()'s status-code keys.
func isDefaultErrorMapping(v any) bool {
	mapping, ok := v.(abstractions.ErrorMappings)
	if !ok || mapping == nil {
		return false
	}
	for _, code := range []string{"400", "401", "403", "404", "429", "5XX", "XXX"} {
		if _, ok := mapping[code]; !ok {
			return false
		}
	}
	return len(mapping) == 7
}

// TestNodesRequestBuilder_Get_PassesDefaultErrorMapping guards against #565:
// CDM builders previously passed literal nil instead of core.DefaultErrorMapping(),
// so ServiceNow API errors never mapped to a typed core.ServiceNowError.
func TestNodesRequestBuilder_Get_PassesDefaultErrorMapping(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	mockRes := core.NewBaseServiceNowCollectionResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.MatchedBy(isDefaultErrorMapping)).Return(mockRes, nil)

	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	_, err := builder.Get(context.Background(), nil)

	require.NoError(t, err)
	adapter.AssertExpectations(t)
}
