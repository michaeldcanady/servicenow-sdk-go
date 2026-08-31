// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

// FindServiceQueryParameters represents the query parameters for a find_service request.
type FindServiceQueryParameters struct {
	// Name of the application service (required if number is not provided)
	Name *string `uriparametername:"name"`
	// Number of the application service (required if name is not provided)
	Number *string `uriparametername:"number"`
}
