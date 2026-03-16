package oauth2

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client is a generic OAuth2 client that handles various grant types and management operations.
type Client struct {
	// ClientID is the public identifier for the application.
	ClientID string
	// ClientSecret is the secret identifier for the application (confidential clients only).
	ClientSecret string //nolint:gosec // G117: Needed for flow, no secret
	// Endpoints holds the URLs for the various OAuth2 service endpoints.
	Endpoints *Endpoints
	// AuthMethod specifies how the client identifies itself to the authorization server.
	AuthMethod AuthMethod
	// HTTPClient is the underlying client used to make HTTP requests.
	HTTPClient HTTPClient
	// ErrorParser is an optional hook to parse errors from the authorization server.
	ErrorParser func(statusCode int, body []byte) error
}

// Exchange performs a generic token exchange by sending the provided parameters to the token endpoint.
func (c *Client) Exchange(ctx context.Context, params url.Values) (*Token, error) {
	if err := c.Endpoints.Validate(params.Get(GrantTypeKey)); err != nil {
		return nil, err
	}

	// Always include client_id
	if params.Get(ClientIDKey) == "" && c.ClientID == "" {
		return nil, errors.New("client id is empty")
	}

	if params.Get(ClientIDKey) == "" && c.ClientID != "" {
		params.Set(ClientIDKey, c.ClientID)
	}

	var authHeader string

	// Switch on auth method
	switch c.AuthMethod {
	case AuthMethodClientSecretBasic:
		if c.ClientID == "" || c.ClientSecret == "" {
			return nil, errors.New("client_secret_basic requires client_id and client_secret")
		}
		authHeader = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.ClientID+":"+c.ClientSecret))

	case AuthMethodClientSecretPost:
		if c.ClientSecret != "" {
			params.Set(ClientSecretKey, c.ClientSecret)
		}

	case AuthMethodNone:
		// No secret needed (public client)

	default:
		return nil, fmt.Errorf("unsupported auth method: %d", c.AuthMethod)
	}

	// Build request once
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Endpoints.TokenURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentTypeKey, FormURLEncodedContentType)
	req.Header.Set(AcceptKey, JSONContentType)
	if authHeader != "" {
		req.Header.Set(AuthorizationKey, authHeader)
	}

	req.Header.Set(ContentTypeKey, FormURLEncodedContentType)
	req.Header.Set(AcceptKey, JSONContentType)

	return c.doRequest(req)
}

func (c *Client) doRequest(req *http.Request) (*Token, error) {
	res, err := c.httpClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("token request failed: %w", err)
	}
	defer res.Body.Close() //nolint:errcheck

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if c.ErrorParser != nil {
			if err := c.ErrorParser(res.StatusCode, body); err != nil {
				return nil, err
			}
		}
		te := &TokenError{StatusCode: res.StatusCode, RawBody: string(body)}
		_ = json.Unmarshal(body, te)
		return nil, te
	}

	var tok Token
	if err := json.Unmarshal(body, &tok); err != nil {
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(body, &raw); err == nil {
		tok.Raw = raw
	} else {
		tok.Raw = map[string]any{"_raw": string(body)}
	}
	return &tok, nil
}

var defaultHTTPClient = &http.Client{
	Timeout: 30 * time.Second,
}

func (c *Client) httpClient() HTTPClient {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return defaultHTTPClient
}

// ExchangeClientCredentials performs a Client Credentials Grant exchange.
func (c *Client) ExchangeClientCredentials(ctx context.Context, scopes []string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeClientCreds)
	if len(scopes) > 0 {
		params.Set(ScopeKey, strings.Join(scopes, " "))
	}
	return c.Exchange(ctx, params)
}

// ExchangeRefreshToken performs a Refresh Token Grant exchange.
func (c *Client) ExchangeRefreshToken(ctx context.Context, refresh string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeRefreshToken)
	params.Set(RefreshTokenKey, refresh)
	return c.Exchange(ctx, params)
}

// ExchangePassword performs a Resource Owner Password Credentials Grant exchange.
func (c *Client) ExchangePassword(ctx context.Context, user, pass string, scopes []string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypePassword)
	params.Set(UsernameKey, user)
	params.Set(PasswordKey, pass)
	if len(scopes) > 0 {
		params.Set(ScopeKey, strings.Join(scopes, " "))
	}
	return c.Exchange(ctx, params)
}

// ExchangeCode performs an Authorization Code Grant exchange.
func (c *Client) ExchangeCode(ctx context.Context, code, redirectURI, verifier, state string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeAuthCode)
	params.Set(CodeKey, code)
	if redirectURI != "" {
		params.Set(RedirectURIKey, redirectURI)
	}
	if state != "" {
		params.Set(StateKey, state)
	}
	if verifier != "" {
		params.Set(CodeVerifierKey, verifier)
	}
	return c.Exchange(ctx, params)
}

// ExchangeJWT performs a JWT Bearer Token Grant exchange.
func (c *Client) ExchangeJWT(ctx context.Context, assertion string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeJWTBearer)
	params.Set(AssertionKey, assertion)
	return c.Exchange(ctx, params)
}

