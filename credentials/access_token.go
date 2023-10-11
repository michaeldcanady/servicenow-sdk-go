package credentials

import "time"

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresAt    time.Time
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

//IsExpired Checks if the access token is expired
func (t *AccessToken) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}
