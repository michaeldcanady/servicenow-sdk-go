package credentials

import (
	"context"
	"fmt"
	"net/http"
)

// HTTPServer represents an HTTP server for OAuth2 redirection.
type HTTPServer struct {
	server *http.Server
}

// NewHTTPServer creates a new HTTP Server
func NewHTTPServer(address string) *HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth_redirect.do", OauthRedirectHandler)

	httpServer := &http.Server{
		Addr:              address,
		Handler:           mux,
		ReadHeaderTimeout: 1000,
	}

	return &HTTPServer{server: httpServer}
}

// Start starts the HTTP Server
func (s *HTTPServer) Start() {
	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()
}

// Stop stops the HTTP Server
func (s *HTTPServer) Stop() {
	err := s.server.Shutdown(context.TODO())
	if err != nil {
		fmt.Println("Error shutting down HTTP server:", err)
	}
}

func OauthRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the OAuth2 redirect request here.
	// You can extract the OAuth2 token from the request and store it in your TokenCredential instance.
	// Make sure to customize this handler according to your ServiceNow OAuth2 implementation.
	// After obtaining the token, you can store it in TokenCredential.Token.
	// Example:
	// token := r.FormValue("access_token")
	// tokenCredential.Token = token

	// Respond with a success message or a web page as needed.
	_, _ = fmt.Fprint(w, "OAuth2 token obtained successfully!")
}
