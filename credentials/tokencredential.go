package credentials

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// TokenCredential represents the OAuth2 token credentials.
type TokenCredential struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	Token        *AccessToken
	BaseURL      string
	server       *HTTPServer
}

func NewTokenCredential(clientID, clientSecret, username, password, baseURL string) *TokenCredential {

	address := "localhost:5000"

	return &TokenCredential{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		BaseURL:      baseURL,
		server:       NewHTTPServer(address),
	}
}

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
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", tc.ClientID)
	data.Set("client_secret", tc.ClientSecret)
	data.Set("username", tc.Username)
	data.Set("password", tc.Password)

	// Make a POST request to [baseUrl]/oauth.do.
	oauthURL := fmt.Sprintf("%s/oauth_token.do", tc.BaseURL)
	resp, err := http.PostForm(oauthURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OAuth2 token request failed with status code %d", resp.StatusCode)
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
		var errVal interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errVal); err != nil {
			return nil, err
		}
		fmt.Println(errVal)
		return nil, fmt.Errorf("OAuth2 token request failed with status code %d", resp.StatusCode)
	}

	var accessToken AccessToken
	if err := json.NewDecoder(resp.Body).Decode(&accessToken); err != nil {
		return nil, err
	}

	accessToken.ExpiresAt = time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second)

	return &accessToken, nil
}
