package credentials

import (
	"encoding/json"
	"net/http"
	"time"
)

func decodeAccessToken(response *http.Response) (*AccessToken, error) {
	defer response.Body.Close()
	var accessToken AccessToken
	if err := json.NewDecoder(response.Body).Decode(&accessToken); err != nil {
		return nil, err
	}

	accessToken.ExpiresAt = time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second)
	return &accessToken, nil
}
