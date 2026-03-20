package oauth2

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

// AuthorizationResult holds the data or error captured by the callback server.
type AuthorizationResult struct {
	// Code is the authorization code sent by the server.
	Code string
	// State is the state parameter returned by the server.
	State string
	// Err is set if an error occurred during the callback.
	Err error
}

// Server is a local HTTP server used to capture authorization codes in redirect flows.
type Server struct {
	// Addr is the full URL where the server is listening.
	Addr   string
	server *http.Server
	result chan AuthorizationResult
	state  string
}

const (
	localhost  = "localhost"
	tcpNetwork = "tcp"
)

// NewServer starts a new local HTTP server on the specified port (or a random one if port is 0).
func NewServer(state string, port int) (*Server, error) {
	listener, err := net.Listen(tcpNetwork, fmt.Sprintf("%s:%d", localhost, port))
	if err != nil && port != 0 {
		// Fallback to random port if specific port is taken
		listener, err = net.Listen(tcpNetwork, fmt.Sprintf("%s:0", localhost))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to start local server: %w", err)
	}

	_, portStr, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse listener address: %w", err)
	}

	s := &Server{
		Addr:   fmt.Sprintf("http://%s:%s", localhost, portStr),
		result: make(chan AuthorizationResult, 1),
		state:  state,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleCallback)

	s.server = &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := s.server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.result <- AuthorizationResult{Err: err}
		}
	}()

	return s, nil
}

// Addr returns the full URL where the server is listening.
func (s *Server) GetAddr() string {
	return s.Addr
}

// Result waits for and returns the authorization result.
func (s *Server) Result(ctx context.Context) AuthorizationResult {
	select {
	case res := <-s.result:
		return res
	case <-ctx.Done():
		return AuthorizationResult{Err: ctx.Err()}
	}
}

// Shutdown gracefully stops the local HTTP server.
func (s *Server) Shutdown(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}

func (s *Server) handleCallback(w http.ResponseWriter, r *http.Request) {
	errParam := r.URL.Query().Get(ErrorKey)
	errDesc := r.URL.Query().Get(ErrorDescriptionKey)
	code := r.URL.Query().Get(CodeKey)
	state := r.URL.Query().Get(StateKey)

	if errParam != "" {
		errMsg := fmt.Sprintf("server returned error: %s", errParam)
		if errDesc != "" {
			errMsg += fmt.Sprintf(" (%s)", errDesc)
		}
		err := errors.New(errMsg)
		s.resultCh(AuthorizationResult{Err: err, State: state})
		s.writeResponse(w, "Authentication Failed", errMsg, true)
		return
	}

	if code == "" {
		err := errors.New("no code found in callback")
		s.resultCh(AuthorizationResult{Err: err, State: state})
		s.writeResponse(w, "Authentication Failed", "No authorization code was found in the callback URL.", true)
		return
	}

	if s.state != "" && state == "" {
		err := errors.New("state mismatch: missing state")
		s.resultCh(AuthorizationResult{Err: err, State: state})
		s.writeResponse(w, "Authentication Failed", "Security state missing. Expected state but none was found. Please try again.", true)
		return
	}

	if s.state != "" && state != s.state {
		err := fmt.Errorf("state mismatch: expected %s, got %s", s.state, state)
		s.resultCh(AuthorizationResult{Err: err, State: state})
		s.writeResponse(w, "Authentication Failed", fmt.Sprintf("Security state mismatch. Expected %s, but got %s. Please try again.", s.state, state), true)
		return
	}

	s.resultCh(AuthorizationResult{Code: code, State: state})
	s.writeResponse(w, "Authentication Successful", "You have successfully authenticated. You can now close this window.", false)
}

func (s *Server) resultCh(res AuthorizationResult) {
	select {
	case s.result <- res:
	default:
	}
}

func (s *Server) writeResponse(w http.ResponseWriter, title, message string, isError bool) {
	w.Header().Set(ContentTypeKey, "text/html; charset=utf-8")
	if isError {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	color := "#28a745" // Success green
	if isError {
		color = "#dc3545" // Error red
	}

	// nolint: gosec // G705 //text comes from exception information not user
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>%s</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif; line-height: 1.6; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; background-color: #f8f9fa; }
        .container { text-align: center; background: white; padding: 2rem; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); max-width: 400px; }
        h1 { color: %s; margin-top: 0; }
        p { color: #6c757d; }
    </style>
</head>
<body>
    <div class="container">
        <h1>%s</h1>
        <p>%s</p>
    </div>
</body>
</html>
`, title, color, title, message)
}
