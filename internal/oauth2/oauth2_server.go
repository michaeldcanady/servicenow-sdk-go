package oauth2

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

const (
	defaultMinPort = 8000
	defaultMaxPort = 9000
)

var localhostNames = []string{"localhost", "127.0.0.1"}

type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type oauthServer struct {
	http.Server
}

func (s *oauthServer) GetURL() string {
	return s.Addr
}

type ServerOption func(*serverConfig)

type serverConfig struct {
	hostname string
	minPort  int
	maxPort  int
}

func NewOauthServer(address string, handler http.Handler, opts ...ServerOption) (Server, error) {
	if handler == nil {
		return nil, fmt.Errorf("handler cannot be nil")
	}

	uri, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}

	hostname := uri.Hostname()

	if !slices.Contains(localhostNames, hostname) {
		return nil, fmt.Errorf("hostname must be one of the following: %s", strings.Join(localhostNames, ", "))
	}

	cfg := &serverConfig{
		minPort: defaultMinPort,
		maxPort: defaultMaxPort,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if uri.Port() == "" {
		port, err := findOpenPort(hostname, cfg.minPort, cfg.maxPort)
		if err != nil {
			return nil, fmt.Errorf("failed to find open port: %w", err)
		}
		uri.Host = fmt.Sprintf("%s:%d", hostname, port)
		address = uri.String()
	}

	return &oauthServer{
		http.Server{
			Addr:    uri.Host,
			Handler: handler,
		},
	}, nil
}

func findOpenPort(hostname string, minPort, maxPort int) (int, error) {
	for port := minPort; port <= maxPort; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
		if err == nil {
			ln.Close()
			return port, nil
		}
	}
	return 0, fmt.Errorf("no open port found between %d and %d", minPort, maxPort)
}
