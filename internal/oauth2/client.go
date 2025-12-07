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
)

type Client struct {
	ClientID     string
	ClientSecret string
	Endpoints    *Endpoints
	AuthMethod   AuthMethod
	HTTPClient   HTTPClient
}

func (c *Client) Exchange(ctx context.Context, params url.Values) (*Token, error) {
	if c.Endpoints == nil || strings.TrimSpace(c.Endpoints.TokenURL) == "" {
		return nil, errors.New("token endpoint is not set")
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
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
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
	tok.Headers = res.Header
	return &tok, nil
}

func (c *Client) httpClient() HTTPClient {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return http.DefaultClient
}

func (c *Client) ExchangeClientCredentials(ctx context.Context, scopes []string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeClientCreds)
	if len(scopes) > 0 {
		params.Set(ScopeKey, strings.Join(scopes, " "))
	}
	return c.Exchange(ctx, params)
}

func (c *Client) ExchangeRefreshToken(ctx context.Context, refresh string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeRefreshToken)
	params.Set(RefreshTokenKey, refresh)
	return c.Exchange(ctx, params)
}

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

func (c *Client) ExchangeCode(ctx context.Context, code, redirectURI, verifier string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeAuthCode)
	params.Set(CodeKey, code)
	if redirectURI != "" {
		params.Set(RedirectURIKey, redirectURI)
	}
	if verifier != "" {
		params.Set(CodeVerifierKey, verifier)
	}
	return c.Exchange(ctx, params)
}

func (c *Client) ExchangeJWT(ctx context.Context, assertion string) (*Token, error) {
	params := url.Values{}
	params.Set(GrantTypeKey, GrantTypeJWTBearer)
	params.Set(AssertionKey, assertion)
	return c.Exchange(ctx, params)
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
