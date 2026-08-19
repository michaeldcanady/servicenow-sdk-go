package cdmapplicationsapi

// ExportsRequestBuilderGetQueryParameters represents the GET query parameters for the Exports resource.
type ExportsRequestBuilderGetQueryParameters struct {
	AppName        *string `uriparametername:"appName"`
	DeployableName *string `uriparametername:"deployableName"`
}
