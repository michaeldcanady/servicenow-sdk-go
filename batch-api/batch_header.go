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

// BatchHeaderable represents a Service-Now batch header
type BatchHeaderable interface {
	GetName() (*string, error)
	SetName(*string) error
	GetValue() (*string, error)
	SetValue(*string) error
	serialization.Parsable
	newInternal.Model
}

// BatchHeader implementation of BatchHeaderable
type BatchHeader struct {
	newInternal.Model
}

// NewBatchHeader creates new instance of BatchHeader
func NewBatchHeader() *BatchHeader {
	return &BatchHeader{
		newInternal.NewBaseModel(),
	}
}

// CreateBatchHeader2FromDiscriminatorValue is a parsable factory for creating a BatchRequestable
func CreateBatchHeader2FromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewBatchHeader(), nil
}

// Serialize writes the objects properties to the current writer.
func (bH *BatchHeader) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(bH) {
		return nil
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
func (bH *BatchHeader) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	if internal.IsNil(bH) {
		return nil
	}

	return map[string]func(serialization.ParseNode) error{
		nameKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (nameKey) not implemented")
		},
		valueKey: func(pn serialization.ParseNode) error {
			return errors.New("deserializer (valueKey) not implemented")
		},
	}
}

// GetName returns the name of the header
func (bH *BatchHeader) GetName() (*string, error) {
	if internal.IsNil(bH) {
		return nil, nil
	}

	name, err := bH.GetBackingStore().Get(nameKey)
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
func (bH *BatchHeader) SetName(name *string) error {
	if internal.IsNil(bH) {
		return nil
	}

	return bH.GetBackingStore().Set(nameKey, name)
}

// GetValue returns the value of the header
func (bH *BatchHeader) GetValue() (*string, error) {
	if internal.IsNil(bH) {
		return nil, nil
	}

	value, err := bH.GetBackingStore().Get(valueKey)
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
func (bH *BatchHeader) SetValue(value *string) error {
	if internal.IsNil(bH) {
		return nil
	}

	return bH.GetBackingStore().Set(valueKey, value)
}

// headers support headers types
type headers interface {
	*abstractions.RequestHeaders
}

// createBatchableHeadersFromHeaders converts headers to BatchHeaderable
func createBatchableHeadersFromHeaders[h headers](headers h) ([]BatchHeaderable, error) {
	batchHeaders := make([]BatchHeaderable, 0)

	if requestHeaders, ok := interface{}(headers).(*abstractions.RequestHeaders); ok {
		for _, key := range requestHeaders.ListKeys() {
			batchHeader := NewBatchHeader()
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
