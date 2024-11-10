package auth

import (
	"errors"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// AccessToken represents an OAuth2 access token and associated metadata.
type AccessToken struct {
	// AccessToken is the actual access token issued by the authorization server.
	AccessToken string `json:"access_token"`
	// ExpiresAt is the time at which the access token expires.
	ExpiresAt time.Time
	// Scope is a space-delimited list of permissions or access rights granted by the token.
	Scope []string `json:"scope"`
}

// IsExpired Checks if the access token is expired
func (t *AccessToken) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}

const (
	accessTokenKey  = "access_token"
	expiresInKey    = "expires_in"
	refreshTokenKey = "refresh_token"
	scopeKey        = "scope"
	tokenTypeKey    = "token_type"
)

type AccessTokenable interface {
	GetAccessToken() (*string, error)
	SetAccessToken(accessToken *string) error
	GetExpiresIn() (*int64, error)
	SetExpiresIn(expiresIn *int64) error
	GetRefreshToken() (*string, error)
	SetRefreshToken(refreshToken *string) error
	GetScopes() ([]string, error)
	SetScopes(scopes []string) error
	GetTokenType() (*string, error)
	SetTokenType(tokenType *string) error
	serialization.Parsable
	store.BackedModel
}

type accessToken struct {
	backingStore store.BackingStore
}

func NewAccessToken() AccessTokenable {
	return &accessToken{
		backingStore: store.NewInMemoryBackingStore(),
	}
}

func CreateAccessTokenFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewAccessToken(), nil
}

// Serialize writes the objects properties to the current writer.
func (p *accessToken) Serialize(writer serialization.SerializationWriter) error {
	return errors.New("Serialize is not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (p *accessToken) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		accessTokenKey: func(pn serialization.ParseNode) error {
			token, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return p.SetAccessToken(token)
		},
		expiresInKey: func(pn serialization.ParseNode) error {
			expiresIn, err := pn.GetInt64Value()
			if err != nil {
				return err
			}

			return p.SetExpiresIn(expiresIn)
		},
		refreshTokenKey: func(pn serialization.ParseNode) error {
			token, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return p.SetAccessToken(token)
		},
		scopeKey: func(pn serialization.ParseNode) error {
			scopes, err := pn.GetCollectionOfPrimitiveValues("string")
			if err != nil {
				return err
			}

			var typedScopes []string
			for _, scope := range scopes {
				typedScope, ok := scope.(string)
				if !ok {
					return errors.New("scopes is not []string")
				}
				typedScopes = append(typedScopes, typedScope)
			}

			return p.SetScopes(typedScopes)
		},
		tokenTypeKey: func(pn serialization.ParseNode) error {
			tokenType, err := pn.GetStringValue()
			if err != nil {
				return err
			}

			return p.SetAccessToken(tokenType)
		},
	}
}

func (p *accessToken) GetBackingStore() store.BackingStore {
	if core.IsNil(p) {
		return nil
	}

	if core.IsNil(p.backingStore) {
		p.backingStore = store.NewInMemoryBackingStore()
	}

	return p.backingStore
}

func (p *accessToken) GetAccessToken() (*string, error) {
	if core.IsNil(p) {
		return nil, nil
	}

	accessToken, err := p.GetBackingStore().Get(accessTokenKey)
	if err != nil {
		return nil, err
	}

	typedAccessToken, ok := accessToken.(*string)
	if !ok {
		return nil, errors.New("accessToken is not *string")
	}
	return typedAccessToken, nil
}

func (p *accessToken) SetAccessToken(accessToken *string) error {
	if core.IsNil(p) {
		return nil
	}

	return p.GetBackingStore().Set(accessTokenKey, accessToken)
}

func (p *accessToken) GetExpiresIn() (*int64, error) {
	if core.IsNil(p) {
		return nil, nil
	}

	expiresIn, err := p.GetBackingStore().Get(expiresInKey)
	if err != nil {
		return nil, err
	}

	typedExpiresIn, ok := expiresIn.(*int64)
	if !ok {
		return nil, errors.New("expiresIn is not *int")
	}
	return typedExpiresIn, nil
}

func (p *accessToken) SetExpiresIn(expiresIn *int64) error {
	if core.IsNil(p) {
		return nil
	}

	return p.GetBackingStore().Set(expiresInKey, expiresIn)
}

func (p *accessToken) GetRefreshToken() (*string, error) {
	if core.IsNil(p) {
		return nil, nil
	}

	refreshToken, err := p.GetBackingStore().Get(refreshTokenKey)
	if err != nil {
		return nil, err
	}

	typedRefreshToken, ok := refreshToken.(*string)
	if !ok {
		return nil, errors.New("refreshToken is not *string")
	}
	return typedRefreshToken, nil
}

func (p *accessToken) SetRefreshToken(refreshToken *string) error {
	if core.IsNil(p) {
		return nil
	}

	return p.GetBackingStore().Set(refreshTokenKey, refreshToken)
}

func (p *accessToken) GetScopes() ([]string, error) {
	if core.IsNil(p) {
		return nil, nil
	}

	scopes, err := p.GetBackingStore().Get(scopeKey)
	if err != nil {
		return nil, err
	}

	typedScopes, ok := scopes.([]string)
	if !ok {
		return nil, errors.New("scopes is not []string")
	}

	return typedScopes, nil
}

func (p *accessToken) SetScopes(scopes []string) error {
	if core.IsNil(p) {
		return nil
	}

	return p.GetBackingStore().Set(scopeKey, scopes)
}

func (p *accessToken) GetTokenType() (*string, error) {
	if core.IsNil(p) {
		return nil, nil
	}

	tokenType, err := p.GetBackingStore().Get(tokenTypeKey)
	if err != nil {
		return nil, err
	}

	typedTokenType, ok := tokenType.(*string)
	if !ok {
		return nil, errors.New("scopes is not *string")
	}

	return typedTokenType, nil
}

func (p *accessToken) SetTokenType(tokenType *string) error {
	if core.IsNil(p) {
		return nil
	}

	return p.GetBackingStore().Set(tokenTypeKey, tokenType)
}
