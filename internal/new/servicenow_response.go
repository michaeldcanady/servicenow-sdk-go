package internal

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// ServiceNowResponse represents a Service-Now response
type ServiceNowResponse[T any] interface {
	GetResult() (T, error)
	SetResult(T) error
	Serialize(writer serialization.SerializationWriter) error
	GetFactory() (serialization.ParsableFactory, error)
	store.BackedModel
}

type ServiceNowResponseImpl[T any] struct {
	factory             serialization.ParsableFactory
	backingStoreFactory store.BackingStoreFactory
	backingStore        store.BackingStore
}

func NewServiceNowResponse[T any](responseType ServiceNowResponseType, factory serialization.ParsableFactory, storeFactory store.BackingStoreFactory) *ServiceNowResponseImpl[T] {
	return &ServiceNowResponseImpl[T]{
		factory:             factory,
		backingStore:        storeFactory(),
		backingStoreFactory: storeFactory,
	}
}

// CreateServiceNowResponseFromDiscriminatorValue is a factory for creating a ServiceNowResponse
func CreateServiceNowResponseFromDiscriminatorValue(factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		resultNode, err := parseNode.GetChildNode(resultKey)
		if err != nil {
			return nil, err
		}
		value, err := resultNode.GetRawValue()
		if err != nil {
			return nil, err
		}
		switch v := reflect.ValueOf(value); v.Kind() {
		case reflect.Slice:
			return NewServiceNowCollectionResponse[serialization.Parsable](factory, store.BackingStoreFactoryInstance), nil
		case reflect.Struct:
			return NewServiceNowItemResponse[serialization.Parsable](factory, store.BackingStoreFactoryInstance), nil
		default:
			return nil, fmt.Errorf("unsupported kind (%s)", v)
		}
	}
}

// Serialize writes the objects properties to the current writer
func (tE *ServiceNowResponseImpl[T]) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(tE) {
		return nil
	}
	return errors.New("doesn't support serialization")
}

// BUG: has possible nil error
// TODO: add nil check to every function that calls this

// GetBackingStore returns the backing store of the record
func (tE *ServiceNowResponseImpl[T]) GetBackingStore() store.BackingStore {
	if internal.IsNil(tE) {
		return nil
	}

	if internal.IsNil(tE.backingStore) {
		if internal.IsNil(tE.backingStoreFactory) {
			return nil
		}
		tE.backingStore = tE.backingStoreFactory()
	}

	return tE.backingStore
}

// GetResult returns result from Service-Now Response
func (tE *ServiceNowResponseImpl[T]) GetResult() (T, error) {
	var empty T
	if internal.IsNil(tE) {
		return empty, nil
	}

	val, err := tE.GetBackingStore().Get(resultKey)
	if err != nil {
		return empty, err
	}

	if val == nil {
		return empty, nil
	}

	typedVal, ok := val.(T)
	if !ok {
		return empty, fmt.Errorf("val is not %T", empty)
	}

	return typedVal, nil
}

// setResult sets the slice for the Service-Now Response
func (tE *ServiceNowResponseImpl[T]) SetResult(result T) error {
	if internal.IsNil(tE) {
		return nil
	}

	return tE.GetBackingStore().Set(resultKey, result)
}

func (tE *ServiceNowResponseImpl[T]) GetFactory() (serialization.ParsableFactory, error) {
	if internal.IsNil(tE) || internal.IsNil(tE.factory) {
		return nil, nil
	}

	return tE.factory, nil
}
