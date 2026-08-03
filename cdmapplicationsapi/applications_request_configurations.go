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

// ExportsRequestBuilderGetQueryParameters represents the GET query parameters for the Exports resource.
type ExportsRequestBuilderGetQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	DeployableName *string `uriparametername:"deployableName"`
}

// SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters represents the GET query parameters for the Shared Libraries Components Applications resource.
type SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters struct {
	AppName             *string `uriparametername:"appName"`
	SharedComponentName *string `uriparametername:"sharedComponentName"`
	Name                *string `uriparametername:"name"`
}

// UploadsCollectionsFileRequestBuilderPostQueryParameters represents the POST query parameters for the Uploads Collections File resource.
type UploadsCollectionsFileRequestBuilderPostQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	CollectionName *string `uriparametername:"collectionName"`
}

// UploadsDeployablesFileRequestBuilderPostQueryParameters represents the POST query parameters for the Uploads Deployables File resource.
type UploadsDeployablesFileRequestBuilderPostQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	DeployableName *string `uriparametername:"deployableName"`
}

// DeployablesRequestBuilderDeleteRequestConfiguration represents the configuration for a Delete request.
type DeployablesRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[DeployablesRequestBuilderDeleteQueryParameters]

// SharedComponentsRequestBuilderDeleteRequestConfiguration represents the DELETE request configuration for the Shared Components resource.
type SharedComponentsRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[SharedComponentsRequestBuilderDeleteQueryParameters]

// UploadStatusItemRequestBuilderGetRequestConfiguration represents the GET request configuration for the Upload Status Item resource.
type UploadStatusItemRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// ExportsRequestBuilderGetRequestConfiguration represents the GET request configuration for the Exports resource.
type ExportsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ExportsRequestBuilderGetQueryParameters]

// ExportItemStatusRequestBuilderGetRequestConfiguration represents the GET request configuration for the Export Item Status resource.
type ExportItemStatusRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// ExportItemContentRequestBuilderGetRequestConfiguration represents the GET request configuration for the Export Item Content resource.
type ExportItemContentRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration represents the GET request configuration for the Shared Libraries Components Applications resource.
type SharedLibrariesComponentsApplicationsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters]

// UploadsComponentsRequestBuilderPostRequestConfiguration represents the POST request configuration for the Uploads Components resource.
type UploadsComponentsRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// UploadsComponentsVarsRequestBuilderPostRequestConfiguration represents the POST request configuration for the Uploads Components Vars resource.
type UploadsComponentsVarsRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// UploadsCollectionsRequestBuilderPostRequestConfiguration represents the POST request configuration for the Uploads Collections resource.
type UploadsCollectionsRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// UploadsCollectionsFileRequestBuilderPostRequestConfiguration represents the POST request configuration for the Uploads Collections File resource.
type UploadsCollectionsFileRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[UploadsCollectionsFileRequestBuilderPostQueryParameters]

// UploadsDeployablesFileRequestBuilderPostRequestConfiguration represents the POST request configuration for the Uploads Deployables File resource.
type UploadsDeployablesFileRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[UploadsDeployablesFileRequestBuilderPostQueryParameters]

// DeployablesRequestBuilderPutRequestConfiguration represents the PUT request configuration for the Deployables resource.
type DeployablesRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// SharedComponentsRequestBuilderPutRequestConfiguration represents the PUT request configuration for the Shared Components resource.
type SharedComponentsRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
