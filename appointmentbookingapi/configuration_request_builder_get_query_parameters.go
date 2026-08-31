// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

// ConfigurationRequestBuilderGetQueryParameters represents the query parameters for GET /configuration.
type ConfigurationRequestBuilderGetQueryParameters struct {
	// TODO: required
	// CatalogID Sys_id of the record producer.
	CatalogID *string `uriparametername:"catalog_id"`
}
