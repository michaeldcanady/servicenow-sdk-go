package appointmentbookingapi

// CalendarRequestBuilderGetQueryParameters represents the query parameters for GET /calendar.
type CalendarRequestBuilderGetQueryParameters struct {
	// CatalogID Sys_id of the record producer.
	CatalogID *string `uriparametername:"catalog_id"`
	// Location Sys_id of the location.
	Location *string `uriparametername:"location"`
	// OpenedFor Sys_id of the user.
	OpenedFor *string `uriparametername:"opened_for"`
}

// ConfigurationRequestBuilderGetQueryParameters represents the query parameters for GET /configuration.
type ConfigurationRequestBuilderGetQueryParameters struct {
	// CatalogID Sys_id of the record producer.
	CatalogID *string `uriparametername:"catalog_id"`
}
