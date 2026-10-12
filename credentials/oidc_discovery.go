// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
)

// oidcWellKnownPath is the standard OIDC discovery path appended to an issuer URL.
const oidcWellKnownPath = "/.well-known/openid-configuration"

// oidcDiscoveryDocument is the subset of an OIDC provider metadata document the
// SDK uses to talk to the external identity provider.
type oidcDiscoveryDocument struct {
	// Issuer is the issuer identifier the provider advertises.
	Issuer string `json:"issuer"`
	// AuthorizationEndpoint is where the user is sent to authenticate.
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	// TokenEndpoint is where the authorization code is exchanged for tokens.
	TokenEndpoint string `json:"token_endpoint"`
	// RevocationEndpoint is optional; it is empty when the provider does not
	// advertise one.
	RevocationEndpoint string `json:"revocation_endpoint"`
}

// endpoints maps the discovery document onto the internal OAuth2 endpoints.
func (d *oidcDiscoveryDocument) endpoints() *oauth2.Endpoints {
	return &oauth2.Endpoints{
		AuthURL:       d.AuthorizationEndpoint,
		TokenURL:      d.TokenEndpoint,
		RevocationURL: d.RevocationEndpoint,
	}
}

// validate ensures the document carries the endpoints the flow requires.
func (d *oidcDiscoveryDocument) validate() error {
	if strings.TrimSpace(d.AuthorizationEndpoint) == "" {
		return fmt.Errorf("OIDC discovery document is missing authorization_endpoint")
	}

	if strings.TrimSpace(d.TokenEndpoint) == "" {
		return fmt.Errorf("OIDC discovery document is missing token_endpoint")
	}

	return nil
}

// oidcDiscoveryURL resolves the metadata URL for an issuer that is either a
// plain issuer URL (the well-known path is appended) or a full metadata URL
// (used as-is).
func oidcDiscoveryURL(issuer string) (string, error) {
	trimmed := strings.TrimSpace(issuer)
	if trimmed == "" {
		return "", ErrEmptyOIDCIssuer
	}

	parsed, err := url.Parse(trimmed)
	if err != nil {
		return "", fmt.Errorf("invalid OIDC issuer %q: %w", issuer, err)
	}

	if !parsed.IsAbs() || parsed.Host == "" {
		return "", fmt.Errorf("OIDC issuer %q must be an absolute URL", issuer)
	}

	if strings.HasSuffix(strings.TrimRight(parsed.Path, "/"), oidcWellKnownPath) {
		return parsed.String(), nil
	}

	parsed.Path = strings.TrimRight(parsed.Path, "/") + oidcWellKnownPath
	parsed.RawQuery = ""
	parsed.Fragment = ""

	return parsed.String(), nil
}

// fetchOIDCDiscovery retrieves and validates the discovery document for issuer.
func fetchOIDCDiscovery(ctx context.Context, issuer string, httpClient *http.Client) (*oidcDiscoveryDocument, error) {
	discoveryURL, err := oidcDiscoveryURL(issuer)
	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, discoveryURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch OIDC discovery document from %s: %w", discoveryURL, err)
	}
	defer func() { _ = res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read OIDC discovery document from %s: %w", discoveryURL, err)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("OIDC discovery request to %s failed with status %d", discoveryURL, res.StatusCode)
	}

	var document oidcDiscoveryDocument
	if err := json.Unmarshal(body, &document); err != nil {
		return nil, fmt.Errorf("failed to parse OIDC discovery document from %s: %w", discoveryURL, err)
	}

	if err := document.validate(); err != nil {
		return nil, err
	}

	return &document, nil
}
