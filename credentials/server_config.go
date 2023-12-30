package credentials

// serverConfig is now exported for user customization
type serverConfig struct {
	Host           string
	PortRangeStart int
	PortRangeEnd   int
}

// validateServerConfig validates and adjusts the server configuration
func validateServerConfig(config *serverConfig) {

	if config.Host == "" {
		config.Host = defaultHost
	}

	if config.PortRangeStart == 0 {
		config.PortRangeStart = defaultPortStart
	}

	if config.PortRangeEnd == 0 {
		config.PortRangeStart = defaultPortEnd
	}
}
