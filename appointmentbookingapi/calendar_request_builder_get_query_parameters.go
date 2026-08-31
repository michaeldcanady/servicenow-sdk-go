// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

// CalendarRequestBuilderGetQueryParameters represents the query parameters for GET /calendar.
type CalendarRequestBuilderGetQueryParameters struct {
	// TODO: required
	// CatalogID Sys_id of the record producer.
	CatalogID *string `uriparametername:"catalog_id"`
	// TODO: required
	// Location Sys_id of the location.
	Location *string `uriparametername:"location"`
	// TODO: required
	// OpenedFor Sys_id of the user.
	OpenedFor *string `uriparametername:"opened_for"`
}
