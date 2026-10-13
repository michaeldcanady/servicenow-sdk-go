// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

// CredentialError represents an error related to credentials or authentication.
//
// Deprecated: return the shared sentinels from the errors package instead.
type CredentialError struct {
	// Message is the error message associated with the credential error.
	Message string
}

// Oauth2Error represents an error related to OAuth2 authentication and extends CredentialError.
//
// Deprecated: return the shared sentinels from the errors package instead.
type Oauth2Error struct {
	CredentialError
}

// NewOauth2Error creates a new Oauth2Error with the specified error message.
//
// Deprecated: return the shared sentinels from the errors package instead.
func NewOauth2Error(message string) *Oauth2Error {
	return &Oauth2Error{
		CredentialError: *NewCredentialError(message),
	}
}

// NewCredentialError creates a new CredentialError with the specified error message.
//
// Deprecated: return the shared sentinels from the errors package instead.
func NewCredentialError(message string) *CredentialError {
	return &CredentialError{Message: message}
}

// Error returns the error message associated with the CredentialError.
func (e *CredentialError) Error() string {
	if e == nil {
		return "nil credential error"
	}
	return e.Message
}

// Sentinel Oauth2Error values for missing required credential fields.
//
// Deprecated: use the matching sentinel in the errors package instead. The
// values are identical, so errors.Is works across both names.
var (
	ErrEmptyClientID      = snerrors.ErrEmptyClientID
	ErrEmptyClientSecret  = snerrors.ErrEmptyClientSecret
	ErrEmptyBaseURL       = snerrors.ErrEmptyBaseURL
	ErrEmptyUsername      = snerrors.ErrEmptyUsername
	ErrEmptyPassword      = snerrors.ErrEmptyPassword
	ErrEmptyTokenProvider = snerrors.ErrEmptyTokenProvider
)
