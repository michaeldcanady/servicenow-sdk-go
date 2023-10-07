package credentials

import (
	"fmt"
	"net/http"
)

// HTTPServer represents an HTTP server for OAuth2 redirection.
type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(address string) *HTTPServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth_redirect.do", oauthRedirectHandler)

	httpServer := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return &HTTPServer{server: httpServer}
}

func (s *HTTPServer) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()
}

func (s *HTTPServer) Stop() {
	if err := s.server.Shutdown(nil); err != nil {
		fmt.Println("Error shutting down HTTP server:", err)
	}
}

func oauthRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the OAuth2 redirect request here.
	// You can extract the OAuth2 token from the request and store it in your TokenCredential instance.
	// Make sure to customize this handler according to your ServiceNow OAuth2 implementation.
	// After obtaining the token, you can store it in TokenCredential.Token.
	// Example:
	// token := r.FormValue("access_token")
	// tokenCredential.Token = token

	// Respond with a success message or a web page as needed.
	fmt.Fprint(w, "OAuth2 token obtained successfully!")
}
