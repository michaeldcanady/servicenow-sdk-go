// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package errors // nolint: revive // renaming this public package would be a breaking API change; see CLAUDE.md's error-sentinel layout

import "errors"

// Standard sentinel errors for consistent error handling.
var (
	ErrNilRequestAdapter         = errors.New("requestAdapter cannot be nil")
	ErrNilRequestBuilder         = errors.New("request builder is nil")
	ErrNilResponse               = errors.New("response cannot be nil")
	ErrNilContext                = errors.New("context cannot be nil")
	ErrNilConfig                 = errors.New("config is nil")
	ErrNilBody                   = errors.New("body is nil")
	ErrNilInput                  = errors.New("input is nil")
	ErrNilRequestConfiguration   = errors.New("requestConfiguration is nil")
	ErrNilQueryParameters        = errors.New("requestConfiguration.QueryParameters is nil")
	ErrNilFactory                = errors.New("factory is nil")
	ErrNilStore                  = errors.New("store is nil")
	ErrNilPathParameters         = errors.New("pathParameters is nil")
	ErrEmptyPathParameters       = errors.New("pathParameters is empty")
	ErrNilMutator                = errors.New("mutator is nil")
	ErrNilModel                  = errors.New("model is nil")
	ErrEmptyMiddleware           = errors.New("middleware is empty")
	ErrEmptyKey                  = errors.New("key is empty")
	ErrNilRequestInfo            = errors.New("requestInfo cannot be nil")
	ErrNilClient                 = errors.New("client cannot be nil")
	ErrNilResult                 = errors.New("result property missing in response object")
	ErrWrongResponseType         = errors.New("incorrect Response Type")
	ErrParsing                   = errors.New("parsing nextLink url failed")
	ErrEmptyURI                  = errors.New("URI is empty")
	ErrNilCallback               = errors.New("callback cannot be nil")
	ErrNilParams                 = errors.New("params cannot be nil")
	ErrNilWriter                 = errors.New("serialization writer is nil")
	ErrNilProvider               = errors.New("provider is nil")
	ErrNilAuthenticationProvider = errors.New("authenticationProvider is nil")
	ErrMissingAuthentication     = errors.New("must provide either an AuthenticationProvider or a RequestAdapter")
	ErrEmptyInstance             = errors.New("instance is empty")
	ErrInvalidURL                = errors.New("invalid URL")
	ErrInsecureURL               = errors.New("URL must use HTTPS")
	ErrEmptyClientID             = errors.New("client ID is empty")
	ErrEmptyClientSecret         = errors.New("client secret is empty")
	ErrEmptyUsername             = errors.New("username is empty")
	ErrEmptyPassword             = errors.New("password is empty")
	ErrEmptyToken                = errors.New("token is empty")
	ErrEmptyAPIKey               = errors.New("API key is empty")
	ErrInvalidAuthLocation       = errors.New("invalid authentication location")
	ErrInvalidAuthMethod         = errors.New("invalid client authentication method")
	ErrNotInitialized            = errors.New("provider is not initialized")
	ErrAlreadyInitialized        = errors.New("provider is already initialized")
	ErrTokenAcquisition          = errors.New("token acquisition failed")
	ErrReauthRequired            = errors.New("reauthentication required")
	ErrUnsupportedTokenKind      = errors.New("unsupported token kind")
	ErrInvalidAssertion          = errors.New("invalid JWT assertion")
	ErrMissingTokenExpiry        = errors.New("token has no parseable expiry")
	ErrMissingAccessToken        = errors.New("token response does not include an access_token")
	ErrEmptySecret               = errors.New("secret is empty")
	ErrInvalidHMACAlgorithm      = errors.New("HMAC algorithm is invalid")
	// ErrUnknownEnumValue reports a wire value that does not correspond to any member of
	// an enum. Parse<Enum> functions wrap this with the enum name and the offending value,
	// so callers can match the class of failure with errors.Is regardless of which enum
	// produced it.
	ErrUnknownEnumValue = errors.New("unknown enum value")
	// ErrNotAuthenticated reports that ServiceNow rejected the credentials with a
	// "User Not Authenticated" response. core.UnauthorizedError matches it through
	// errors.Is, so callers can distinguish ServiceNow rejection from an identity
	// provider that failed to issue a token (credentials.ErrTokenAcquisition)
	// without asserting on the concrete error type.
	ErrNotAuthenticated = errors.New("user not authenticated")
)

// NewValidationError creates a standardized validation error message.
func NewValidationError(parameter string) error {
	return errors.New(parameter + " cannot be nil")
}
