package appserviceapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAppServiceRequestBuilderInternal(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	require.NotNil(t, builder)
	assert.Equal(t, appServiceURLTemplate, builder.GetURLTemplate())
}

func TestAppServiceRequestBuilder_Create(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	createBuilder := builder.Create()
	assert.NotNil(t, createBuilder)
	assert.Equal(t, createURLTemplate, createBuilder.GetURLTemplate())

	// Test URL and body serialization
	req := NewCreateServiceRequest()
	name := "AppService-CreateTest"
	comments := "Testing creation endpoint"
	_ = req.setName(&name)
	_ = req.setComments(&comments)

	requestInfo, err := createBuilder.ToPostRequestInformation(context.Background(), req, nil)
	require.NoError(t, err)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/app_service/create", uri.String())
}

func TestAppServiceRequestBuilder_Csdm(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	csdmBuilder := builder.Csdm()
	assert.NotNil(t, csdmBuilder)
	assert.Equal(t, csdmAppServiceURLTemplate, csdmBuilder.GetURLTemplate())
}
