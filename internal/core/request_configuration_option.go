package core

// RequestConfigurationOption the functional option type for configuring requests.
type RequestConfigurationOption func(*RequestConfiguration)

// WithHeader sets the header for the request.
func WithHeader[T any](header T) RequestConfigurationOption {
	return func(config *RequestConfiguration) {
		config.Header = header
	}
}

// WithQueryParameters sets the query parameters for the request.
func WithQueryParameters[T any](queryParams T) RequestConfigurationOption {
	return func(config *RequestConfiguration) {
		config.QueryParameters = queryParams
	}
}

// WithData sets the data for the request.
func WithData(data interface{}) RequestConfigurationOption {
	return func(config *RequestConfiguration) {
		config.Data = data
	}
}

// WithErrorMapping sets the error mapping for the request.
func WithErrorMapping(errorMapping ErrorMapping) RequestConfigurationOption {
	return func(config *RequestConfiguration) {
		config.ErrorMapping = errorMapping
	}
}

// WithResponse sets the response handler for the request.
func WithResponse(response Response) RequestConfigurationOption {
	return func(config *RequestConfiguration) {
		config.Response = response
	}
}

// ApplyOptions applies the given options to a RequestConfiguration.
func ApplyOptions(opts ...RequestConfigurationOption) *RequestConfiguration {
	config := &RequestConfiguration{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}
