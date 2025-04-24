package batchapi

import (
	"errors"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	nameKey  = "name"
	valueKey = "value"
)

// RestRequestHeader represents a Service-Now batch header
type RestRequestHeader interface {
	// GetName returns the key value for the header
	GetName() (*string, error)
	// setName sets the key value for the header
	SetName(*string) error
	// GetValue returns the value of the header
	GetValue() (*string, error)
	// SetValue sets the value of the header
	SetValue(*string) error
	serialization.Parsable
	newInternal.Model
}

// RestRequestHeaderModel implementation of RestRequestHeader
type RestRequestHeaderModel struct {
	newInternal.Model
}

// NewRestRequestHeader creates new instance of BatchHeader
func NewRestRequestHeader() *RestRequestHeaderModel {
	return &RestRequestHeaderModel{
		newInternal.NewBaseModel(),
	}
}

// CreateRestRequestHeaderFromDiscriminatorValue is a parsable factory for creating a BatchRequest
func CreateRestRequestHeaderFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRestRequestHeader(), nil
}

// Serialize writes the objects properties to the current writer.
func (bH *RestRequestHeaderModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bH) {
		return nil
	}

	if internal.IsNil(writer) {
		return errors.New("write is nil")
	}

	serializers := []func(serialization.SerializationWriter) error{
		func(sw serialization.SerializationWriter) error {
			name, err := bH.GetName()
			if err != nil {
				return err
			}
			return sw.WriteStringValue(nameKey, name)
		},
		func(sw serialization.SerializationWriter) error {
			value, err := bH.GetValue()
			if err != nil {
				return err
			}
			return sw.WriteStringValue(valueKey, value)
		},
	}

	for _, serializer := range serializers {
		if err := serializer(writer); err != nil {
			return err
		}
	}
	return nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (bH *RestRequestHeaderModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(bH) {
		return nil
	}

	return nil
}

// GetName returns the name of the header
func (bH *RestRequestHeaderModel) GetName() (*string, error) {
	if internal.IsNil(bH) {
		return nil, nil
	}

	backingStore := bH.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("backingStore is nil")
	}

	name, err := backingStore.Get(nameKey)
	if err != nil {
		return nil, err
	}

	typedName, ok := name.(*string)
	if !ok {
		return nil, errors.New("name is not *string")
	}

	return typedName, nil
}

// SetName sets name to provided value
func (bH *RestRequestHeaderModel) SetName(name *string) error {
	if internal.IsNil(bH) {
		return nil
	}

	backingStore := bH.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("backingStore is nil")
	}

	return backingStore.Set(nameKey, name)
}

// GetValue returns the value of the header
func (bH *RestRequestHeaderModel) GetValue() (*string, error) {
	if internal.IsNil(bH) {
		return nil, nil
	}

	backingStore := bH.GetBackingStore()
	if internal.IsNil(backingStore) {
		return nil, errors.New("backingStore is nil")
	}

	value, err := backingStore.Get(valueKey)
	if err != nil {
		return nil, err
	}

	typedValue, ok := value.(*string)
	if !ok {
		return nil, errors.New("value is not *string")
	}

	return typedValue, nil
}

// SetValue sets the value to the provided value
func (bH *RestRequestHeaderModel) SetValue(value *string) error {
	if internal.IsNil(bH) {
		return nil
	}

	backingStore := bH.GetBackingStore()
	if internal.IsNil(backingStore) {
		return errors.New("backingStore is nil")
	}

	return backingStore.Set(valueKey, value)
}

// headers support headers types
type headers interface {
	*abstractions.RequestHeaders
}

// createRestRequestHeaderFromHeaders converts headers to RestRequestHeader
func createRestRequestHeaderFromHeaders[h headers](headers h) ([]RestRequestHeader, error) {
	batchHeaders := make([]RestRequestHeader, 0)

	if requestHeaders, ok := interface{}(headers).(*abstractions.RequestHeaders); ok {
		for _, key := range requestHeaders.ListKeys() {
			batchHeader := NewRestRequestHeader()
			values := requestHeaders.Get(key)
			if err := batchHeader.SetName(&key); err != nil {
				return nil, err
			}
			valuesString := strings.Join(values, ", ")
			if err := batchHeader.SetValue(&valuesString); err != nil {
				return nil, err
			}
			batchHeaders = append(batchHeaders, batchHeader)
		}
		return batchHeaders, nil
	}

	return nil, nil
}
