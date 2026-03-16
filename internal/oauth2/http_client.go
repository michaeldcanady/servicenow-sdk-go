package oauth2

import "net/http"

// HTTPClient is an interface that abstracts the standard http.Client.
type HTTPClient interface {
	// Do executes an HTTP request and returns an HTTP response.
	Do(req *http.Request) (*http.Response, error)
}
