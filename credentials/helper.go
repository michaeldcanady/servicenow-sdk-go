package credentials

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

func decodeAccessToken(response *http.Response) (*AccessToken, error) {
	defer response.Body.Close() //nolint:errcheck
	var accessToken AccessToken
	if err := json.NewDecoder(response.Body).Decode(&accessToken); err != nil {
		return nil, err
	}

	accessToken.ExpiresAt = time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second)
	return &accessToken, nil
}

func convertToken(t *oauth2.Token) *AccessToken {
	if t == nil {
		return nil
	}
	return &AccessToken{
		AccessToken:  t.AccessToken,
		ExpiresIn:    int(t.ExpiresIn),
		ExpiresAt:    time.Now().Add(time.Duration(t.ExpiresIn) * time.Second),
		RefreshToken: t.RefreshToken,
		Scope:        t.Scope,
		TokenType:    t.TokenType,
	}
}
