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

// Request Configurations
type DeployablesRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[DeployablesRequestBuilderDeleteQueryParameters]
type SharedComponentsRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[SharedComponentsRequestBuilderDeleteQueryParameters]
type UploadStatusItemRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
