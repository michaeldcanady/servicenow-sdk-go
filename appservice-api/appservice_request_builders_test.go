package appserviceapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
)

func TestAppServiceRequestBuilder_CreateOrUpdateService(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	
	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	createOrUpdateBuilder := builder.CreateOrUpdateService()
	assert.NotNil(t, createOrUpdateBuilder)
	assert.Equal(t, createOrUpdateServiceURLTemplate, createOrUpdateBuilder.GetURLTemplate())

	// Test URL and body serialization
	req := NewCreateOrUpdateServiceRequest()
	name := "AppService-Test"
	_ = req.setName(&name)

	requestInfo, err := createOrUpdateBuilder.ToPostRequestInformation(context.Background(), req, nil)
	assert.Nil(t, err)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/app_service/createOrUpdateService", uri.String())
}


func TestAppServiceRequestBuilder_ByID(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewAppServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	itemBuilder := builder.ByID("service123")
	assert.NotNil(t, itemBuilder)
	assert.Equal(t, appServiceItemURLTemplate, itemBuilder.GetURLTemplate())

	getContentBuilder := itemBuilder.GetContent()
	assert.NotNil(t, getContentBuilder)
	assert.Equal(t, getContentURLTemplate, getContentBuilder.GetURLTemplate())

	// Test GetContent URL serialization with query parameter
	mode := "Full"
	config := &GetContentRequestBuilderGetRequestConfiguration{
		QueryParameters: &GetContentRequestBuilderGetQueryParameters{
			Mode: &mode,
		},
	}

	requestInfo, err := getContentBuilder.ToGetRequestInformation(context.Background(), config)
	assert.Nil(t, err)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/v1/cmdb/app_service/service123/getContent?Mode=Full", uri.String())
}

func TestCreateOrUpdateServiceRequest_Serialization(t *testing.T) {
	req := NewCreateOrUpdateServiceRequest()
	assert.NotNil(t, req)

	name := "AppService-Test"
	typ := "manual"

	err := req.setName(&name)
	assert.Nil(t, err)
	err = req.setType(&typ)
	assert.Nil(t, err)

	retName, err := req.GetName()
	assert.Nil(t, err)
	assert.Equal(t, &name, retName)

	retType, err := req.GetType()
	assert.Nil(t, err)
	assert.Equal(t, &typ, retType)
}

func TestCreateOrUpdateServiceResult_Serialization(t *testing.T) {
	res := NewCreateOrUpdateServiceResult()
	assert.NotNil(t, res)

	sysID := "sys-12345"
	name := "MyManualService"
	className := "cmdb_ci_service_discovered"

	err := res.setSysId(&sysID)
	assert.Nil(t, err)
	err = res.setName(&name)
	assert.Nil(t, err)
	err = res.setClassName(&className)
	assert.Nil(t, err)

	retSysID, err := res.GetSysId()
	assert.Nil(t, err)
	assert.Equal(t, &sysID, retSysID)

	retName, err := res.GetName()
	assert.Nil(t, err)
	assert.Equal(t, &name, retName)

	retClassName, err := res.GetClassName()
	assert.Nil(t, err)
	assert.Equal(t, &className, retClassName)
}

func TestGetContentResult_Serialization(t *testing.T) {
	res := NewGetContentResult()
	assert.NotNil(t, res)

	ci := NewCIInfo()
	sysID := "ci-123"
	ciName := "CI-One"
	ciClass := "cmdb_ci_linux_server"
	_ = ci.setSysId(&sysID)
	_ = ci.setName(&ciName)
	_ = ci.setClassName(&ciClass)

	rel := NewRelationshipInfo()
	parent := "ci-123"
	child := "ci-456"
	relType := "Depends on::Used by"
	_ = rel.setParent(&parent)
	_ = rel.setChild(&child)
	_ = rel.setType(&relType)

	err := res.setCis([]*CIInfo{ci})
	assert.Nil(t, err)
	err = res.setRelations([]*RelationshipInfo{rel})
	assert.Nil(t, err)

	retCis, err := res.GetCis()
	assert.Nil(t, err)
	assert.Len(t, retCis, 1)
	
	retSysID, err := retCis[0].GetSysId()
	assert.Nil(t, err)
	assert.Equal(t, &sysID, retSysID)

	retRels, err := res.GetRelations()
	assert.Nil(t, err)
	assert.Len(t, retRels, 1)
	
	retParent, err := retRels[0].GetParent()
	assert.Nil(t, err)
	assert.Equal(t, &parent, retParent)
}


func TestAppServiceRequestBuilder_RawURL(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewAppServiceRequestBuilder("https://example.service-now.com/api/now/v1/cmdb/app_service", adapter)
	assert.NotNil(t, builder)
	assert.Equal(t, appServiceURLTemplate, builder.GetURLTemplate())
}

func TestGetContentResponse_Discriminator(t *testing.T) {
	p, err := CreateGetContentResponseFromDiscriminatorValue(nil)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestCreateOrUpdateServiceResponse_Discriminator(t *testing.T) {
	p, err := CreateCreateOrUpdateServiceResponseFromDiscriminatorValue(nil)
	assert.Nil(t, err)
	assert.NotNil(t, p)
}
