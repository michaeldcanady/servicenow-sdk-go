package authentication

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
)

var _ authentication.AccessTokenProvider = (*externalJWTTokenProvider)(nil)

const (
	tokenKey = ""
)

type externalJWTTokenProvider struct {
	flow        *jwtBearerFlow
	refreshFlow *refreshTokenFlow
	cache       Cache
}

func newExternalJWTTokenProvider(tokenProvider authentication.AccessTokenProvider, cache Cache, clientID, clientSecret string) *externalJWTTokenProvider {
	return &externalJWTTokenProvider{
		flow:        newJWTBearerFlow(tokenProvider, newRequestAdapter(), clientID, clientSecret),
		refreshFlow: newRefreshTokenFlow(newRequestAdapter(), clientID, clientSecret),
		cache:       cache,
	}
}

// GetAuthorizationToken returns the access token for the provided url.
func (provider *externalJWTTokenProvider) GetAuthorizationToken(context context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	jsonCredential, err := provider.cache.Get(tokenKey)
	if err != nil {
		return "", err
	}

	if jsonCredential == "" {
		record, err := provider.flow.AcquireAuthRecord(context, uri, additionalAuthenticationContext)
		if err != nil {
			return "", err
		}

		writer := jsonserialization.NewJsonSerializationWriter()

		if err := record.Serialize(writer); err != nil {
			return "", err
		}

		content, err := writer.GetSerializedContent()
		if err != nil {
			return "", err
		}

		if err := provider.cache.Set(tokenKey, string(content)); err != nil {
			return "", err
		}

		accessToken, err := record.GetAccessToken()
		if err != nil {
			return "", err
		}
		return *accessToken, nil
	}

	node, err := jsonserialization.NewJsonParseNode([]byte(jsonCredential))
	if err != nil {
		return "", err
	}

	potentialRecord, err := node.GetObjectValue(CreateAuthenticationTokenResponseFromDiscriminatorValue)
	if err != nil {
		return "", err
	}

	record, ok := potentialRecord.(authenticationTokenResponsable)
	if !ok {
		return "", errors.New("potentialRecord is not authenticationTokenResponsable")
	}

	if record.GetExpiresIn() < time.Now() {

		token, err := record.GetAccessToken()
		if err != nil {
			return "", err
		}

		return *token, nil
	}

	token, err := record.GetRefreshToken()
	if err != nil {
		return "", err
	}

	record, err = provider.refreshFlow.AcquireAuthRecord(context, uri, additionalAuthenticationContext, *token)
	if err != nil {
		return "", err
	}

	writer := jsonserialization.NewJsonSerializationWriter()

	if err := record.Serialize(writer); err != nil {
		return "", err
	}

	content, err := writer.GetSerializedContent()
	if err != nil {
		return "", err
	}

	if err := provider.cache.Set(tokenKey, string(content)); err != nil {
		return "", err
	}

	accessToken, err := record.GetAccessToken()
	if err != nil {
		return "", err
	}
	return *accessToken, nil
}

// GetAllowedHostsValidator returns the hosts validator.
func (provider *externalJWTTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
