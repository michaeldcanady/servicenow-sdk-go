package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ServiceNowCollectionResponse[T serialization.Parsable] interface {
	GetResult() ([]T, error)
	GetNextLink() (*string, error)
	GetPreviousLink() (*string, error)
	GetFirstLink() (*string, error)
	GetLastLink() (*string, error)
	setNextLink(*string) error
	setPreviousLink(*string) error
	setFirstLink(*string) error
	setLastLink(*string) error
	setResult([]T) error
	serialization.Parsable
	store.BackedModel
}

type serviceNowCollectionResponse[T serialization.Parsable] struct {
	factory      serialization.ParsableFactory
	backingStore store.BackingStore
}

func NewServiceNowCollectionResponse[T serialization.Parsable](factory serialization.ParsableFactory) ServiceNowCollectionResponse[T] {
	return &serviceNowCollectionResponse[T]{
		factory:      factory,
		backingStore: store.BackingStoreFactoryInstance(),
	}
}

func CreateServiceNowCollectionResponseFromDiscriminatorValue[T serialization.Parsable](factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewServiceNowCollectionResponse[T](factory), nil
	}
}

func (tE *serviceNowCollectionResponse[T]) GetBackingStore() store.BackingStore {
	if IsNil(tE) {
		return nil
	}

	if IsNil(tE.backingStore) {
		tE.backingStore = store.BackingStoreFactoryInstance()
	}

	return tE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (tE *serviceNowCollectionResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	if IsNil(tE) {
		return nil
	}
	return nil
}

func (tE *serviceNowCollectionResponse[T]) GetResult() ([]T, error) {
	if IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(resultKey)
	if err != nil {
		return nil, err
	}
	if IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.([]T)
	if !ok {
		return nil, errors.New("val is not serialization.Parsable")
	}

	return typedVal, nil
}

func (tE *serviceNowCollectionResponse[T]) setResult(result []T) error {
	if IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(resultKey, result)
}

func (tE *serviceNowCollectionResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if IsNil(tE) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
			elem, err := pn.GetCollectionOfObjectValues(tE.factory)
			if err != nil {
				return err
			}

			typedElems := make([]T, len(elem))
			for index, el := range elem {
				typedEl, ok := el.(T)
				if !ok {
					return errors.New("el is not serialization.Parsable")
				}
				typedElems[index] = typedEl
			}

			if err := tE.setResult(typedElems); err != nil {
				return err
			}

			return nil
		},
	}
}

func (tE *serviceNowCollectionResponse[T]) GetNextLink() (*string, error) {
	if IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(nextLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

func (tE *serviceNowCollectionResponse[T]) GetPreviousLink() (*string, error) {
	if IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(prevLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

func (tE *serviceNowCollectionResponse[T]) GetFirstLink() (*string, error) {
	if IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(firstLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

func (tE *serviceNowCollectionResponse[T]) GetLastLink() (*string, error) {
	if IsNil(tE) {
		return nil, nil
	}

	val, err := tE.GetBackingStore().Get(lastLinkHeaderKey)
	if err != nil {
		return nil, err
	}
	if IsNil(val) {
		return nil, nil
	}

	typedVal, ok := val.(*string)
	if !ok {
		return nil, errors.New("val is not *string")
	}

	return typedVal, nil
}

func (tE *serviceNowCollectionResponse[T]) setNextLink(nextLink *string) error {
	if IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(nextLinkHeaderKey, nextLink)
}

func (tE *serviceNowCollectionResponse[T]) setPreviousLink(previousLink *string) error {
	if IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(prevLinkHeaderKey, previousLink)
}

func (tE *serviceNowCollectionResponse[T]) setFirstLink(firstLink *string) error {
	if IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(firstLinkHeaderKey, firstLink)
}

func (tE *serviceNowCollectionResponse[T]) setLastLink(lastLink *string) error {
	if IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(lastLinkHeaderKey, lastLink)
}