// RequestDeviceAuthorization initiates the device authorization flow.
func (c *Client) RequestDeviceAuthorization(ctx context.Context, scopes []string) (*DeviceAuthorizationResponse, error) {
	if err := c.Endpoints.Validate(GrantTypeDeviceCode); err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set(ClientIDKey, c.ClientID)
	if len(scopes) > 0 {
		params.Set(ScopeKey, strings.Join(scopes, " "))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Endpoints.DeviceURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentTypeKey, FormURLEncodedContentType)
	req.Header.Set(AcceptKey, JSONContentType)

	res, err := c.httpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if c.ErrorParser != nil {
			if err := c.ErrorParser(res.StatusCode, body); err != nil {
				return nil, err
			}
		}
		te := &TokenError{StatusCode: res.StatusCode, RawBody: string(body)}
		_ = json.Unmarshal(body, te)
		return nil, te
	}

	var dar DeviceAuthorizationResponse
	if err := json.Unmarshal(body, &dar); err != nil {
		return nil, fmt.Errorf("failed to parse device authorization response: %w", err)
	}
	return &dar, nil
}

// ExchangeDeviceCode exchanges a device code for an access token.
func (c *Client) ExchangeDeviceCode(ctx context.Context, deviceCode string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeDeviceCode)
	params.Set(DeviceCodeKey, deviceCode)
	return c.Exchange(ctx, params)
}

// Revoke invalidates the provided token (RFC 7009).
func (c *Client) Revoke(ctx context.Context, token, tokenTypeHint string) error {
	if c.Endpoints == nil || strings.TrimSpace(c.Endpoints.RevocationURL) == "" {
		return errors.New("revocation endpoint is not set")
	}

	params := url.Values{}
	params.Set(TokenKey, token)
	if tokenTypeHint != "" {
		params.Set(TokenTypeHintKey, tokenTypeHint)
	}

	req, err := c.newAuthenticatedRequest(ctx, http.MethodPost, c.Endpoints.RevocationURL, params)
	if err != nil {
		return err
	}

	res, err := c.httpClient().Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		if c.ErrorParser != nil {
			if err := c.ErrorParser(res.StatusCode, body); err != nil {
				return err
			}
		}
		te := &TokenError{StatusCode: res.StatusCode, RawBody: string(body)}
		_ = json.Unmarshal(body, te)
		return te
	}

	return nil
}

// Introspect checks the status and metadata of a token (RFC 7662).
func (c *Client) Introspect(ctx context.Context, token, tokenTypeHint string) (*IntrospectionResponse, error) {
	if c.Endpoints == nil || strings.TrimSpace(c.Endpoints.IntrospectionURL) == "" {
		return nil, errors.New("introspection endpoint is not set")
	}

	params := url.Values{}
	params.Set(TokenKey, token)
	if tokenTypeHint != "" {
		params.Set(TokenTypeHintKey, tokenTypeHint)
	}

	req, err := c.newAuthenticatedRequest(ctx, http.MethodPost, c.Endpoints.IntrospectionURL, params)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		if c.ErrorParser != nil {
			if err := c.ErrorParser(res.StatusCode, body); err != nil {
				return nil, err
			}
		}
		te := &TokenError{StatusCode: res.StatusCode, RawBody: string(body)}
		_ = json.Unmarshal(body, te)
		return nil, te
	}

	var ir IntrospectionResponse
	if err := json.Unmarshal(body, &ir); err != nil {
		return nil, fmt.Errorf("failed to parse introspection response: %w", err)
	}
	_ = json.Unmarshal(body, &ir.Raw)
	return &ir, nil
}

func (c *Client) newAuthenticatedRequest(ctx context.Context, method, url string, params url.Values) (*http.Request, error) {
	if params.Get(ClientIDKey) == "" && c.ClientID != "" {
		params.Set(ClientIDKey, c.ClientID)
	}

	var authHeader string

	switch c.AuthMethod {
	case AuthMethodClientSecretBasic:
		if c.ClientID != "" && c.ClientSecret != "" {
			authHeader = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.ClientID+":"+c.ClientSecret))
		}
	case AuthMethodClientSecretPost:
		if c.ClientSecret != "" {
			params.Set(ClientSecretKey, c.ClientSecret)
		}
	case AuthMethodNone:
		// No secret needed
	}

	req, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentTypeKey, FormURLEncodedContentType)
	req.Header.Set(AcceptKey, JSONContentType)
	if authHeader != "" {
		req.Header.Set(AuthorizationKey, authHeader)
	}
	return req, nil
}

// AuthCodeURL builds the authorization URL for 3-legged flows.
func (c *Client) AuthCodeURL(redirectURI, state, codeChallenge, codeChallengeMethod string, scopes []string) (string, error) {
	if c.Endpoints == nil || strings.TrimSpace(c.Endpoints.AuthURL) == "" {
		return "", errors.New("authorization endpoint is not set")
	}

	u, err := url.Parse(c.Endpoints.AuthURL)
	if err != nil {
		return "", fmt.Errorf("invalid auth URL: %w", err)
	}

	q := u.Query()
	q.Set(ResponseTypeKey, ResponseTypeCode)
	q.Set(ClientIDKey, c.ClientID)
	if redirectURI != "" {
		q.Set(RedirectURIKey, redirectURI)
	}
	if state != "" {
		q.Set(StateKey, state)
	}
	if len(scopes) > 0 {
		q.Set(ScopeKey, strings.Join(scopes, " "))
	}
	if codeChallenge != "" {
		q.Set(CodeChallengeKey, codeChallenge)
		if codeChallengeMethod != "" {
			q.Set(CodeChallengeMethodKey, codeChallengeMethod)
		}
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
