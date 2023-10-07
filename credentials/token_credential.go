package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/abstraction"
)

// TokenCredential represents the OAuth2 token credentials.
type TokenCredential struct {
	ClientID     string
	ClientSecret string
	Token        *AccessToken
	BaseURL      string
	Prompt       func() (string, string, error)
	server       *HTTPServer
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
func NewTokenCredential(clientId, clientSecret, baseURL string, prompt func() (string, string, error)) (*TokenCredential, error) {

	address := "localhost:5000"

	if prompt == nil {
		prompt = DefaultPrompt
	}

	if clientId == "" {
		return nil, errors.New("clientId is empty")
	}

	if clientSecret == "" {
		return nil, errors.New("clientSecret is empty")
	}

	if baseURL == "" {
		return nil, errors.New("baseURL is empty")
	}

	return &TokenCredential{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
		Prompt:       prompt,
		server:       NewHTTPServer(address),
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

func (tc *TokenCredential) retrieveOAuthToken() (*AccessToken, error) {

	username, password, err := tc.promptUser()
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", tc.ClientID)
	data.Set("client_secret", tc.ClientSecret)
	data.Set("username", username)
	data.Set("password", password)

	// Make a POST request to [baseUrl]/oauth.do.
	oauthURL := fmt.Sprintf("%s/oauth_token.do", tc.BaseURL)
	resp, err := http.PostForm(oauthURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OAuth2 token request failed with status code? %d", resp.StatusCode)
	}

	var accessToken AccessToken
	if err := json.NewDecoder(resp.Body).Decode(&accessToken); err != nil {
		return nil, err
	}

	accessToken.ExpiresAt = time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second)

	return &accessToken, nil
}

func (tc *TokenCredential) refreshOAuthToken() (*AccessToken, error) {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", tc.ClientID)
	data.Set("client_secret", tc.ClientSecret)
	data.Set("refresh_token ", tc.Token.RefreshToken)

	// Make a POST request to [baseUrl]/oauth.do.
	oauthURL := fmt.Sprintf("%s/oauth_token.do", tc.BaseURL)
	resp, err := http.PostForm(oauthURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errVal map[string]string
		if err := json.NewDecoder(resp.Body).Decode(&errVal); err != nil {
			return nil, err
		}

		return nil, &abstraction.ServiceNowError{
			Exception: abstraction.Exception{
				Message: errVal["error"],
				Detail:  errVal["error_description"],
			},
		}
	}

	var accessToken AccessToken
	if err := json.NewDecoder(resp.Body).Decode(&accessToken); err != nil {
		return nil, err
	}

	accessToken.ExpiresAt = time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second)

	return &accessToken, nil
}
