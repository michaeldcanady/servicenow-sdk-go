package cdmeditorapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		sysId := "123"
		config := &NodesRequestBuilderGetRequestConfiguration{
			QueryParameters: &NodesRequestBuilderGetQueryParameters{
				SysId: &sysId,
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

		assert.NoError(t, err)
		assert.Equal(t, mockRes, resp)
	})
}

func TestNodesRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*NodeResultModel](CreateNodeResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewNodeCreateRequest(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
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

		assert.NoError(t, err)
		assert.Equal(t, mockRes, resp)
	})
}

func TestNodeItemRequestBuilder_Delete(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewNodeItemRequestBuilderInternal(map[string]string{
		"baseurl":     "https://example.service-now.com",
		"node_sys_id": "node123",
	}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*MessageResult](CreateMessageResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Delete(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestValidationRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewValidationRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	t.Run("URI Construction", func(t *testing.T) {
		cdmId := "cdm456"
		config := &ValidationRequestBuilderGetRequestConfiguration{
			QueryParameters: &ValidationRequestBuilderGetQueryParameters{
				CdmId: &cdmId,
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

		assert.NoError(t, err)
		assert.Equal(t, mockRes, resp)
	})
}
