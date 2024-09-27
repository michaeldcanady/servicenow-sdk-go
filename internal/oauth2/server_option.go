package oauth2

type ServerOption func(*serverConfig)

func WithMinPort(port int) ServerOption {
	return func(sc *serverConfig) {
		sc.minPort = port
	}
}

func WithMaxPort(port int) ServerOption {
	return func(sc *serverConfig) {
		sc.maxPort = port
	}
}
