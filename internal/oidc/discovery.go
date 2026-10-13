// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oidc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
)

// DiscoveryDocument represents the OIDC discovery document as defined in RFC 8414.
// Only the fields used by the authorization code + PKCE flow are included.
// The JWKS URL is intentionally ignored per ADR-014.
type DiscoveryDocument struct {
	// Issuer is the URL of the OIDC provider's issuer identifier.
	Issuer string `json:"issuer"`
	// AuthorizationEndpoint is the URL of the authorization server's authorization endpoint.
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	// TokenEndpoint is the URL of the authorization server's token endpoint.
	TokenEndpoint string `json:"token_endpoint"`
	// RevocationEndpoint is the URL of the authorization server's token revocation endpoint (optional).
	RevocationEndpoint string `json:"revocation_endpoint,omitempty"`
}

// Discover fetches and validates the OIDC discovery document from the given issuer URL.
// The issuerURL can be either the base issuer URL (e.g., "https://idp.example.com") or
// a full metadata URL (e.g., "https://idp.example.com/.well-known/openid-configuration").
// It uses the provided HTTP client for fetching the document and validates the response.
// Returns an *oauth2.Endpoints populated with the discovered endpoints, or an error wrapped
// with the shared OIDC discovery sentinels for matchable error handling.
func Discover(ctx context.Context, issuerURL string, httpClient oauth2.HTTPClient) (*oauth2.Endpoints, error) {
	metadataURL := resolveMetadataURL(issuerURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, metadataURL, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", snerrors.ErrOIDCDiscoveryFailed, err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", snerrors.ErrOIDCDiscoveryFailed, err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %w", snerrors.ErrOIDCDiscoveryFailed, fmt.Errorf("HTTP %d", resp.StatusCode))
	}

	var doc DiscoveryDocument
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("%w: %w", snerrors.ErrOIDCDiscoveryFailed, err)
	}

	// Validate required fields
	if strings.TrimSpace(doc.Issuer) == "" {
		return nil, fmt.Errorf("%w: issuer", snerrors.ErrOIDCDiscoveryMissingEndpoint)
	}
	if strings.TrimSpace(doc.AuthorizationEndpoint) == "" {
		return nil, fmt.Errorf("%w: authorization_endpoint", snerrors.ErrOIDCDiscoveryMissingEndpoint)
	}
	if strings.TrimSpace(doc.TokenEndpoint) == "" {
		return nil, fmt.Errorf("%w: token_endpoint", snerrors.ErrOIDCDiscoveryMissingEndpoint)
	}

	// Validate issuer matches the requested issuer (normalized comparison)
	if !issuerMatches(doc.Issuer, issuerURL) {
		return nil, fmt.Errorf("%w: expected %q, got %q", snerrors.ErrOIDCDiscoveryIssuerMismatch, normalizeIssuer(issuerURL), doc.Issuer)
	}

	// Build the oauth2.Endpoints struct with discovered URLs
	endpoints := &oauth2.Endpoints{
		AuthURL:          doc.AuthorizationEndpoint,
		TokenURL:         doc.TokenEndpoint,
		RevocationURL:    doc.RevocationEndpoint,
		DeviceURL:        "",
		IntrospectionURL: "",
	}

	return endpoints, nil
}

// resolveMetadataURL constructs the OIDC discovery metadata URL from the issuer URL.
// If the input already ends with the well-known path, it's returned as-is.
// Otherwise, the standard ".well-known/openid-configuration" path is appended.
func resolveMetadataURL(issuerURL string) string {
	trimmed := strings.TrimSuffix(issuerURL, "/")
	const wellKnownPath = "/.well-known/openid-configuration"
	if strings.HasSuffix(trimmed, wellKnownPath) {
		return trimmed
	}
	return trimmed + wellKnownPath
}

// issuerMatches checks if the document's issuer matches the expected issuer.
// Comparison is done after normalizing both URLs (removing trailing slashes).
func issuerMatches(docIssuer, expectedIssuer string) bool {
	return normalizeIssuer(docIssuer) == normalizeIssuer(expectedIssuer)
}

// normalizeIssuer normalizes an issuer URL by removing trailing slashes and lowercasing the host.
func normalizeIssuer(issuer string) string {
	trimmed := strings.TrimSuffix(strings.TrimSpace(issuer), "/")
	// Parse to normalize host casing
	if u, err := url.Parse(trimmed); err == nil && u.Host != "" {
		// Rebuild with normalized host
		return strings.TrimSuffix(trimmed, u.Host) + u.Host + strings.TrimPrefix(trimmed, u.Scheme+"://"+u.Host)
	}
	return trimmed
}
