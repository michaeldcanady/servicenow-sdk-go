package oauth2

import "fmt"

// TokenError represents an error returned by the OAuth2 server during token exchange or management.
type TokenError struct {
	// Err is the error code returned by the server (e.g., invalid_request).
	Err string `json:"error"`
	// ErrorDescription is a human-readable string providing additional information.
	ErrorDescription string `json:"error_description,omitempty"`
	// ErrorURI is a URL identifying a human-readable web page with information about the error.
	ErrorURI string `json:"error_uri,omitempty"`
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"-"`
	// RawBody is the raw HTTP response body.
	RawBody string `json:"-"`
}

// Error returns the string representation of the TokenError.
func (e *TokenError) Error() string {
	msg := fmt.Sprintf("oauth2 error: %s", e.Err)
	if e.ErrorDescription != "" {
		msg += fmt.Sprintf(" (%s)", e.ErrorDescription)
	}
	if e.StatusCode != 0 {
		msg += fmt.Sprintf(" (status code: %d)", e.StatusCode)
	}
	return msg
}
