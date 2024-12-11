package authentication

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	grantTypeKey = "grant_type"
)

type grantTypeRequestable interface {
	GetGrantType() (*grantType, error)
	setGrantType(*grantType) error
	serialization.Parsable
	store.BackedModel
}

type grantTypeRequest struct {
	backingStoreFactory store.BackingStoreFactory
	backingStore        store.BackingStore
}

type grantTypeRequestOption func(*grantTypeRequest)

// WithBackingStoreFactory sets a custom backing store factory.
//func withBackingStoreFactory(factory store.BackingStoreFactory) grantTypeRequestOption {
//	return func(req *grantTypeRequest) {
//		req.backingStoreFactory = factory
//	}
//}

func newGrantTypeRequest(opts ...grantTypeRequestOption) grantTypeRequestable {
	req := &grantTypeRequest{
		backingStoreFactory: store.NewInMemoryBackingStore,
	}

	for _, opt := range opts {
		opt(req)
	}

	req.backingStore = req.backingStoreFactory()

	return req
}

// GetBackingStore returns the store that is backing the model.
func (request *grantTypeRequest) GetBackingStore() store.BackingStore {
	if internal.IsNil(request) {
		return nil
	}
	if internal.IsNil(request.backingStore) {
		request.backingStore = request.backingStoreFactory()
	}

	return request.backingStore
}

// Serialize writes the objects properties to the current writer.
func (request *grantTypeRequest) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(request) {
		return nil
	}

	fieldSerializers := []func(serialization.SerializationWriter) error{
		func(writer serialization.SerializationWriter) error {
			grantType, err := request.GetGrantType()
			if err != nil {
				return err
			}

			rawGrantType := grantType.String()

			return writer.WriteStringValue(grantTypeKey, &rawGrantType)
		},
	}

	for _, fieldSerializer := range fieldSerializers {
		if err := fieldSerializer(writer); err != nil {
			return err
		}
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (request *grantTypeRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func (request *grantTypeRequest) GetGrantType() (*grantType, error) {
	if internal.IsNil(request) {
		return nil, nil
	}

	unknownGrantType, err := request.GetBackingStore().Get(grantTypeKey)
	if err != nil {
		return nil, err
	}

	typedGrantType, ok := unknownGrantType.(*grantType)
	if !ok {
		return nil, errors.New("grantType is not grantType")
	}

	return typedGrantType, nil
}

func (request *grantTypeRequest) setGrantType(grantType *grantType) error {
	if internal.IsNil(request) {
		return nil
	}

	return request.GetBackingStore().Set(grantTypeKey, grantType)
}
