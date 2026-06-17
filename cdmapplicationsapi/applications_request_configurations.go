package cdmapplicationsapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// DeployablesRequestBuilderDeleteQueryParameters represents query parameters for DELETE /applications/deployables.
type DeployablesRequestBuilderDeleteQueryParameters struct {
	AppName *string `url:"appName,omitempty"`
	Name    *string `url:"name,omitempty"`
}

// SharedComponentsRequestBuilderDeleteQueryParameters represents query parameters for DELETE /applications/shared_components.
type SharedComponentsRequestBuilderDeleteQueryParameters struct {
	AppName             *string `url:"appName,omitempty"`
	SharedComponentName *string `url:"sharedComponentName,omitempty"`
	Name                *string `url:"name,omitempty"` // Support name too just in case
}

type ExportsRequestBuilderGetQueryParameters struct {
	AppName        *string `url:"appName,omitempty"`
	DeployableName *string `url:"deployableName,omitempty"`
}

type SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters struct {
	AppName             *string `url:"appName,omitempty"`
	SharedComponentName *string `url:"sharedComponentName,omitempty"`
	Name                *string `url:"name,omitempty"`
}

type UploadsCollectionsFileRequestBuilderPostQueryParameters struct {
	AppName        *string `url:"appName,omitempty"`
	CollectionName *string `url:"collectionName,omitempty"`
}

type UploadsDeployablesFileRequestBuilderPostQueryParameters struct {
	AppName        *string `url:"appName,omitempty"`
	DeployableName *string `url:"deployableName,omitempty"`
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
