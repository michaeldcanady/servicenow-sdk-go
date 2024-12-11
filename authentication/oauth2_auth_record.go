package authentication

import (
	"errors"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type oauth2AuthRecordable interface {
	authRecordable
	GetAccessToken() (*string, error)
	SetAccessToken(*string) error
	GetExpirationDate() (*time.Time, error)
	SetExpirationDate(*time.Time) error
}

type oauth2AuthRecord struct {
	authRecordable
}

func newOIDCAuthRecord(storeFactory store.BackingStoreFactory) oauth2AuthRecordable {
	return &oauth2AuthRecord{
		newAuthRecord(storeFactory),
	}
}

// Serialize writes the objects properties to the current writer.
func (aR *oauth2AuthRecord) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(aR) {
		return nil
	}

	return errors.New("not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (aR *oauth2AuthRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (aR *oauth2AuthRecord) GetAccessToken() (*string, error) {
	if internal.IsNil(aR) {
		return nil, nil
	}

	accessToken, err := aR.GetBackingStore().Get("accessToken")
	if err != nil {
		return nil, err
	}

	typedAccessToken, ok := accessToken.(*string)
	if !ok {
		return nil, errors.New("accessToken is not *string")
	}

	return typedAccessToken, nil
}

func (aR *oauth2AuthRecord) GetExpirationDate() (*time.Time, error) {
	if internal.IsNil(aR) {
		return nil, nil
	}

	expirationDate, err := aR.GetBackingStore().Get("expirationDate")
	if err != nil {
		return nil, err
	}

	typedExpirationDate, ok := expirationDate.(*time.Time)
	if !ok {
		return nil, errors.New("expirationDate is not *string")
	}

	return typedExpirationDate, nil
}

func (aR *oauth2AuthRecord) GetRefreshToken() (*string, error) {
	if internal.IsNil(aR) {
		return nil, nil
	}

	refreshToken, err := aR.GetBackingStore().Get("refreshToken")
	if err != nil {
		return nil, err
	}

	typedRefreshToken, ok := refreshToken.(*string)
	if !ok {
		return nil, errors.New("refreshToken is not *string")
	}

	return typedRefreshToken, nil
}

func (aR *oauth2AuthRecord) SetAccessToken(accessToken *string) error {
	if internal.IsNil(aR) {
		return nil
	}
	return aR.GetBackingStore().Set("accessToken", accessToken)
}

func (aR *oauth2AuthRecord) SetExpirationDate(accessToken *time.Time) error {
	if internal.IsNil(aR) {
		return nil
	}
	return aR.GetBackingStore().Set("expirationDate", accessToken)
}

func (aR *oauth2AuthRecord) SetRefreshToken(refreshToken *string) error {
	if internal.IsNil(aR) {
		return nil
	}
	return aR.GetBackingStore().Set("refreshToken", refreshToken)
}
