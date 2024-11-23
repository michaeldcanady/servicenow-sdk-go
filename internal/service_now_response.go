package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

const (
	resultKey          = "Result"
	firstLinkHeaderKey = "first"
	prevLinkHeaderKey  = "prev"
	nextLinkHeaderKey  = "next"
	lastLinkHeaderKey  = "last"
)

type ServiceNowResponse[T serialization.Parsable] interface {
	GetResult() (T, error)
	setResult(T) error
	serialization.Parsable
	store.BackedModel
}

type serviceNowResponse[T serialization.Parsable] struct {
	factory      serialization.ParsableFactory
	backingStore store.BackingStore
}

func NewServiceNowResponse[T serialization.Parsable](factory serialization.ParsableFactory) ServiceNowResponse[T] {
	return &serviceNowResponse[T]{
		factory:      factory,
		backingStore: store.BackingStoreFactoryInstance(),
	}
}

func CreateServiceNowResponseFromDiscriminatorValue[T serialization.Parsable](factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewServiceNowResponse[T](factory), nil
	}
}

func (tE *serviceNowResponse[T]) GetBackingStore() store.BackingStore {
	return tE.backingStore
}

// Serialize writes the objects properties to the current writer.
func (tE *serviceNowResponse[T]) Serialize(writer serialization.SerializationWriter) error {
	return nil
}

func (tE *serviceNowResponse[T]) GetResult() (T, error) {
	var res T

	if IsNil(tE) {
		return res, nil
	}

	val, err := tE.GetBackingStore().Get(resultKey)
	if err != nil {
		return res, err
	}

	typedVal, ok := val.(T)
	if !ok {
		return res, errors.New("val is not serialization.Parsable")
	}

	return typedVal, nil
}

func (tE *serviceNowResponse[T]) setResult(result T) error {
	if IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(resultKey, result)
}

func (tE *serviceNowResponse[T]) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		resultKey: func(pn serialization.ParseNode) error {
			val, err := pn.GetObjectValue(tE.factory)
			if err != nil {
				return nil
			}

			elem, ok := val.(T)
			if !ok {
				return errors.New("elem is not serialization.Parsable")
			}

			if err := tE.setResult(elem); err != nil {
				return err
			}

			return nil
		},
	}
}
