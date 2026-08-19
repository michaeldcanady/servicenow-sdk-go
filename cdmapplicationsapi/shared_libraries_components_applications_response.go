package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedLibrariesComponentsApplicationsResponse represents a response containing a collection of applications associated with shared libraries.
type SharedLibrariesComponentsApplicationsResponse interface {
	core.ServiceNowCollectionResponse[*SharedLibraryComponentApplication]
}

// CreateSharedLibrariesComponentsApplicationsResponseFromDiscriminatorValue instantiates a new SharedLibrariesComponentsApplicationsResponse.
func CreateSharedLibrariesComponentsApplicationsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*SharedLibraryComponentApplication](CreateSharedLibraryComponentApplicationFromDiscriminatorValue), nil
}
