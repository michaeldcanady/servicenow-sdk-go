package snauth

import (
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/auth"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

type serviceNowROPCConfig struct {
	clientID     string
	clientSecret string
	refreshToken string
}

func NewROPCAuthorizationStrategy(
	clientID string,
	clientSecret string,
	decoder TokenDecoder,
	instance string,
	// TODO: make these options
	cache auth.Cache[oauth2.Oauth2Token],
	client *http.Client,
) {
	oauth2.NewOauth2TokenStrategy(
		// TODO: build out tokenURL
		NewROPCTokenStrategy(client, decoder, instance),
		cache,
		serviceNowROPCConfig{
			clientID:     clientID,
			clientSecret: clientSecret,
		},
		serviceNowOptionsGenerator,
	)
}

func serviceNowOptionsGenerator(config serviceNowROPCConfig) []oauth2.TokenOption[ROPCTokenConfig] {
	opts := make([]ROPCTokenOption, 0, 5)

	if config.clientID != "" {
		opts = append(opts, WithClientID(config.clientID))
	}

	if config.clientSecret != "" {
		opts = append(opts, WithClientSecret(config.clientSecret))
	}

	return opts
}
