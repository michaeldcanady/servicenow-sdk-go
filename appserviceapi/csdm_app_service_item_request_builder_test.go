package appserviceapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCsdmAppServiceItemRequestBuilderInternal(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewCsdmAppServiceItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com", "sys_id": "service123"}, adapter)

	require.NotNil(t, builder)
	assert.Equal(t, csdmAppServiceItemURLTemplate, builder.GetURLTemplate())
}

func TestCsdmAppServiceItemRequestBuilder_PopulateService(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	builder := NewCsdmAppServiceItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com", "sys_id": "service123"}, adapter)

	populateBuilder := builder.PopulateService()
	assert.NotNil(t, populateBuilder)
	assert.Equal(t, populateServiceURLTemplate, populateBuilder.GetURLTemplate())

	requestInfo, err := populateBuilder.ToPutRequestInformation(context.Background(), NewPopulateServiceRequest(), nil)
	require.NoError(t, err)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/cmdb/csdm/app_service/service123/populate_service", uri.String())
}

func TestCsdmAppServiceItemRequestBuilder_ServiceDetails(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	builder := NewCsdmAppServiceItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com", "sys_id": "service123"}, adapter)

	detailsBuilder := builder.ServiceDetails()
	assert.NotNil(t, detailsBuilder)
	assert.Equal(t, serviceDetailsURLTemplate, detailsBuilder.GetURLTemplate())

	req := NewServiceDetailsRequest()
	details := NewBasicDetails()
	env := "Production"
	name := "Service Name Updated"
	_ = details.SetEnvironment(&env)
	_ = details.SetName(&name)
	_ = req.SetBasicDetails(details)

	requestInfo, err := detailsBuilder.ToPutRequestInformation(context.Background(), req, nil)
	require.NoError(t, err)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/cmdb/csdm/app_service/service123/service_details", uri.String())
}
