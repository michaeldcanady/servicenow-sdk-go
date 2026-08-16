package cdmapplicationsapi

// DeployablesRequestBuilderDeleteQueryParameters represents query parameters for DELETE /applications/deployables.
type DeployablesRequestBuilderDeleteQueryParameters struct {
	// TODO: required
	// AppName the name of the CDM application to which the deployable is associated.
	AppName *string `uriparametername:"appName"`
	// Name the name of the deployable to delete.
	Name *string `uriparametername:"name"`
}
