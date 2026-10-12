// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// BasicAuthenticationProvider is an authentication.AuthenticationProvider implementing HTTP
// Basic authentication with a static username/password pair.
type BasicAuthenticationProvider struct {
	username string
	password string
}

// NewBasicProvider instantiates a new BasicAuthenticationProvider.
func NewBasicProvider(username, password string) *BasicAuthenticationProvider {
	return &BasicAuthenticationProvider{
		username: username,
		password: password,
	}
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (b *BasicAuthenticationProvider) AuthenticateRequest(context context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if conversion.IsNil(b) {
		return errors.New("provider is nil")
	}

	if err := context.Err(); err != nil {
		return err
	}

	if strings.TrimSpace(b.username) == "" {
		return snerrors.ErrEmptyUsername
	}

	if strings.TrimSpace(b.password) == "" {
		return snerrors.ErrEmptyPassword
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", b.username, b.password)))

	request.Headers.Add("Authorization", fmt.Sprintf("Basic %s", auth))

	return nil
}
