package appserviceapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
)

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
	assert.Nil(t, err)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/app_service/create", uri.String())
}

func TestAppServiceRequestBuilder_Csdm_FindService(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	csdmBuilder := builder.Csdm()
	assert.NotNil(t, csdmBuilder)
	assert.Equal(t, csdmAppServiceURLTemplate, csdmBuilder.GetURLTemplate())

	findBuilder := csdmBuilder.FindService()
	assert.NotNil(t, findBuilder)
	assert.Equal(t, findServiceURLTemplate, findBuilder.GetURLTemplate())

	name := "Email_East"
	number := "SNSVC0001018"
	config := &FindServiceRequestConfiguration{
		QueryParameters: &FindServiceQueryParameters{
			Name:   &name,
			Number: &number,
		},
	}

	requestInfo, err := findBuilder.ToGetRequestInformation(context.Background(), config)
	assert.Nil(t, err)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/csdm/app_service/find_service?name=Email_East&number=SNSVC0001018", uri.String())
}

func TestAppServiceRequestBuilder_Csdm_RegisterService(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	registerBuilder := builder.Csdm().RegisterService()
	assert.NotNil(t, registerBuilder)
	assert.Equal(t, registerServiceURLTemplate, registerBuilder.GetURLTemplate())

	req := NewRegisterServiceRequest()
	details := NewBasicDetails()
	env := "Test Lab 2"
	name := "Service Name Here"
	_ = details.setEnvironment(&env)
	_ = details.setName(&name)
	_ = req.setBasicDetails(details)

	requestInfo, err := registerBuilder.ToPostRequestInformation(context.Background(), req, nil)
	assert.Nil(t, err)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/csdm/app_service/register_service", uri.String())
}

func TestAppServiceRequestBuilder_Csdm_PopulateService(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	populateBuilder := builder.Csdm().ByID("service123").PopulateService()
	assert.NotNil(t, populateBuilder)
	assert.Equal(t, populateServiceURLTemplate, populateBuilder.GetURLTemplate())

	req := NewPopulateServiceRequest()
	rel := NewServiceRelation()
	parent := "parent123"
	child := "child456"
	relType := "Depends on::Used by"
	_ = rel.setParent(&parent)
	_ = rel.setChild(&child)
	_ = rel.setType(&relType)
	_ = req.setServiceRelations([]*ServiceRelation{rel})

	requestInfo, err := populateBuilder.ToPutRequestInformation(context.Background(), req, nil)
	assert.Nil(t, err)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/csdm/app_service/service123/populate_service", uri.String())
}

func TestAppServiceRequestBuilder_Csdm_ServiceDetails(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	detailsBuilder := builder.Csdm().ByID("service123").ServiceDetails()
	assert.NotNil(t, detailsBuilder)
	assert.Equal(t, serviceDetailsURLTemplate, detailsBuilder.GetURLTemplate())

	req := NewServiceDetailsRequest()
	details := NewBasicDetails()
	env := "Production"
	name := "Service Name Updated"
	_ = details.setEnvironment(&env)
	_ = details.setName(&name)
	_ = req.setBasicDetails(details)

	requestInfo, err := detailsBuilder.ToPutRequestInformation(context.Background(), req, nil)
	assert.Nil(t, err)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/csdm/app_service/service123/service_details", uri.String())
}
