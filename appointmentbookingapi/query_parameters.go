package appointmentbookingapi

// CalendarRequestBuilderGetQueryParameters represents the query parameters for GET /calendar.
type CalendarRequestBuilderGetQueryParameters struct {
	// CatalogID Sys_id of the record producer.
	CatalogID *string `url:"catalog_id,omitempty"`
	// Location Sys_id of the location.
	Location *string `url:"location,omitempty"`
	// OpenedFor Sys_id of the user.
	OpenedFor *string `url:"opened_for,omitempty"`
}

// ConfigurationRequestBuilderGetQueryParameters represents the query parameters for GET /configuration.
type ConfigurationRequestBuilderGetQueryParameters struct {
	// CatalogID Sys_id of the record producer.
	CatalogID *string `url:"catalog_id,omitempty"`
}
