package tableapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRequestBuilder2Internal(t *testing.T) {
	pathParams := map[string]string{"test": "value"}
	requestAdapter := &mocking.RequestAdapter{}
	factory := mocking.ParsableFactory

	builder := newRequestBuilder2Internal(pathParams, requestAdapter, factory)
	assert.NotNil(t, builder)
	assert.Equal(t, pathParams, builder.BaseRequestBuilder.PathParameters)
	assert.Equal(t, requestAdapter, builder.BaseRequestBuilder.RequestAdapter)
}

func TestNewRequestBuilderBuilder2(t *testing.T) {
	rawURL := "https://example.com"
	requestAdapter := &mocking.RequestAdapter{}
	factory := mocking.ParsableFactory

	builder := newRequestBuilderBuilder2(rawURL, requestAdapter, factory)
	assert.NotNil(t, builder)
	assert.Equal(t, rawURL, builder.BaseRequestBuilder.PathParameters["request-raw-url"])
	assert.Equal(t, requestAdapter, builder.BaseRequestBuilder.RequestAdapter)
}

func TestTableRequestBuilder2_ByID(t *testing.T) {
	builder := getTestBuilder()
	sysID := "1234"
	itemBuilder := builder.ByID(sysID)
	assert.NotNil(t, itemBuilder)
	assert.Equal(t, sysID, itemBuilder.BaseRequestBuilder.PathParameters["sysid"])
}

func TestTableRequestBuilder2_Get(t *testing.T) {
	requestConfiguration := &TableRequestBuilder2GetRequestConfiguration{}
	ctx := context.Background()

	builder := getTestBuilder()
	builder.RequestAdapter.(*mocking.RequestAdapter).On("Send", context.Background(), mock.AnythingOfType("*abstractions.RequestInformation"), mock.AnythingOfType("serialization.ParsableFactory"), abstractions.ErrorMappings{}).Return(new(serviceNowCollectionResponse), nil)

	_, err := builder.Get(ctx, requestConfiguration)
	assert.Nil(t, err)
}

func TestTableRequestBuilder2_Post(t *testing.T) {
	builder := getTestBuilder()
	body := &tableRecord{}
	requestConfiguration := &TableRequestBuilder2PostRequestConfiguration{}
	ctx := context.Background()

	_, err := builder.Post(ctx, body, requestConfiguration)
	assert.Nil(t, err)
}

func getTestBuilder() *TableRequestBuilder2 {
	pathParams := map[string]string{"test": "value"}
	requestAdapter := &mocking.RequestAdapter{}
	factory := mocking.ParsableFactory
	return newRequestBuilder2Internal(pathParams, requestAdapter, factory)
}
