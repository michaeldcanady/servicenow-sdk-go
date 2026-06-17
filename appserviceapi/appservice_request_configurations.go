package appserviceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CreateRequestConfiguration represents the configuration for a Create request.
type CreateRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// FindServiceQueryParameters represents the query parameters for a find_service request.
type FindServiceQueryParameters struct {
	// Name of the application service (required if number is not provided)
	Name *string `url:"name,omitempty"`
	// Number of the application service (required if name is not provided)
	Number *string `url:"number,omitempty"`
}

// FindServiceRequestConfiguration represents the configuration for a find_service request.
type FindServiceRequestConfiguration = abstractions.RequestConfiguration[FindServiceQueryParameters]

// RegisterServiceRequestConfiguration represents the configuration for a register_service request.
type RegisterServiceRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// PopulateServiceRequestConfiguration represents the configuration for a populate_service request.
type PopulateServiceRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// ServiceDetailsRequestConfiguration represents the configuration for a service_details request.
type ServiceDetailsRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
