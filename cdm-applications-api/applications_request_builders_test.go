package cdmapplicationsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestApplicationsRequestBuilder_Deployables(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	deployablesBuilder := builder.Deployables()
	assert.NotNil(t, deployablesBuilder)
	assert.Equal(t, deployablesURLTemplate, deployablesBuilder.GetURLTemplate())
}

func TestApplicationsRequestBuilder_SharedComponents(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	sharedComponentsBuilder := builder.SharedComponents()
	assert.NotNil(t, sharedComponentsBuilder)
	assert.Equal(t, sharedComponentsURLTemplate, sharedComponentsBuilder.GetURLTemplate())
}

func TestApplicationsRequestBuilder_UploadStatus(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	uploadStatusBuilder := builder.UploadStatus()
	assert.NotNil(t, uploadStatusBuilder)
	assert.Equal(t, uploadStatusURLTemplate, uploadStatusBuilder.GetURLTemplate())
}

func TestDeployablesRequestBuilder_Delete(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	appName := "test_app"
	name := "Dep-1"
	config := &DeployablesRequestBuilderDeleteRequestConfiguration{
		QueryParameters: &DeployablesRequestBuilderDeleteQueryParameters{
			AppName: &appName,
			Name:    &name,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables?appName=test_app&name=Dep-1", uri.String())
}

func TestSharedComponentsRequestBuilder_Delete(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	appName := "test_app"
	sharedComponentName := "Comp-1"
	config := &SharedComponentsRequestBuilderDeleteRequestConfiguration{
		QueryParameters: &SharedComponentsRequestBuilderDeleteQueryParameters{
			AppName:             &appName,
			SharedComponentName: &sharedComponentName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/shared_components?appName=test_app&sharedComponentName=Comp-1", uri.String())
}

func TestUploadStatusItemRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewUploadStatusRequestBuilderInternal(map[string]string{
		"baseurl": "https://example.service-now.com",
	}, adapter)

	itemBuilder := builder.ByID("upload123")
	assert.NotNil(t, itemBuilder)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, itemBuilder.GetURLTemplate(), itemBuilder.GetPathParameters())

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/upload-status/upload123", uri.String())
}
