package caseapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCaseRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	t.Run("Search Cases", func(t *testing.T) {
		requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, caseURLTemplate, requestInfo.UrlTemplate)
	})

	t.Run("Get Case By ID", func(t *testing.T) {
		itemBuilder := builder.ByID("test-id")
		requestInfo, err := itemBuilder.ToGetRequestInformation(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, caseItemURLTemplate, requestInfo.UrlTemplate)
		assert.Equal(t, "test-id", requestInfo.PathParameters["id"])
	})
}

func TestCaseRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowCollectionResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestCaseRequestBuilder_Post(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewCaseResult(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestCaseItemRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestCaseItemRequestBuilder_Put(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*CaseResultModel](CreateCaseResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Put(context.Background(), NewCaseResult(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestCaseActivitiesRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseActivitiesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*ActivitiesResultModel](CreateActivitiesResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestCaseFieldValuesRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseFieldValuesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "field_name": "state"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*FieldValuesResultModel](CreateFieldValuesResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestCaseRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	assert.NotNil(t, builder.ByID("test-id"))
	assert.NotNil(t, builder.FieldValues("state"))
}

func TestCaseItemRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCaseItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "id": "test-id"}, adapter)

	assert.NotNil(t, builder.Activities())
	assert.NotNil(t, builder.FieldValues("state"))
}

func TestCustomerServiceRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCustomerServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	assert.NotNil(t, builder.Case())
}
