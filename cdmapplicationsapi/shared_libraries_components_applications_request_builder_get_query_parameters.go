package cdmapplicationsapi

// SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters represents the GET query parameters for the Shared Libraries Components Applications resource.
type SharedLibrariesComponentsApplicationsRequestBuilderGetQueryParameters struct {
	AppName             *string `uriparametername:"appName"`
	SharedComponentName *string `uriparametername:"sharedComponentName"`
	Name                *string `uriparametername:"name"`
}
