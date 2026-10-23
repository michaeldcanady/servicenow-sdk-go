// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

func normalizeBaseURL(raw string) (string, error) {
	if strings.TrimSpace(raw) == "" {
		return "", snerrors.ErrEmptyURI
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("%w: %w", snerrors.ErrInvalidURL, err)
	}
	if !strings.EqualFold(parsed.Scheme, "https") {
		return "", snerrors.ErrInsecureURL
	}
	if parsed.Host == "" {
		return "", fmt.Errorf("%w: host is empty", snerrors.ErrInvalidURL)
	}

	return strings.ToLower(parsed.Scheme) + "://" + strings.ToLower(parsed.Host), nil
}

func normalizeScopes(scopes []string) []string {
	if len(scopes) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(scopes))
	normalized := make([]string, 0, len(scopes))
	for _, scope := range scopes {
		scope = strings.TrimSpace(scope)
		if scope == "" {
			continue
		}
		if _, ok := seen[scope]; ok {
			continue
		}
		seen[scope] = struct{}{}
		normalized = append(normalized, scope)
	}
	sort.Strings(normalized)
	return normalized
}

func newAllowedHostsValidator(config ProviderConfig) (*authentication.AllowedHostsValidator, error) {
	return authentication.NewAllowedHostsValidatorErrorCheck(config.AllowedHosts)
}

func requestURI(ctx context.Context, request *abstractions.RequestInformation, validator *authentication.AllowedHostsValidator) (*url.URL, bool, error) {
	if ctx == nil {
		return nil, false, snerrors.ErrNilContext
	}
	if request == nil {
		return nil, false, snerrors.ErrNilRequestInfo
	}
	if request.Headers == nil {
		request.Headers = abstractions.NewRequestHeaders()
	}

	uri, err := request.GetUri()
	if err != nil {
		return nil, false, err
	}
	if validator != nil && !validator.IsUrlHostValid(uri) {
		return uri, false, nil
	}
	if !strings.EqualFold(uri.Scheme, "https") {
		return uri, false, snerrors.ErrInsecureURL
	}
	return uri, true, nil
}

func setHeaderIfAbsent(request *abstractions.RequestInformation, name, value string) {
	if request == nil || request.Headers == nil || value == "" {
		return
	}
	request.Headers.TryAdd(name, value)
}

func (m ClientAuthMethod) oauth2Method() (oauth2.AuthMethod, error) {
	switch m {
	case ClientAuthMethodDefault, ClientAuthMethodClientSecretPost:
		return oauth2.AuthMethodClientSecretPost, nil
	case ClientAuthMethodClientSecretBasic:
		return oauth2.AuthMethodClientSecretBasic, nil
	case ClientAuthMethodNone:
		return oauth2.AuthMethodNone, nil
	default:
		return oauth2.AuthMethodUnknown, snerrors.ErrInvalidAuthMethod
	}
}
