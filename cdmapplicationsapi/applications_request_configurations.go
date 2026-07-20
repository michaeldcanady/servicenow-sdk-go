package cdmapplicationsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// DeployablesRequestBuilderDeleteQueryParameters represents query parameters for DELETE /applications/deployables.
type DeployablesRequestBuilderDeleteQueryParameters struct {
	AppName *string `uriparametername:"appName"`
	Name    *string `uriparametername:"name"`
}

// SharedComponentsRequestBuilderDeleteQueryParameters represents query parameters for DELETE /applications/shared_components.
type SharedComponentsRequestBuilderDeleteQueryParameters struct {
	AppName             *string `uriparametername:"appName"`
	SharedComponentName *string `uriparametername:"sharedComponentName"`
	Name                *string `uriparametername:"name"` // Support name too just in case
}

type ExportsRequestBuilderGetQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	DeployableName *string `uriparametername:"deployableName"`
}

type SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters struct {
	AppName             *string `uriparametername:"appName"`
	SharedComponentName *string `uriparametername:"sharedComponentName"`
	Name                *string `uriparametername:"name"`
}

type UploadsCollectionsFileRequestBuilderPostQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	CollectionName *string `uriparametername:"collectionName"`
}

type UploadsDeployablesFileRequestBuilderPostQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	DeployableName *string `uriparametername:"deployableName"`
}

// Request Configurations
type DeployablesRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[DeployablesRequestBuilderDeleteQueryParameters]
type SharedComponentsRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[SharedComponentsRequestBuilderDeleteQueryParameters]
type UploadStatusItemRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

type ExportsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ExportsRequestBuilderGetQueryParameters]
type ExportItemStatusRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type ExportItemContentRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

type SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters]

type UploadsComponentsRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type UploadsComponentsVarsRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type UploadsCollectionsRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type UploadsCollectionsFileRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[UploadsCollectionsFileRequestBuilderPostQueryParameters]
type UploadsDeployablesFileRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[UploadsDeployablesFileRequestBuilderPostQueryParameters]

type DeployablesRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type SharedComponentsRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
