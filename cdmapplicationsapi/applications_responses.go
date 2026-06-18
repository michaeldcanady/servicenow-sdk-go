package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadStatusResponse represents a response for the upload status.
type UploadStatusResponse interface {
	core.ServiceNowItemResponse[*UploadStatusResult]
}

// CreateUploadStatusResponseFromDiscriminatorValue instantiates a new UploadStatusResponse.
func CreateUploadStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil
}

// ExportsResponse represents a response containing a collection of export results.
type ExportsResponse interface {
	core.ServiceNowCollectionResponse[*ExportResult]
}

// CreateExportsResponseFromDiscriminatorValue instantiates a new ExportsResponse.
func CreateExportsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ExportResult](CreateExportResultFromDiscriminatorValue), nil
}

// ExportStatusResponse represents a response containing an export status result.
type ExportStatusResponse interface {
	core.ServiceNowItemResponse[*ExportStatusResult]
}

// CreateExportStatusResponseFromDiscriminatorValue instantiates a new ExportStatusResponse.
func CreateExportStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*ExportStatusResult](CreateExportStatusResultFromDiscriminatorValue), nil
}

// SharedLibrariesComponentsApplicationsResponse represents a response containing a collection of applications associated with shared libraries.
type SharedLibrariesComponentsApplicationsResponse interface {
	core.ServiceNowCollectionResponse[*SharedLibraryComponentApplication]
}

// CreateSharedLibrariesComponentsApplicationsResponseFromDiscriminatorValue instantiates a new SharedLibrariesComponentsApplicationsResponse.
func CreateSharedLibrariesComponentsApplicationsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*SharedLibraryComponentApplication](CreateSharedLibraryComponentApplicationFromDiscriminatorValue), nil
}

// DeployableUpdateResponse represents a response for deployable updates.
type DeployableUpdateResponse interface {
	core.ServiceNowItemResponse[*UploadStatusResult]
}

// CreateDeployableUpdateResponseFromDiscriminatorValue instantiates a new DeployableUpdateResponse.
func CreateDeployableUpdateResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil
}

// SharedComponentUpdateResponse represents a response for shared component updates.
type SharedComponentUpdateResponse interface {
	core.ServiceNowItemResponse[*UploadStatusResult]
}

// CreateSharedComponentUpdateResponseFromDiscriminatorValue instantiates a new SharedComponentUpdateResponse.
func CreateSharedComponentUpdateResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*UploadStatusResult](CreateUploadStatusResultFromDiscriminatorValue), nil
}
