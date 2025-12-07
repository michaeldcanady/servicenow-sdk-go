package oauth2

import "fmt"

type TokenError struct {
	Err              string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
	ErrorURI         string `json:"error_uri,omitempty"`
	StatusCode       int    `json:"-"`
	RawBody          string `json:"-"`
}

func (e *TokenError) Error() string {
	if e.ErrorDescription != "" {
		return fmt.Sprintf("oauth2 error: %s (%s)", e.Err, e.ErrorDescription)
	}
	return fmt.Sprintf("oauth2 error: %s", e.Err)
}
