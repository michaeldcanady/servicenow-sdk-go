package cdmapplicationsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
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
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
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
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
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

func TestApplicationsRequestBuilder_SharedLibraries(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	sharedLibrariesBuilder := builder.SharedLibraries()
	assert.NotNil(t, sharedLibrariesBuilder)
	assert.Equal(t, "{+baseurl}/api/sn_cdm/applications/shared_libraries", sharedLibrariesBuilder.GetURLTemplate())
}

func TestApplicationsRequestBuilder_Uploads(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	uploadsBuilder := builder.Uploads()
	assert.NotNil(t, uploadsBuilder)
	assert.Equal(t, "{+baseurl}/api/sn_cdm/applications/uploads", uploadsBuilder.GetURLTemplate())
}

func TestDeployablesRequestBuilder_Put(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables", uri.String())
}

func TestSharedComponentsRequestBuilder_Put(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/shared_components", uri.String())
}

func TestExportsRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).Exports()

	appName := "test_app"
	deployableName := "test_dep"
	config := &ExportsRequestBuilderGetRequestConfiguration{
		QueryParameters: &ExportsRequestBuilderGetQueryParameters{
			AppName:        &appName,
			DeployableName: &deployableName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables/exports?appName=test_app&deployableName=test_dep", uri.String())
}

func TestExportItemRequestBuilder(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).Exports().ByID("exp-123")

	statusBuilder := builder.Status()
	assert.NotNil(t, statusBuilder)
	requestInfoStatus := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, statusBuilder.GetURLTemplate(), statusBuilder.GetPathParameters())
	uriStatus, err := requestInfoStatus.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables/exports/exp-123/status", uriStatus.String())

	contentBuilder := builder.Content()
	assert.NotNil(t, contentBuilder)
	requestInfoContent := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, contentBuilder.GetURLTemplate(), contentBuilder.GetPathParameters())
	uriContent, err := requestInfoContent.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/deployables/exports/exp-123/content", uriContent.String())
}

func TestSharedLibrariesComponentsApplicationsRequestBuilder_Get(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		SharedLibraries().Components().Applications()

	appName := "test_app"
	sharedCompName := "shared_comp"
	name := "name"
	config := &SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration{
		QueryParameters: &SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters{
			AppName:             &appName,
			SharedComponentName: &sharedCompName,
			Name:                &name,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/shared_libraries/components/applications?appName=test_app&sharedComponentName=shared_comp&name=name", uri.String())
}

func TestUploadsComponentsRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Components()

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/components", uri.String())
}

func TestUploadsComponentsVarsRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Components().Vars()

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/components/vars", uri.String())
}

func TestUploadsCollectionsRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Collections()

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/collections", uri.String())
}

func TestUploadsCollectionsFileRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Collections().File()

	appName := "test_app"
	collectionName := "test_coll"
	config := &UploadsCollectionsFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &UploadsCollectionsFileRequestBuilderPostQueryParameters{
			AppName:        &appName,
			CollectionName: &collectionName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/collections/file?appName=test_app&collectionName=test_coll", uri.String())
}

func TestUploadsDeployablesFileRequestBuilder_Post(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	builder := NewApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter).
		Uploads().Deployables().File()

	appName := "test_app"
	deployableName := "test_dep"
	config := &UploadsDeployablesFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &UploadsDeployablesFileRequestBuilderPostQueryParameters{
			AppName:        &appName,
			DeployableName: &deployableName,
		},
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	kiotaRequestInfo.AddQueryParameters(config.QueryParameters)

	uri, err := requestInfo.GetUri()
	assert.Nil(t, err)
	assert.Equal(t, "https://example.service-now.com/api/sn_cdm/applications/uploads/deployables/file?appName=test_app&deployableName=test_dep", uri.String())
}
