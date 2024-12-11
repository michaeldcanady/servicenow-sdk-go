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

var _ authentication.AccessTokenProvider = (*resourceOwnerPasswordTokenProvider)(nil)

type refreshable interface {
	GetRefreshToken() (*string, error)
}

type resourceOwnerPasswordTokenProvider struct {
	flow        *resourceOwnerPasswordFlow
	refreshFlow *refreshTokenFlow
	cache       Cache
}

func newResourceOwnerPasswordTokenProvider(clientID, clientSecret, username, password string, cache Cache) *resourceOwnerPasswordTokenProvider {
	return &resourceOwnerPasswordTokenProvider{
		flow:        newResourceOwnerPasswordFlow(clientID, newRequestAdapter(), clientSecret, username, password),
		refreshFlow: newRefreshTokenFlow(newRequestAdapter(), clientID, clientSecret),
		cache:       cache,
	}
}

func (provider *resourceOwnerPasswordTokenProvider) getCachedAuthRecord() (authRecordable, error) {
	if internal.IsNil(provider) {
		return nil, nil
	}

	jsonCredential, err := provider.cache.Get(tokenKey)
	if err != nil {
		return nil, err
	}

	node, err := jsonserialization.NewJsonParseNode([]byte(jsonCredential))
	if err != nil {
		return nil, err
	}

	potentialRecord, err := node.GetObjectValue(CreateAuthenticationTokenResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}

	record, ok := potentialRecord.(authRecordable)
	if !ok {
		return nil, errors.New("potentialRecord is not AuthRecordable")
	}

	return record, nil
}

func (provider *resourceOwnerPasswordTokenProvider) setCachedAuthRecord(record authRecordable) error {
	if internal.IsNil(provider) {
		return nil
	}

	writer := jsonserialization.NewJsonSerializationWriter()

	if err := record.Serialize(writer); err != nil {
		return err
	}

	content, err := writer.GetSerializedContent()
	if err != nil {
		return err
	}

	return provider.cache.Set(tokenKey, string(content))
}

func (provider *resourceOwnerPasswordTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", nil
	}

	authRecord, err := provider.getCachedAuthRecord()
	if err != nil {
		return "", err
	}

	properAuthRecord, ok := authRecord.(oauth2AuthRecordable)
	if !ok {
		return "", errors.New("authRecord is not oauth2AuthRecordable")
	}

	expirationDate, err := properAuthRecord.GetExpirationDate()
	if err != nil {
		return "", err
	}

	// token is still valid
	if !internal.IsNil(authRecord) && time.Now().Before(*expirationDate) {
		token, err := properAuthRecord.GetAccessToken()
		if err != nil {
			return "", err
		}

		return *token, nil
	}

	// retrieve auth record if record is nil
	if internal.IsNil(authRecord) {
		response, err := provider.flow.AcquireAuthRecord(ctx, uri, additionalAuthenticationContext)
		if err != nil {
			return "", err
		}
		authRecord, err = newAuthRecordFromAuthenticationTokenResponsable(response)
		if err != nil {
			return "", err
		}
	} else {
		refreshable, ok := (authRecord).(refreshable)
		if !ok {
			return "", errors.New("existing auth record isn't refreshable")
		}

		token, err := refreshable.GetRefreshToken()
		if err != nil {
			return "", err
		}

		response, err := provider.refreshFlow.AcquireAuthRecord(ctx, uri, additionalAuthenticationContext, *token)
		if err != nil {
			return "", err
		}
		authRecord, err = newAuthRecordFromAuthenticationTokenResponsable(response)
		if err != nil {
			return "", err
		}
		properAuthRecord, ok = authRecord.(oauth2AuthRecordable)
		if !ok {
			return "", errors.New("authRecord is not oauth2AuthRecordable")
		}
	}

	if provider.setCachedAuthRecord(authRecord) != nil {
		return "", err
	}

	token, err := properAuthRecord.GetAccessToken()
	if err != nil {
		return "", err
	}

	return *token, nil
}

// GetAllowedHostsValidator returns the hosts validator.
func (provider *resourceOwnerPasswordTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
