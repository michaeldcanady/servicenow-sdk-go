package cdmapplicationsapi

// DeployablesRequestBuilderPutQueryParameters defines query parameters for updating a CDM deployable.
type DeployablesRequestBuilderPutQueryParameters struct {
	// TODO: required
	// AppName the name of the CDM application to which the deployable is associated.
	AppName *string `uriparametername:"appName"`
	// Name the name of the deployable to delete.
	Name *string `uriparametername:"name"`
	// NameDescription a description for the CDM Deployable.
	NewDescription *string `uriparametername:"newDescription"`
	// NewServiceID the ID of the desired Application Service/Dynamic CI Group. In the cURL request, provide '’ to disconnect the Deployable from the service.
	NewServiceID *string `uriparametername:"newServiceId"`
	// NewIdentifier the identifier of the Deployable.
	NewIdentifier *string `uriparametername:"newIdentifier"`
	// NewName the new name of the Deployable.
	NewName *string `uriparametername:"newName"`
	// ReturnFields list of fields to be returned as part of the response.
	ReturnFields []*string `uriparametername:"returnFields"`
}
