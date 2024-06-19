package credentials

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenCredential represents the OAuth2 token credentials.
type TokenCredential struct {
	ClientID     string
	ClientSecret string
	Token        *AccessToken
	BaseURL      string
	Prompt       func() (string, string, error)
	server       *HTTPServer
	// Specific to Inbound JWT
	KeyID      string
	User       string
	PrivateKey *rsa.PrivateKey
}

// DefaultPrompt is a default implementation of the user prompt function.
func DefaultPrompt() (string, string, error) {
	var username, password string

	fmt.Print("Enter username: ")
	_, _ = fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	_, _ = fmt.Scanln(&password)

	return username, password, nil
}

// NewTokenCredential creates a new token credential
func NewTokenCredential(clientID, clientSecret, baseURL string, prompt func() (string, string, error)) (*TokenCredential, error) {
	address := "localhost:5000"

	if prompt == nil {
		prompt = DefaultPrompt
	}

	if clientID == "" {
		return nil, EmptyClientID
	}

	if clientSecret == "" {
		return nil, EmptyClientSecret
	}

	if baseURL == "" {
		return nil, EmptyBaseURL
	}

	return &TokenCredential{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
		Prompt:       prompt,
		server:       NewHTTPServer(address),
	}, nil
}

// NewJwtTokenCredential creates a new token credential using Inbound JWT authentication
func NewJwtTokenCredential(clientID, clientSecret, kid, baseURL, user, privateKeyPath string) (*TokenCredential, error) {
	if clientID == "" {
		return nil, EmptyClientID
	}

	if clientSecret == "" {
		return nil, EmptyClientSecret
	}

	if baseURL == "" {
		return nil, EmptyBaseURL
	}

	if kid == "" {
		return nil, EmptyKeyID
	}

	if user == "" {
		return nil, EmptyUser
	}

	if privateKeyPath == "" {
		return nil, EmptyPrivateKeyPath
	}

	fd, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(fd)
	if err != nil {
		return nil, err
	}

	return &TokenCredential{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
		KeyID:        kid,
		User:         user,
		PrivateKey:   privateKey,
	}, nil
}

func (tc *TokenCredential) promptUser() (string, string, error) {
	// Use the provided function to get the username and password from the user.
	username, password, err := tc.Prompt()
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}

func (tc *TokenCredential) generateJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &jwt.RegisteredClaims{
		Audience:  jwt.ClaimStrings{tc.ClientID},
		Issuer:    tc.ClientID,
		Subject:   tc.User,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	})

	token.Header["kid"] = tc.KeyID
	return token.SignedString(tc.PrivateKey)
}

// GetAuthentication gets the authentication header value
func (tc *TokenCredential) GetAuthentication() (string, error) {
	if tc.Token == nil {
		token, err := tc.retrieveOAuthToken()
		if err != nil {
			return "", err
		}
		tc.Token = token
	} else if tc.Token.IsExpired() {
		token, err := tc.refreshOAuthToken()
		if err != nil {
			return "", err
		}
		tc.Token = token
	}

	return fmt.Sprintf("Bearer %s", tc.Token.AccessToken), nil
}

func (tc *TokenCredential) GetOauth2Url() string {
	return fmt.Sprintf("%s/oauth_token.do", tc.BaseURL)
}

func (tc *TokenCredential) requestToken(data url.Values) (*AccessToken, error) {
	if tc.server != nil {
		tc.server.Start()
	}

	oauthURL := tc.GetOauth2Url()
	req, err := http.NewRequest("POST", oauthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OAuth2 token request failed with status code? %d", resp.StatusCode)
	}

	AccessToken, err := decodeAccessToken(resp)
	if err != nil {
		return nil, err
	}

	return AccessToken, nil
}

func (tc *TokenCredential) retrieveOAuthToken() (*AccessToken, error) {
	data := url.Values{}
	data.Set("client_id", tc.ClientID)
	data.Set("client_secret", tc.ClientSecret)
	if tc.KeyID == "" {
		username, password, err := tc.promptUser()
		if err != nil {
			return nil, err
		}

		data.Set("grant_type", "password")
		data.Set("username", username)
		data.Set("password", password)
	} else {
		data.Set("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
		jwtToken, err := tc.generateJwtToken()
		if err != nil {
			return nil, err
		}
		data.Set("assertion", jwtToken)
	}

	return tc.requestToken(data)
}

func (tc *TokenCredential) refreshOAuthToken() (*AccessToken, error) {
	// JWT doesn't support refresh token, just fetch a new one
	if tc.KeyID != "" {
		return tc.retrieveOAuthToken()
	}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", tc.ClientID)
	data.Set("client_secret", tc.ClientSecret)
	data.Set("refresh_token ", tc.Token.RefreshToken)

	return tc.requestToken(data)
}
