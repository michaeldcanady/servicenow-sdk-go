package cdmapplicationsapi

type SharedComponentsRequestBuilderPutQueryParameters struct {
	// ChangesetNumber unique identifier of the associated changeset.
	ChangesetNumber *string `uriparametername:"changesetNumber"`
	// SharedComponentName name of the shared component.
	SharedComponentName *string `uriparametername:"sharedComponentName"`
	// SharedLibraryName name of the shared library.
	SharedLibraryName *string `uriparametername:"sharedLibraryName"`
	// Version version name associated with the shared component.
	Version *string `uriparametername:"version"`
	// ReturnFields list of fields to return as part of the response.
	ReturnFields []*string `uriparametername:"returnFields"`
}
