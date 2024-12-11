package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type oidcRefreshAuthRecordable interface {
	oauth2AuthRecordable
	Refreshable
}

type authRecordable interface {
	GetRecordType() (*string, error)
	SetRecordType(*string) error
	store.BackedModel
	serialization.Parsable
}

type Refreshable interface {
	GetRefreshToken() (*string, error)
	SetRefreshToken(*string) error
}

type authRecord struct {
	backingStoreFactory store.BackingStoreFactory
	backingStore        store.BackingStore
}

func newAuthRecordFromAuthenticationTokenResponsable(response authenticationTokenResponsable) (authRecordable, error) {
	authRecord := newAuthRecord(store.NewInMemoryBackingStore)

}

func newAuthRecord(storeFactory store.BackingStoreFactory) authRecordable {
	return &authRecord{
		backingStoreFactory: storeFactory,
		backingStore:        storeFactory(),
	}
}

func newRefreshableAuthRecord(storeFactory store.BackingStoreFactory) authRecordable {
	return &authRecord{
		backingStoreFactory: storeFactory,
		backingStore:        storeFactory(),
	}
}

// GetBackingStore retrieves the backing store for the model.
func (aR *authRecord) GetBackingStore() store.BackingStore {
	if internal.IsNil(aR) {
		return nil
	}

	if internal.IsNil(aR.backingStore) {
		aR.backingStore = aR.backingStoreFactory()
	}

	return aR.backingStore
}

// Serialize writes the objects properties to the current writer.
func (aR *authRecord) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(aR) {
		return nil
	}

	return errors.New("not implemented")
}

// GetFieldDeserializers returns the deserialization information for this object.
func (aR *authRecord) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (aR *authRecord) GetRecordType() (*string, error) {
	if internal.IsNil(aR) {
		return nil, nil
	}

	recordType, err := aR.GetBackingStore().Get("recordType")
	if err != nil {
		return nil, err
	}

	typedRecordType, ok := recordType.(*string)
	if !ok {
		return nil, errors.New("recordType is not *string")
	}

	return typedRecordType, nil
}

func (aR *authRecord) SetRecordType(recordType *string) error {
	if internal.IsNil(aR) {
		return nil
	}

	return aR.GetBackingStore().Set("recordType", recordType)
}
