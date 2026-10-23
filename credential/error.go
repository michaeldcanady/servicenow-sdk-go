// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"fmt"
)

// OAuthError describes a token endpoint or token lifecycle failure. It never
// contains a token, secret, password, API key, authorization code, PKCE
// verifier, or complete authorization URL.
type OAuthError struct {
	// Operation is the lifecycle step that failed, such as "acquire" or
	// "refresh".
	Operation string
	// Flow identifies the OAuth flow.
	Flow Flow
	// Code is the OAuth error code returned by the authority.
	Code string
	// Description is the OAuth error description returned by the authority.
	Description string
	// URI is the RFC 6749 error_uri returned by the authority.
	URI string
	// StatusCode is the HTTP status returned by the authority.
	StatusCode int
	// Err is the underlying sentinel or cause.
	Err error
}

// Error returns a redacted description of the failure.
func (e *OAuthError) Error() string {
	if e == nil {
		return "nil OAuth error"
	}

	message := fmt.Sprintf("oauth2 %s failed for flow %s", e.Operation, e.Flow)
	if e.Code != "" {
		message += ": " + e.Code
	}
	if e.StatusCode != 0 {
		message += fmt.Sprintf(" (status code: %d)", e.StatusCode)
	}
	return message
}

// Unwrap returns the underlying error so errors.Is can reach the sentinel.
func (e *OAuthError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}
